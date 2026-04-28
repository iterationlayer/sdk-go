package iterationlayer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExtractDocumentPostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/document-extraction/v1/extract" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Fatalf("missing auth header")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true, "data": {}}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	_, err := client.ExtractDocument(ExtractDocumentRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestExtractDocumentAsyncPostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"success": true, "async": true, "message": "Processing started"}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	result, err := client.ExtractDocumentAsync(ExtractDocumentAsyncRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Message != "Processing started" {
		b, _ := json.Marshal(result)
		t.Fatalf("unexpected result: %s", b)
	}
}

func TestGenerateDocumentPostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/document-generation/v1/generate" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Fatalf("missing auth header")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true, "data": {}}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	_, err := client.GenerateDocument(GenerateDocumentRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGenerateDocumentAsyncPostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"success": true, "async": true, "message": "Processing started"}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	result, err := client.GenerateDocumentAsync(GenerateDocumentAsyncRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Message != "Processing started" {
		b, _ := json.Marshal(result)
		t.Fatalf("unexpected result: %s", b)
	}
}

func TestConvertDocumentToMarkdownPostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/document-to-markdown/v1/convert" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Fatalf("missing auth header")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true, "data": {}}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	_, err := client.ConvertDocumentToMarkdown(ConvertDocumentToMarkdownRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestConvertDocumentToMarkdownAsyncPostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"success": true, "async": true, "message": "Processing started"}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	result, err := client.ConvertDocumentToMarkdownAsync(ConvertDocumentToMarkdownAsyncRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Message != "Processing started" {
		b, _ := json.Marshal(result)
		t.Fatalf("unexpected result: %s", b)
	}
}

func TestGenerateImagePostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/image-generation/v1/generate" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Fatalf("missing auth header")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true, "data": {}}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	_, err := client.GenerateImage(GenerateImageRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGenerateImageAsyncPostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"success": true, "async": true, "message": "Processing started"}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	result, err := client.GenerateImageAsync(GenerateImageAsyncRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Message != "Processing started" {
		b, _ := json.Marshal(result)
		t.Fatalf("unexpected result: %s", b)
	}
}

func TestTransformImagePostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/image-transformation/v1/transform" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Fatalf("missing auth header")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true, "data": {}}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	_, err := client.TransformImage(TransformImageRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTransformImageAsyncPostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"success": true, "async": true, "message": "Processing started"}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	result, err := client.TransformImageAsync(TransformImageAsyncRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Message != "Processing started" {
		b, _ := json.Marshal(result)
		t.Fatalf("unexpected result: %s", b)
	}
}

func TestGenerateSheetPostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/sheet-generation/v1/generate" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Fatalf("missing auth header")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true, "data": {}}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	_, err := client.GenerateSheet(GenerateSheetRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGenerateSheetAsyncPostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"success": true, "async": true, "message": "Processing started"}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	result, err := client.GenerateSheetAsync(GenerateSheetAsyncRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Message != "Processing started" {
		b, _ := json.Marshal(result)
		t.Fatalf("unexpected result: %s", b)
	}
}

func TestExtractWebsitePostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/website-extraction/v1/extract" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Fatalf("missing auth header")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true, "data": {}}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	_, err := client.ExtractWebsite(ExtractWebsiteRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestExtractWebsiteAsyncPostsToEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"success": true, "async": true, "message": "Processing started"}`))
	}))
	defer server.Close()
	client := NewClient("test-key", WithBaseURL(server.URL))
	result, err := client.ExtractWebsiteAsync(ExtractWebsiteAsyncRequest{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Message != "Processing started" {
		b, _ := json.Marshal(result)
		t.Fatalf("unexpected result: %s", b)
	}
}

func TestBinaryFieldsEncodeAndDecode(t *testing.T) {
	requestBody, err := json.Marshal(struct {
		Base64 []byte `json:"base64"`
	}{Base64: []byte{1, 2, 3}})
	if err != nil {
		t.Fatalf("failed to marshal request: %v", err)
	}
	if !bytes.Contains(requestBody, []byte(`"base64":"AQID"`)) {
		t.Fatalf("request did not encode base64 bytes: %s", requestBody)
	}

	var result BinaryResult
	if err := json.Unmarshal([]byte(`{"buffer":"AQID","mime_type":"application/octet-stream"}`), &result); err != nil {
		t.Fatalf("failed to unmarshal binary result: %v", err)
	}
	if !bytes.Equal(result.Buffer, []byte{1, 2, 3}) {
		t.Fatalf("unexpected buffer: %v", result.Buffer)
	}
	if result.MimeType != "application/octet-stream" {
		t.Fatalf("unexpected mime type: %s", result.MimeType)
	}
}
