package iterationlayer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	defaultBaseURL    = "https://api.iterationlayer.com"
	defaultTimeoutInS = 300
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

type Option func(*Client)

func WithBaseURL(baseURL string) Option {
	return func(client *Client) {
		client.baseURL = baseURL
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

func NewClient(apiKey string, opts ...Option) *Client {
	client := &Client{
		apiKey:  apiKey,
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: defaultTimeoutInS * time.Second,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

type Error struct {
	StatusCode   int
	ErrorMessage string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Iteration Layer API error (%d): %s", e.StatusCode, e.ErrorMessage)
}

func (c *Client) ConvertToMarkdown(req ConvertRequest) (*MarkdownFileResult, error) {
	rawData, err := c.post("/document-to-markdown/v1/convert", req)
	if err != nil {
		return nil, err
	}

	var result MarkdownFileResult
	if err := json.Unmarshal(rawData, &result); err != nil {
		return nil, fmt.Errorf("failed to parse convert result: %w", err)
	}

	return &result, nil
}

func (c *Client) Extract(req ExtractRequest) (*ExtractionResult, error) {
	rawData, err := c.post("/document-extraction/v1/extract", req)
	if err != nil {
		return nil, err
	}

	var extraction ExtractionResult
	if err := json.Unmarshal(rawData, &extraction); err != nil {
		return nil, fmt.Errorf("failed to parse extraction result: %w", err)
	}

	return &extraction, nil
}

func (c *Client) ExtractAsync(req ExtractAsyncRequest) (*AsyncResult, error) {
	return c.postAsync("/document-extraction/v1/extract", req)
}

func (c *Client) Transform(req TransformRequest) (*BinaryResult, error) {
	rawData, err := c.post("/image-transformation/v1/transform", req)
	if err != nil {
		return nil, err
	}

	var binary BinaryResult
	if err := json.Unmarshal(rawData, &binary); err != nil {
		return nil, fmt.Errorf("failed to parse transform result: %w", err)
	}

	return &binary, nil
}

func (c *Client) TransformAsync(req TransformAsyncRequest) (*AsyncResult, error) {
	return c.postAsync("/image-transformation/v1/transform", req)
}

func (c *Client) GenerateImage(req GenerateImageRequest) (*BinaryResult, error) {
	rawData, err := c.post("/image-generation/v1/generate", req)
	if err != nil {
		return nil, err
	}

	var binary BinaryResult
	if err := json.Unmarshal(rawData, &binary); err != nil {
		return nil, fmt.Errorf("failed to parse image generation result: %w", err)
	}

	return &binary, nil
}

func (c *Client) GenerateImageAsync(req GenerateImageAsyncRequest) (*AsyncResult, error) {
	return c.postAsync("/image-generation/v1/generate", req)
}

func (c *Client) GenerateDocument(req GenerateDocumentRequest) (*BinaryResult, error) {
	rawData, err := c.post("/document-generation/v1/generate", req)
	if err != nil {
		return nil, err
	}

	var binary BinaryResult
	if err := json.Unmarshal(rawData, &binary); err != nil {
		return nil, fmt.Errorf("failed to parse document generation result: %w", err)
	}

	return &binary, nil
}

func (c *Client) GenerateDocumentAsync(req GenerateDocumentAsyncRequest) (*AsyncResult, error) {
	return c.postAsync("/document-generation/v1/generate", req)
}

func (c *Client) GenerateSheet(req GenerateSheetRequest) (*BinaryResult, error) {
	rawData, err := c.post("/sheet-generation/v1/generate", req)
	if err != nil {
		return nil, err
	}

	var binary BinaryResult
	if err := json.Unmarshal(rawData, &binary); err != nil {
		return nil, fmt.Errorf("failed to parse sheet generation result: %w", err)
	}

	return &binary, nil
}

func (c *Client) GenerateSheetAsync(req GenerateSheetAsyncRequest) (*AsyncResult, error) {
	return c.postAsync("/sheet-generation/v1/generate", req)
}

func (c *Client) post(path string, body any) (json.RawMessage, error) {
	respBody, statusCode, err := c.doRequest(path, body)
	if err != nil {
		return nil, err
	}

	var parsed struct {
		Success bool            `json:"success"`
		Data    json.RawMessage `json:"data"`
		Error   string          `json:"error"`
	}

	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if !parsed.Success {
		return nil, &Error{
			StatusCode:   statusCode,
			ErrorMessage: parsed.Error,
		}
	}

	return parsed.Data, nil
}

func (c *Client) postAsync(path string, body any) (*AsyncResult, error) {
	respBody, statusCode, err := c.doRequest(path, body)
	if err != nil {
		return nil, err
	}

	var parsed struct {
		Success bool   `json:"success"`
		Async   bool   `json:"async"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}

	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if !parsed.Success {
		return nil, &Error{
			StatusCode:   statusCode,
			ErrorMessage: parsed.Error,
		}
	}

	return &AsyncResult{Message: parsed.Message}, nil
}

func (c *Client) doRequest(path string, body any) ([]byte, int, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", c.baseURL+path, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read response: %w", err)
	}

	return respBody, resp.StatusCode, nil
}
