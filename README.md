# Iteration Layer Go SDK

Official Go SDK for the [Iteration Layer API](https://iterationlayer.com).

## Installation

```sh
go get github.com/iterationlayer/sdk-go
```

## Usage

```go
package main

import (
    il "github.com/iterationlayer/sdk-go"
)

func main() {
    client := il.NewClient("il_your_api_key")
}
```

### Document Extraction

Extract structured data from documents using AI.

```go
result, err := client.ExtractDocument(il.ExtractDocumentRequest{
    Files: []il.FileInput{
        il.NewFileFromURL("invoice.pdf", "https://example.com/invoice.pdf"),
    },
    Schema: il.ExtractionSchema{
        "company_name": il.NewTextFieldConfig("company_name", "The company name"),
        "total":        il.NewCurrencyAmountFieldConfig("total", "The invoice total"),
    },
})
if err != nil {
    log.Fatal(err)
}

companyName := (*result)["company_name"]
fmt.Println(companyName.Value)      // "Acme Corp"
fmt.Println(companyName.Confidence) // 0.95
```

### Website Extraction

Extract structured data from public website pages. Static fetching is used by default; set `ShouldRenderJavascript` when a page needs browser rendering.

```go
shouldRenderJavascript := true

result, err := client.ExtractWebsite(il.ExtractWebsiteRequest{
    File: il.FileInputURL{
        Type: "url",
        URL:  "https://example.com/pricing",
        FetchOptions: &il.FileFetchOptions{
            ShouldRenderJavascript: &shouldRenderJavascript,
        },
    },
    Schema: il.ExtractionSchema{
        "plan_name": il.NewTextFieldConfig("plan_name", "The pricing plan name"),
        "price":     il.NewCurrencyAmountFieldConfig("price", "The monthly price"),
    },
})
if err != nil {
    log.Fatal(err)
}

planName := (*result)["plan_name"]
fmt.Println(planName.Value) // "Pro"
```

### Image Transformation

Resize, crop, convert, and apply effects to images.

```go
result, err := client.TransformImage(il.TransformImageRequest{
    File: il.NewFileFromURL("photo.jpg", "https://example.com/photo.jpg"),
    Operations: []il.TransformOperation{
        il.NewResizeOperation(800, 600, "cover"),
        il.NewConvertOperation("webp"),
    },
})
if err != nil {
    log.Fatal(err)
}

imageBase64 := result.Buffer
```

### Image Generation

Compose images from layer definitions.

```go
result, err := client.GenerateImage(il.GenerateImageRequest{
    Dimensions: il.Dimensions{Width: 1200, Height: 630},
    Layers: []il.Layer{
        il.NewSolidColorBackgroundLayer(0, "#1a1a2e"),
        il.NewTextLayer(
            1,
            "Hello World",
            "Inter",
            48,
            "#ffffff",
            il.Position{X: 50, Y: 50},
            il.Dimensions{Width: 1100, Height: 530},
        ),
    },
    OutputFormat: "png",
})
if err != nil {
    log.Fatal(err)
}

imageBase64 := result.Buffer
```

### Document Generation

Generate PDF, DOCX, EPUB, or PPTX from structured data.

```go
result, err := client.GenerateDocument(il.GenerateDocumentRequest{
    Format: "pdf",
    Document: il.DocumentDefinition{
        Metadata: il.DocumentMetadata{Title: "Invoice #123"},
        Page: il.DocumentPage{
            Size:    il.DocPageSize{Preset: "A4"},
            Margins: il.DocMargins{TopInPt: 36, BottomInPt: 36, LeftInPt: 36, RightInPt: 36},
        },
        Styles: il.DocumentStyles{
            Text:      il.TextStyle{FontFamily: "Helvetica", FontSizeInPt: 12, LineHeight: 1.5, Color: "#000000"},
            Headline:  il.HeadlineStyle{FontFamily: "Helvetica", FontSizeInPt: 24, Color: "#000000", SpacingBeforeInPt: 12, SpacingAfterInPt: 6},
            Link:      il.LinkStyle{Color: "#0066cc"},
            List:      il.ListStyle{IndentInPt: 18, SpacingBetweenItemsInPt: 4},
            Table: il.TableStyle{
                Header: il.TableHeaderStyle{BackgroundColor: "#f0f0f0", FontFamily: "Helvetica", FontSizeInPt: 12, Color: "#000000", PaddingInPt: 6},
                Body:   il.TableBodyStyle{FontFamily: "Helvetica", FontSizeInPt: 12, Color: "#000000", PaddingInPt: 6},
            },
            Grid:      il.GridStyle{GapInPt: 12},
            Separator: il.SeparatorStyle{Color: "#cccccc", ThicknessInPt: 1, MarginTopInPt: 12, MarginBottomInPt: 12},
            Image:     il.ImageStyle{Alignment: "center", MarginTopInPt: 8, MarginBottomInPt: 8},
        },
        Content: []il.ContentBlock{
            il.NewHeadlineBlock("h1", "Invoice #123"),
            il.NewParagraphBlock(),
        },
    },
})
if err != nil {
    log.Fatal(err)
}

pdfBase64 := result.Buffer
```

### Sheet Generation

Generate CSV, Markdown, or XLSX spreadsheets from structured data.

```go
result, err := client.GenerateSheet(il.GenerateSheetRequest{
    Format: "xlsx",
    Sheets: []il.Sheet{
        {
            Name: "Invoices",
            Columns: []il.SheetColumn{
                {Name: "Company", Width: 20},
                {Name: "Total", Width: 15},
            },
            Rows: [][]il.SheetCell{
                {
                    {Value: "Acme Corp"},
                    {Value: 1500.50, Format: "currency", CurrencyCode: "EUR"},
                },
            },
        },
    },
})
if err != nil {
    log.Fatal(err)
}

sheetBase64 := result.Buffer
```

### Webhooks (Async)

Use the `*Async` methods to receive results via webhook instead of waiting for the response.

```go
result, err := client.ExtractDocumentAsync(il.ExtractDocumentAsyncRequest{
    Files: []il.FileInput{
        il.NewFileFromURL("invoice.pdf", "https://example.com/invoice.pdf"),
    },
    Schema: il.ExtractionSchema{
        "total": il.NewCurrencyAmountFieldConfig("total", "The invoice total"),
    },
    WebhookURL: "https://your-app.com/webhooks/extraction",
})

fmt.Println(result.Message) // "Request accepted..."
```

### Error Handling

```go
result, err := client.ExtractDocument(req)
if err != nil {
    var apiErr *il.Error
    if errors.As(err, &apiErr) {
        fmt.Println(apiErr.StatusCode)    // 422
        fmt.Println(apiErr.ErrorMessage)  // "Validation error: ..."
    }
}
```

### Options

```go
// Custom base URL
client := il.NewClient("il_your_api_key", il.WithBaseURL("https://custom.api.com"))

// Custom HTTP client
client := il.NewClient("il_your_api_key", il.WithHTTPClient(&http.Client{Timeout: 60 * time.Second}))
```

## Documentation

Full documentation is available at [https://iterationlayer.com/docs](https://iterationlayer.com/docs).

## Issues & Feedback

Please report bugs and request features in the [issues repository](https://github.com/iterationlayer/issues).
