package iterationlayer

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("uses default base URL", func(t *testing.T) {
		client := NewClient("test-key")

		if client.baseURL != "https://api.iterationlayer.com" {
			t.Errorf("expected default base URL, got %s", client.baseURL)
		}
	})

	t.Run("uses custom base URL with WithBaseURL", func(t *testing.T) {
		client := NewClient("test-key", WithBaseURL("https://custom.example.com"))

		if client.baseURL != "https://custom.example.com" {
			t.Errorf("expected custom base URL, got %s", client.baseURL)
		}
	})

	t.Run("uses custom HTTP client with WithHTTPClient", func(t *testing.T) {
		customHTTPClient := &http.Client{}
		client := NewClient("test-key", WithHTTPClient(customHTTPClient))

		if client.httpClient != customHTTPClient {
			t.Error("expected custom HTTP client to be used")
		}
	})

	t.Run("stores API key", func(t *testing.T) {
		client := NewClient("my-secret-key")

		if client.apiKey != "my-secret-key" {
			t.Errorf("expected API key 'my-secret-key', got %s", client.apiKey)
		}
	})
}

func TestExtract(t *testing.T) {
	t.Run("sends correct request and parses extraction result", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("expected POST, got %s", r.Method)
			}

			if r.URL.Path != "/document-extraction/v1/extract" {
				t.Errorf("expected /document-extraction/v1/extract, got %s", r.URL.Path)
			}

			if r.Header.Get("Content-Type") != "application/json" {
				t.Errorf("expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
			}

			if r.Header.Get("Authorization") != "Bearer test-key" {
				t.Errorf("expected Authorization Bearer test-key, got %s", r.Header.Get("Authorization"))
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("failed to read request body: %v", err)
			}

			var requestPayload struct {
				Files  []json.RawMessage          `json:"files"`
				Schema map[string]json.RawMessage `json:"schema"`
			}
			if err := json.Unmarshal(body, &requestPayload); err != nil {
				t.Fatalf("failed to parse request body: %v", err)
			}

			if len(requestPayload.Files) != 1 {
				t.Fatalf("expected 1 file, got %d", len(requestPayload.Files))
			}

			if len(requestPayload.Schema) != 1 {
				t.Fatalf("expected 1 schema field, got %d", len(requestPayload.Schema))
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
        "success": true,
        "data": {
          "invoice_number": {
            "value": "INV-001",
            "confidence": 0.95,
            "citations": ["page 1"],
            "source": "document",
            "type": "TEXT"
          }
        }
      }`))
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		result, err := client.Extract(ExtractRequest{
			Files: []FileInput{
				NewFileFromURL("invoice.pdf", "https://example.com/invoice.pdf"),
			},
			Schema: ExtractionSchema{
				"invoice_number": NewTextFieldConfig("invoice_number", "The invoice number"),
			},
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result == nil {
			t.Fatal("expected extraction result, got nil")
		}

		fieldResult, ok := (*result)["invoice_number"]
		if !ok {
			t.Fatal("expected invoice_number field in result")
		}

		if fieldResult.Value != "INV-001" {
			t.Errorf("expected value INV-001, got %v", fieldResult.Value)
		}

		if fieldResult.Confidence != 0.95 {
			t.Errorf("expected confidence 0.95, got %f", fieldResult.Confidence)
		}

		if fieldResult.Type != "TEXT" {
			t.Errorf("expected type TEXT, got %s", fieldResult.Type)
		}
	})
}

func TestTransform(t *testing.T) {
	t.Run("sends correct request and parses binary result", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("expected POST, got %s", r.Method)
			}

			if r.URL.Path != "/image-transformation/v1/transform" {
				t.Errorf("expected /image-transformation/v1/transform, got %s", r.URL.Path)
			}

			if r.Header.Get("Authorization") != "Bearer test-key" {
				t.Errorf("expected Authorization Bearer test-key, got %s", r.Header.Get("Authorization"))
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("failed to read request body: %v", err)
			}

			var requestPayload struct {
				File       json.RawMessage   `json:"file"`
				Operations []json.RawMessage `json:"operations"`
			}
			if err := json.Unmarshal(body, &requestPayload); err != nil {
				t.Fatalf("failed to parse request body: %v", err)
			}

			if len(requestPayload.Operations) != 1 {
				t.Fatalf("expected 1 operation, got %d", len(requestPayload.Operations))
			}

			var operation struct {
				Type       string `json:"type"`
				WidthInPx  int    `json:"width_in_px"`
				HeightInPx int    `json:"height_in_px"`
				Fit        string `json:"fit"`
			}
			if err := json.Unmarshal(requestPayload.Operations[0], &operation); err != nil {
				t.Fatalf("failed to parse operation: %v", err)
			}

			if operation.Type != "resize" {
				t.Errorf("expected operation type 'resize', got %s", operation.Type)
			}

			if operation.WidthInPx != 800 || operation.HeightInPx != 600 {
				t.Errorf("expected dimensions 800x600, got %dx%d", operation.WidthInPx, operation.HeightInPx)
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
        "success": true,
        "data": {
          "buffer": "aW1hZ2VkYXRh",
          "mime_type": "image/png"
        }
      }`))
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		result, err := client.Transform(TransformRequest{
			File:       NewFileFromURL("photo.jpg", "https://example.com/photo.jpg"),
			Operations: []TransformOperation{NewResizeOperation(800, 600, "cover")},
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result == nil {
			t.Fatal("expected binary result, got nil")
		}

		if result.Buffer != "aW1hZ2VkYXRh" {
			t.Errorf("expected buffer 'aW1hZ2VkYXRh', got %s", result.Buffer)
		}

		if result.MimeType != "image/png" {
			t.Errorf("expected mime type 'image/png', got %s", result.MimeType)
		}
	})
}

func TestGenerateImage(t *testing.T) {
	t.Run("sends correct request and parses binary result", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/image-generation/v1/generate" {
				t.Errorf("expected /image-generation/v1/generate, got %s", r.URL.Path)
			}

			if r.Method != "POST" {
				t.Errorf("expected POST, got %s", r.Method)
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("failed to read request body: %v", err)
			}

			var requestPayload struct {
				Dimensions struct {
					Width  int `json:"width"`
					Height int `json:"height"`
				} `json:"dimensions"`
				Layers []json.RawMessage `json:"layers"`
			}
			if err := json.Unmarshal(body, &requestPayload); err != nil {
				t.Fatalf("failed to parse request body: %v", err)
			}

			if requestPayload.Dimensions.Width != 1200 || requestPayload.Dimensions.Height != 630 {
				t.Errorf("expected dimensions 1200x630, got %dx%d",
					requestPayload.Dimensions.Width, requestPayload.Dimensions.Height)
			}

			if len(requestPayload.Layers) != 2 {
				t.Fatalf("expected 2 layers, got %d", len(requestPayload.Layers))
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
        "success": true,
        "data": {
          "buffer": "cG5nZGF0YQ==",
          "mime_type": "image/png"
        }
      }`))
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		result, err := client.GenerateImage(GenerateImageRequest{
			Dimensions: Dimensions{Width: 1200, Height: 630},
			Layers: []Layer{
				NewSolidColorBackgroundLayer(0, "#ffffff"),
				NewTextLayer(1, "Hello", "Arial", 48, "#000000",
					Position{X: 100, Y: 100},
					Dimensions{Width: 1000, Height: 100}),
			},
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result == nil {
			t.Fatal("expected binary result, got nil")
		}

		if result.Buffer != "cG5nZGF0YQ==" {
			t.Errorf("expected buffer 'cG5nZGF0YQ==', got %s", result.Buffer)
		}

		if result.MimeType != "image/png" {
			t.Errorf("expected mime type 'image/png', got %s", result.MimeType)
		}
	})
}

func TestGenerateDocument(t *testing.T) {
	t.Run("sends correct request and parses binary result", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/document-generation/v1/generate" {
				t.Errorf("expected /document-generation/v1/generate, got %s", r.URL.Path)
			}

			if r.Method != "POST" {
				t.Errorf("expected POST, got %s", r.Method)
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("failed to read request body: %v", err)
			}

			var requestPayload struct {
				Format   string          `json:"format"`
				Document json.RawMessage `json:"document"`
			}
			if err := json.Unmarshal(body, &requestPayload); err != nil {
				t.Fatalf("failed to parse request body: %v", err)
			}

			if requestPayload.Format != "pdf" {
				t.Errorf("expected format 'pdf', got %s", requestPayload.Format)
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
        "success": true,
        "data": {
          "buffer": "cGRmZGF0YQ==",
          "mime_type": "application/pdf"
        }
      }`))
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		result, err := client.GenerateDocument(GenerateDocumentRequest{
			Format: "pdf",
			Document: DocumentDefinition{
				Metadata: DocumentMetadata{Title: "Test Document"},
				Page: &DocumentPage{
					Size:    DocPageSize{Preset: "A4"},
					Margins: DocMargins{TopInPt: 72, RightInPt: 72, BottomInPt: 72, LeftInPt: 72},
				},
				Styles: &DocumentStyles{
					Text:      TextStyle{FontFamily: "Helvetica", FontSizeInPt: 12, LineHeight: 1.5, Color: "#000000"},
					Headline:  HeadlineStyle{FontFamily: "Helvetica", FontSizeInPt: 24, Color: "#000000", SpacingBeforeInPt: 12, SpacingAfterInPt: 6},
					Link:      LinkStyle{Color: "#0000ff"},
					List:      ListStyle{MarkerColor: "#000000", MarkerGapInPt: 4, TextStyle: TextStyle{FontFamily: "Helvetica", FontSizeInPt: 12, LineHeight: 1.5, Color: "#000000"}},
					Table:     TableStyle{Header: TableHeaderStyle{BackgroundColor: "#f0f0f0", TextColor: "#000000", FontSizeInPt: 12}, Body: TableBodyStyle{BackgroundColor: "#ffffff", TextColor: "#000000", FontSizeInPt: 12}},
					Grid:      GridStyle{GapInPt: 8},
					Separator: SeparatorStyle{Color: "#cccccc", ThicknessInPt: 1, SpacingBeforeInPt: 8, SpacingAfterInPt: 8},
					Image:     ImageStyle{BorderColor: "#000000", BorderWidthInPt: 0},
				},
				Content: []ContentBlock{
					NewHeadlineBlock("h1", "Hello World"),
				},
			},
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result == nil {
			t.Fatal("expected binary result, got nil")
		}

		if result.Buffer != "cGRmZGF0YQ==" {
			t.Errorf("expected buffer 'cGRmZGF0YQ==', got %s", result.Buffer)
		}

		if result.MimeType != "application/pdf" {
			t.Errorf("expected mime type 'application/pdf', got %s", result.MimeType)
		}
	})
}

func TestGenerateSheet(t *testing.T) {
	t.Run("sends correct request and parses binary result", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/sheet-generation/v1/generate" {
				t.Errorf("expected /sheet-generation/v1/generate, got %s", r.URL.Path)
			}

			if r.Method != "POST" {
				t.Errorf("expected POST, got %s", r.Method)
			}

			if r.Header.Get("Authorization") != "Bearer test-key" {
				t.Errorf("expected Authorization Bearer test-key, got %s", r.Header.Get("Authorization"))
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("failed to read request body: %v", err)
			}

			var requestPayload struct {
				Format string            `json:"format"`
				Sheets []json.RawMessage `json:"sheets"`
			}
			if err := json.Unmarshal(body, &requestPayload); err != nil {
				t.Fatalf("failed to parse request body: %v", err)
			}

			if requestPayload.Format != "xlsx" {
				t.Errorf("expected format 'xlsx', got %s", requestPayload.Format)
			}

			if len(requestPayload.Sheets) != 1 {
				t.Fatalf("expected 1 sheet, got %d", len(requestPayload.Sheets))
			}

			var sheet struct {
				Name    string            `json:"name"`
				Columns []json.RawMessage `json:"columns"`
				Rows    []json.RawMessage `json:"rows"`
			}
			if err := json.Unmarshal(requestPayload.Sheets[0], &sheet); err != nil {
				t.Fatalf("failed to parse sheet: %v", err)
			}

			if sheet.Name != "Invoices" {
				t.Errorf("expected sheet name 'Invoices', got %s", sheet.Name)
			}

			if len(sheet.Columns) != 2 {
				t.Errorf("expected 2 columns, got %d", len(sheet.Columns))
			}

			if len(sheet.Rows) != 1 {
				t.Errorf("expected 1 row, got %d", len(sheet.Rows))
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
        "success": true,
        "data": {
          "buffer": "eGxzeGRhdGE=",
          "mime_type": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
        }
      }`))
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		result, err := client.GenerateSheet(GenerateSheetRequest{
			Format: "xlsx",
			Sheets: []Sheet{
				{
					Name: "Invoices",
					Columns: []SheetColumn{
						{Name: "Company", Width: 20},
						{Name: "Total", Width: 15},
					},
					Rows: [][]SheetCell{
						{
							{Value: "Acme Corp"},
							{Value: 1500.50, Format: "currency", CurrencyCode: "EUR"},
						},
					},
				},
			},
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result == nil {
			t.Fatal("expected binary result, got nil")
		}

		if result.Buffer != "eGxzeGRhdGE=" {
			t.Errorf("expected buffer 'eGxzeGRhdGE=', got %s", result.Buffer)
		}

		if result.MimeType != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
			t.Errorf("expected mime type 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet', got %s", result.MimeType)
		}
	})
}

func TestGenerateSheetAsync(t *testing.T) {
	t.Run("returns async result", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/sheet-generation/v1/generate" {
				t.Errorf("expected /sheet-generation/v1/generate, got %s", r.URL.Path)
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("failed to read request body: %v", err)
			}

			var requestPayload struct {
				Format     string `json:"format"`
				WebhookURL string `json:"webhook_url"`
			}
			if err := json.Unmarshal(body, &requestPayload); err != nil {
				t.Fatalf("failed to parse request body: %v", err)
			}

			if requestPayload.WebhookURL != "https://example.com/webhook" {
				t.Errorf("expected webhook_url 'https://example.com/webhook', got %s", requestPayload.WebhookURL)
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"success": true, "async": true, "message": "Queued"}`))
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		result, err := client.GenerateSheetAsync(GenerateSheetAsyncRequest{
			GenerateSheetRequest: GenerateSheetRequest{
				Format: "xlsx",
				Sheets: []Sheet{
					{
						Name:    "Sheet1",
						Columns: []SheetColumn{{Name: "Name"}},
						Rows:    [][]SheetCell{{{Value: "Test"}}},
					},
				},
			},
			WebhookURL: "https://example.com/webhook",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result.Message != "Queued" {
			t.Errorf("expected message 'Queued', got %s", result.Message)
		}
	})
}

func TestAsyncResult(t *testing.T) {
	t.Run("ExtractAsync returns async result", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
        "success": true,
        "async": true,
        "message": "Processing started. Results will be sent to your webhook."
      }`))
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		result, err := client.ExtractAsync(ExtractAsyncRequest{
			Files: []FileInput{
				NewFileFromURL("invoice.pdf", "https://example.com/invoice.pdf"),
			},
			Schema: ExtractionSchema{
				"total": NewDecimalFieldConfig("total", "The total amount"),
			},
			WebhookURL: "https://example.com/webhook",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result.Message != "Processing started. Results will be sent to your webhook." {
			t.Errorf("unexpected message: %s", result.Message)
		}
	})

	t.Run("TransformAsync returns async result", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"success": true, "async": true, "message": "Queued"}`))
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		result, err := client.TransformAsync(TransformAsyncRequest{
			File:       NewFileFromURL("img.jpg", "https://example.com/img.jpg"),
			Operations: []TransformOperation{NewResizeOperation(100, 100, "cover")},
			WebhookURL: "https://example.com/webhook",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result.Message != "Queued" {
			t.Errorf("expected message 'Queued', got %s", result.Message)
		}
	})

	t.Run("GenerateImageAsync returns async result", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"success": true, "async": true, "message": "Queued"}`))
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		result, err := client.GenerateImageAsync(GenerateImageAsyncRequest{
			Dimensions: Dimensions{Width: 1200, Height: 630},
			Layers: []Layer{
				NewSolidColorBackgroundLayer(0, "#ffffff"),
			},
			WebhookURL: "https://example.com/webhook",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result.Message != "Queued" {
			t.Errorf("expected message 'Queued', got %s", result.Message)
		}
	})

	t.Run("GenerateDocumentAsync returns async result", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"success": true, "async": true, "message": "Queued"}`))
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		result, err := client.GenerateDocumentAsync(GenerateDocumentAsyncRequest{
			Format: "pdf",
			Document: DocumentDefinition{
				Metadata: DocumentMetadata{Title: "Test"},
				Page: &DocumentPage{
					Size:    DocPageSize{Preset: "A4"},
					Margins: DocMargins{TopInPt: 72, RightInPt: 72, BottomInPt: 72, LeftInPt: 72},
				},
				Styles: &DocumentStyles{
					Text:      TextStyle{FontFamily: "Helvetica", FontSizeInPt: 12, LineHeight: 1.5, Color: "#000000"},
					Headline:  HeadlineStyle{FontFamily: "Helvetica", FontSizeInPt: 24, Color: "#000000", SpacingBeforeInPt: 12, SpacingAfterInPt: 6},
					Link:      LinkStyle{Color: "#0000ff"},
					List:      ListStyle{MarkerColor: "#000000", MarkerGapInPt: 4, TextStyle: TextStyle{FontFamily: "Helvetica", FontSizeInPt: 12, LineHeight: 1.5, Color: "#000000"}},
					Table:     TableStyle{Header: TableHeaderStyle{BackgroundColor: "#f0f0f0", TextColor: "#000000", FontSizeInPt: 12}, Body: TableBodyStyle{BackgroundColor: "#ffffff", TextColor: "#000000", FontSizeInPt: 12}},
					Grid:      GridStyle{GapInPt: 8},
					Separator: SeparatorStyle{Color: "#cccccc", ThicknessInPt: 1, SpacingBeforeInPt: 8, SpacingAfterInPt: 8},
					Image:     ImageStyle{BorderColor: "#000000", BorderWidthInPt: 0},
				},
				Content: []ContentBlock{
					NewHeadlineBlock("h1", "Hello World"),
				},
			},
			WebhookURL: "https://example.com/webhook",
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result.Message != "Queued" {
			t.Errorf("expected message 'Queued', got %s", result.Message)
		}
	})
}

func TestErrorHandling(t *testing.T) {
	t.Run("returns Error when response has success false", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"success": false, "error": "Invalid API key"}`))
		}))
		defer server.Close()

		client := NewClient("bad-key", WithBaseURL(server.URL))
		result, err := client.Extract(ExtractRequest{
			Files: []FileInput{
				NewFileFromURL("doc.pdf", "https://example.com/doc.pdf"),
			},
			Schema: ExtractionSchema{
				"name": NewTextFieldConfig("name", "Name"),
			},
		})

		if result != nil {
			t.Error("expected nil result on error")
		}

		if err == nil {
			t.Fatal("expected error, got nil")
		}

		apiErr, ok := err.(*Error)
		if !ok {
			t.Fatalf("expected *Error, got %T", err)
		}

		if apiErr.StatusCode != http.StatusUnauthorized {
			t.Errorf("expected status 401, got %d", apiErr.StatusCode)
		}

		if apiErr.ErrorMessage != "Invalid API key" {
			t.Errorf("expected error message 'Invalid API key', got %s", apiErr.ErrorMessage)
		}

		expectedErrStr := "Iteration Layer API error (401): Invalid API key"
		if apiErr.Error() != expectedErrStr {
			t.Errorf("expected error string %q, got %q", expectedErrStr, apiErr.Error())
		}
	})

	t.Run("returns Error for bad request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"success": false, "error": "Missing required field: files"}`))
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		_, err := client.Transform(TransformRequest{
			File:       NewFileFromURL("img.jpg", "https://example.com/img.jpg"),
			Operations: []TransformOperation{},
		})

		if err == nil {
			t.Fatal("expected error, got nil")
		}

		apiErr, ok := err.(*Error)
		if !ok {
			t.Fatalf("expected *Error, got %T", err)
		}

		if apiErr.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", apiErr.StatusCode)
		}
	})
}

func TestRequestSerialization(t *testing.T) {
	t.Run("NewFileFromURL serializes with correct type field", func(t *testing.T) {
		file := NewFileFromURL("doc.pdf", "https://example.com/doc.pdf")
		jsonBytes, err := json.Marshal(file)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed map[string]string
		if err := json.Unmarshal(jsonBytes, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed["type"] != "url" {
			t.Errorf("expected type 'url', got %s", parsed["type"])
		}

		if parsed["name"] != "doc.pdf" {
			t.Errorf("expected name 'doc.pdf', got %s", parsed["name"])
		}

		if parsed["url"] != "https://example.com/doc.pdf" {
			t.Errorf("expected url 'https://example.com/doc.pdf', got %s", parsed["url"])
		}
	})

	t.Run("NewFileFromBase64 serializes with correct type field", func(t *testing.T) {
		file := NewFileFromBase64("image.png", "aW1hZ2VkYXRh")
		jsonBytes, err := json.Marshal(file)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed map[string]string
		if err := json.Unmarshal(jsonBytes, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed["type"] != "base64" {
			t.Errorf("expected type 'base64', got %s", parsed["type"])
		}

		if parsed["base64"] != "aW1hZ2VkYXRh" {
			t.Errorf("expected base64 'aW1hZ2VkYXRh', got %s", parsed["base64"])
		}
	})

	t.Run("NewResizeOperation serializes with correct type field", func(t *testing.T) {
		operation := NewResizeOperation(800, 600, "cover")
		jsonBytes, err := json.Marshal(operation)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed map[string]interface{}
		if err := json.Unmarshal(jsonBytes, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed["type"] != "resize" {
			t.Errorf("expected type 'resize', got %v", parsed["type"])
		}

		if parsed["width_in_px"] != float64(800) {
			t.Errorf("expected width_in_px 800, got %v", parsed["width_in_px"])
		}

		if parsed["height_in_px"] != float64(600) {
			t.Errorf("expected height_in_px 600, got %v", parsed["height_in_px"])
		}

		if parsed["fit"] != "cover" {
			t.Errorf("expected fit 'cover', got %v", parsed["fit"])
		}
	})

	t.Run("NewTextFieldConfig serializes with correct type field", func(t *testing.T) {
		field := NewTextFieldConfig("invoice_number", "The invoice number")
		jsonBytes, err := json.Marshal(field)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed map[string]interface{}
		if err := json.Unmarshal(jsonBytes, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed["type"] != "TEXT" {
			t.Errorf("expected type 'TEXT', got %v", parsed["type"])
		}

		if parsed["name"] != "invoice_number" {
			t.Errorf("expected name 'invoice_number', got %v", parsed["name"])
		}

		if parsed["description"] != "The invoice number" {
			t.Errorf("expected description 'The invoice number', got %v", parsed["description"])
		}
	})

	t.Run("NewSolidColorBackgroundLayer serializes with correct type field", func(t *testing.T) {
		layer := NewSolidColorBackgroundLayer(0, "#ff0000")
		jsonBytes, err := json.Marshal(layer)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed map[string]interface{}
		if err := json.Unmarshal(jsonBytes, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed["type"] != "solid-color" {
			t.Errorf("expected type 'solid-color', got %v", parsed["type"])
		}

		if parsed["hex_color"] != "#ff0000" {
			t.Errorf("expected hex_color '#ff0000', got %v", parsed["hex_color"])
		}
	})

	t.Run("NewRemoveBackgroundOperation serializes with correct type field", func(t *testing.T) {
		operation := NewRemoveBackgroundOperation()
		jsonBytes, err := json.Marshal(operation)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed map[string]interface{}
		if err := json.Unmarshal(jsonBytes, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed["type"] != "remove_background" {
			t.Errorf("expected type 'remove_background', got %v", parsed["type"])
		}
	})

	t.Run("ConvertToMarkdown sends correct request and parses result", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("expected POST, got %s", r.Method)
			}

			if r.URL.Path != "/document-to-markdown/v1/convert" {
				t.Errorf("expected /document-to-markdown/v1/convert, got %s", r.URL.Path)
			}

			if r.Header.Get("Authorization") != "Bearer test-key" {
				t.Errorf("expected Authorization Bearer test-key, got %s", r.Header.Get("Authorization"))
			}

			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("failed to read request body: %v", err)
			}

			var requestPayload struct {
				File json.RawMessage `json:"file"`
			}
			if err := json.Unmarshal(body, &requestPayload); err != nil {
				t.Fatalf("failed to parse request body: %v", err)
			}

			if requestPayload.File == nil {
				t.Fatal("expected file field in request body")
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
        "success": true,
        "data": {
          "name": "doc.pdf",
          "mime_type": "application/pdf",
          "markdown": "# Document\n\nSome content."
        }
      }`))
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		result, err := client.ConvertToMarkdown(ConvertRequest{
			File: NewFileFromURL("doc.pdf", "https://example.com/doc.pdf"),
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result == nil {
			t.Fatal("expected result, got nil")
		}

		if result.Name != "doc.pdf" {
			t.Errorf("expected name 'doc.pdf', got %s", result.Name)
		}

		if result.MimeType != "application/pdf" {
			t.Errorf("expected mime type 'application/pdf', got %s", result.MimeType)
		}

		if result.Markdown != "# Document\n\nSome content." {
			t.Errorf("expected markdown content, got %s", result.Markdown)
		}
	})

	t.Run("NewArrayFieldConfig serializes with nested item schema", func(t *testing.T) {
		field := NewArrayFieldConfig("line_items", "Invoice line items", []FieldConfig{
			NewTextFieldConfig("description", "Item description"),
			NewDecimalFieldConfig("amount", "Item amount"),
		})
		jsonBytes, err := json.Marshal(field)
		if err != nil {
			t.Fatalf("failed to marshal: %v", err)
		}

		var parsed map[string]interface{}
		if err := json.Unmarshal(jsonBytes, &parsed); err != nil {
			t.Fatalf("failed to unmarshal: %v", err)
		}

		if parsed["type"] != "ARRAY" {
			t.Errorf("expected type 'ARRAY', got %v", parsed["type"])
		}

		itemSchema, ok := parsed["item_schema"].(map[string]interface{})
		if !ok {
			t.Fatal("expected item_schema to be an object")
		}

		fields, ok := itemSchema["fields"].([]interface{})
		if !ok {
			t.Fatal("expected item_schema.fields to be an array")
		}

		if len(fields) != 2 {
			t.Errorf("expected 2 fields in item_schema, got %d", len(fields))
		}
	})
}
