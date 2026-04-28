package iterationlayer

type BinaryResult struct {
	Buffer   []byte `json:"buffer"`
	MimeType string `json:"mime_type"`
}

type ExtractionFieldResult struct {
	Value      any      `json:"value"`
	Confidence float64  `json:"confidence"`
	Citations  []string `json:"citations"`
	Source     string   `json:"source"`
	Type       string   `json:"type"`
}

type ExtractionResult map[string]ExtractionFieldResult

type MarkdownFileResult struct {
	Name        string `json:"name"`
	MimeType    string `json:"mime_type"`
	Markdown    string `json:"markdown"`
	Description string `json:"description,omitempty"`
}

type AsyncResult struct {
	Message string `json:"message"`
}

type AddressFieldConfig struct {
	// AllowedCountryCodes List of ISO 3166-1 alpha-2 country codes to restrict valid addresses to.
	AllowedCountryCodes []string `json:"allowed_country_codes,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'ADDRESS' for a structured address field.
	Type string `json:"type"`
}
type AngledEdge struct {
	// AngleInDegrees Angle in degrees, from -45 to 45.
	AngleInDegrees float64 `json:"angle_in_degrees"`
	// Edge Which edge to angle.
	Edge string `json:"edge"`
}
type ArrayFieldConfig struct {
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// Fields List of field configurations defining the schema for each item in the array.
	Fields []any `json:"fields"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'ARRAY' for a repeating list of structured items.
	Type string `json:"type"`
}
type AutoContrastOperation struct {
	// Type Must be 'auto_contrast'.
	Type string `json:"type"`
}
type BarcodeBlock struct {
	// BgHexColor Background color as a 6-digit hex code.
	BgHexColor string `json:"bg_hex_color,omitempty"`
	// FgHexColor Foreground (bar) color as a 6-digit hex code.
	FgHexColor string `json:"fg_hex_color,omitempty"`
	// Format Barcode format.
	Format string `json:"format"`
	// HeightInPt Height in points. Must be >= 1.
	HeightInPt float64 `json:"height_in_pt"`
	// Type Must be 'barcode'.
	Type string `json:"type"`
	// Value Data to encode in the barcode.
	Value string `json:"value"`
	// WidthInPt Width in points. Must be >= 1.
	WidthInPt float64 `json:"width_in_pt"`
}
type BarcodeLayer struct {
	// BackgroundHexColor Background color as a 6-digit hex code.
	BackgroundHexColor string `json:"background_hex_color"`
	// Dimensions defines the dimensions.
	Dimensions Dimensions `json:"dimensions"`
	// ForegroundHexColor Bar color as a 6-digit hex code.
	ForegroundHexColor string `json:"foreground_hex_color"`
	// Format Barcode format.
	Format string `json:"format"`
	// Index Z-order index of this layer. Must be >= 0.
	Index int `json:"index"`
	// Opacity Layer opacity from 0 (transparent) to 100 (opaque). Defaults to 100.
	Opacity int `json:"opacity,omitempty"`
	// Position defines the position.
	Position Position `json:"position"`
	// RotationInDegrees Rotation angle in degrees, from -180 to 180.
	RotationInDegrees float64 `json:"rotation_in_degrees,omitempty"`
	// Type Must be 'barcode'.
	Type string `json:"type"`
	// Value Data to encode in the barcode.
	Value string `json:"value"`
}
type BlurOperation struct {
	// Sigma Gaussian blur radius. Must be > 0.
	Sigma float64 `json:"sigma"`
	// Type Must be 'blur'.
	Type string `json:"type"`
}
type BooleanFieldConfig struct {
	// DefaultValue Default boolean value if the field is not found in the document.
	DefaultValue bool `json:"default_value,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'BOOLEAN' for a true/false field.
	Type string `json:"type"`
}
type BorderStyle struct {
	// Color Border color as a 6-digit hex code.
	Color string `json:"color"`
	// WidthInPt Border width in points. Must be >= 0.
	WidthInPt float64 `json:"width_in_pt"`
}
type BorderStyleOverrides struct {
	// Color Override border color as a 6-digit hex code.
	Color string `json:"color,omitempty"`
	// WidthInPt Override border width in points. Must be >= 0.
	WidthInPt float64 `json:"width_in_pt,omitempty"`
}
type CalculatedFieldConfig struct {
	// Description Human-readable description guiding the AI on what this calculated field represents.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Operation Arithmetic operation to apply to the source fields.
	Operation string `json:"operation"`
	// SourceFieldNames Ordered list of field names whose values are used as operands.
	SourceFieldNames []string `json:"source_field_names"`
	// Type Must be 'CALCULATED' for a field whose value is computed from other fields.
	Type string `json:"type"`
	// Unit Unit of measurement for the result (e.g., 'EUR', 'kg').
	Unit string `json:"unit,omitempty"`
}
type CellStyle struct {
	// BackgroundColor Cell background color as a 6-digit hex code.
	BackgroundColor string `json:"background_color,omitempty"`
	// FontColor Text color as a 6-digit hex code.
	FontColor string `json:"font_color,omitempty"`
	// FontFamily Font family name for the cell text.
	FontFamily string `json:"font_family,omitempty"`
	// FontSizeInPt Font size in points. Must be >= 1.
	FontSizeInPt float64 `json:"font_size_in_pt,omitempty"`
	// HorizontalAlignment Horizontal text alignment within the cell.
	HorizontalAlignment string `json:"horizontal_alignment,omitempty"`
	// IsBold Whether the text is bold.
	IsBold bool `json:"is_bold,omitempty"`
	// IsItalic Whether the text is italic.
	IsItalic bool `json:"is_italic,omitempty"`
	// NumberFormat Number format pattern (e.g., '#,##0.00' or '0%').
	NumberFormat string `json:"number_format,omitempty"`
}
type Column struct {
	// Name Column header name. Must not be empty.
	Name string `json:"name"`
	// Width Column width as a fraction. Must be > 0.
	Width float64 `json:"width,omitempty"`
}
type CompressToSizeOperation struct {
	// MaxFileSizeInBytes Target maximum file size in bytes. Must be > 0.
	MaxFileSizeInBytes int `json:"max_file_size_in_bytes"`
	// Type Must be 'compress_to_size'.
	Type string `json:"type"`
}
type ConvertOperation struct {
	// Format Target image format.
	Format string `json:"format"`
	// Quality Compression quality from 1 (lowest) to 100 (highest).
	Quality int `json:"quality,omitempty"`
	// Type Must be 'convert'.
	Type string `json:"type"`
}
type CountryFieldConfig struct {
	// DefaultValue Default ISO 3166-1 alpha-2 country code if the field is not found in the document.
	DefaultValue string `json:"default_value,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'COUNTRY' for an ISO 3166-1 alpha-2 country code field.
	Type string `json:"type"`
}
type CropOperation struct {
	// HeightInPx Height of the crop area in pixels. Must be > 0.
	HeightInPx int `json:"height_in_px"`
	// LeftInPx Left offset in pixels for the crop area. Must be >= 0.
	LeftInPx int `json:"left_in_px"`
	// TopInPx Top offset in pixels for the crop area. Must be >= 0.
	TopInPx int `json:"top_in_px"`
	// Type Must be 'crop'.
	Type string `json:"type"`
	// WidthInPx Width of the crop area in pixels. Must be > 0.
	WidthInPx int `json:"width_in_px"`
}
type CurrencyAmountFieldConfig struct {
	// DecimalPoints Number of decimal places for the amount. Must be >= 0.
	DecimalPoints int `json:"decimal_points,omitempty"`
	// DefaultValue Default amount if the field is not found in the document.
	DefaultValue float64 `json:"default_value,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Max Maximum allowed amount (inclusive).
	Max float64 `json:"max,omitempty"`
	// Min Minimum allowed amount (inclusive).
	Min float64 `json:"min,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'CURRENCY_AMOUNT' for a monetary amount field.
	Type string `json:"type"`
}
type CurrencyCodeFieldConfig struct {
	// DefaultValue Default ISO 4217 currency code if the field is not found in the document.
	DefaultValue string `json:"default_value,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'CURRENCY_CODE' for an ISO 4217 currency code field.
	Type string `json:"type"`
}
type DateFieldConfig struct {
	// AllowFutureDates Whether dates in the future are accepted.
	AllowFutureDates bool `json:"allow_future_dates,omitempty"`
	// AllowPastDates Whether dates in the past are accepted.
	AllowPastDates bool `json:"allow_past_dates,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'DATE' for a date field (YYYY-MM-DD).
	Type string `json:"type"`
}
type DatetimeFieldConfig struct {
	// AllowFutureDates Whether date-times in the future are accepted.
	AllowFutureDates bool `json:"allow_future_dates,omitempty"`
	// AllowPastDates Whether date-times in the past are accepted.
	AllowPastDates bool `json:"allow_past_dates,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'DATETIME' for a date-time field (ISO 8601).
	Type string `json:"type"`
}
type DecimalFieldConfig struct {
	// DecimalPoints Number of decimal places to round to. Must be >= 0.
	DecimalPoints int `json:"decimal_points,omitempty"`
	// DefaultValue Default decimal value if the field is not found in the document.
	DefaultValue float64 `json:"default_value,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Max Maximum allowed value (inclusive).
	Max float64 `json:"max,omitempty"`
	// Min Minimum allowed value (inclusive).
	Min float64 `json:"min,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'DECIMAL' for a decimal number field.
	Type string `json:"type"`
	// Unit Unit of measurement for context (e.g., 'EUR', 'mm').
	Unit string `json:"unit,omitempty"`
}
type DenoiseOperation struct {
	// Size Noise-reduction filter window size. Must be > 0.
	Size int `json:"size"`
	// Type Must be 'denoise'.
	Type string `json:"type"`
}
type Dimensions struct {
	// HeightInPx Height in pixels. Must be >= 1.
	HeightInPx int `json:"height_in_px"`
	// WidthInPx Width in pixels. Must be >= 1.
	WidthInPx int `json:"width_in_px"`
}
type DocumentGenerationRequest struct {
	// Document defines the document.
	Document any `json:"document"`
	// Format defines the format.
	Format string `json:"format"`
	// WebhookUrl HTTPS URL to receive results asynchronously. If provided, returns 201 immediately.
	WebhookUrl string `json:"webhook_url,omitempty"`
}
type DocumentStyles struct {
	// Grid defines the grid.
	Grid GridStyle `json:"grid"`
	// Headline defines the headline.
	Headline HeadlineStyle `json:"headline"`
	// Image defines the image.
	Image ImageStyle `json:"image"`
	// Link defines the link.
	Link LinkStyle `json:"link"`
	// List defines the list.
	List ListStyle `json:"list"`
	// Separator defines the separator.
	Separator SeparatorStyle `json:"separator"`
	// Table defines the table.
	Table TableStyle `json:"table"`
	// Text defines the text.
	Text TextStyle `json:"text"`
}
type DocumentToMarkdownSuccessResponse struct {
	// Data Converted document result with name, mime_type, markdown, optional image description, and optional nested_files.
	Data any `json:"data"`
	// Metadata defines the metadata.
	Metadata WebsiteMetadata `json:"metadata,omitempty"`
	// Success defines the success.
	Success any `json:"success"`
}
type EmailFieldConfig struct {
	// DefaultValue Default email address if the field is not found in the document.
	DefaultValue string `json:"default_value,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'EMAIL' for an email address field.
	Type string `json:"type"`
}
type EnumFieldConfig struct {
	// DefaultValue Default selected values if the field is not found in the document.
	DefaultValue []string `json:"default_value,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// MaxSelected Maximum number of values that can be selected. Must be > 0.
	MaxSelected int `json:"max_selected,omitempty"`
	// MinSelected Minimum number of values that must be selected. Must be >= 0.
	MinSelected int `json:"min_selected,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'ENUM' for a field with predefined allowed values.
	Type string `json:"type"`
	// Values List of allowed enum values the AI can choose from.
	Values []string `json:"values"`
}
type ExtendOperation struct {
	// BottomInPx Pixels to add at the bottom. Must be >= 0.
	BottomInPx int `json:"bottom_in_px"`
	// HexColor Fill color for extended area as a 6-digit hex code (e.g., '#FF0000').
	HexColor string `json:"hex_color"`
	// LeftInPx Pixels to add on the left. Must be >= 0.
	LeftInPx int `json:"left_in_px"`
	// RightInPx Pixels to add on the right. Must be >= 0.
	RightInPx int `json:"right_in_px"`
	// TopInPx Pixels to add at the top. Must be >= 0.
	TopInPx int `json:"top_in_px"`
	// Type Must be 'extend'.
	Type string `json:"type"`
}
type ExtractionSchema struct {
	// Fields defines the fields.
	Fields []any `json:"fields"`
}
type FetchOptions struct {
	// Auth Optional website authentication. Supports bearer, basic, and custom_header auth. Secret values are never returned in metadata.
	Auth any `json:"auth,omitempty"`
	// Headers Optional custom request headers sent to the target website. Unsafe hop-by-hop, cookie, and browser-controlled headers are rejected.
	Headers any `json:"headers,omitempty"`
	// Locale BCP 47 locale tag sent as Accept-Language header (e.g., "en-US", "de-DE", "fr").
	Locale string `json:"locale,omitempty"`
	// ShouldRenderJavascript When true, render the page in Chromium before extracting content.
	ShouldRenderJavascript bool `json:"should_render_javascript,omitempty"`
	// TimeoutMs Maximum fetch timeout in milliseconds. Must be between 1000 and 60000.
	TimeoutMs int `json:"timeout_ms,omitempty"`
	// UserAgent Custom User-Agent header string (1–500 characters).
	UserAgent string `json:"user_agent,omitempty"`
}
type FileInput struct {
	// Base64 Base64-encoded file content. Required when type is 'base64'.
	Base64 []byte `json:"base64,omitempty"`
	// FetchOptions defines the fetch options.
	FetchOptions FetchOptions `json:"fetch_options,omitempty"`
	// Name Optional file name including extension (e.g., 'invoice.pdf'). Used for format detection for direct files.
	Name string `json:"name,omitempty"`
	// Type Input type: 'base64' for inline Base64-encoded data, 'url' for a remote file URL.
	Type string `json:"type"`
	// Url URL to fetch the file from. Required when type is 'url'.
	Url string `json:"url,omitempty"`
}
type FlipOperation struct {
	// Type Must be 'flip'. Mirrors the image vertically (top to bottom).
	Type string `json:"type"`
}
type FlopOperation struct {
	// Type Must be 'flop'. Mirrors the image horizontally (left to right).
	Type string `json:"type"`
}
type FontDefinition struct {
	// File defines the file.
	File FileInput `json:"file"`
	// Name Font family name to reference in text layers.
	Name string `json:"name"`
	// Style Font style: 'normal' or 'italic'.
	Style string `json:"style"`
	// Weight Font weight identifier (e.g., 'regular', 'bold').
	Weight string `json:"weight"`
}
type GammaOperation struct {
	// Gamma Gamma correction factor. Must be > 0. Values < 1 darken, > 1 brighten.
	Gamma float64 `json:"gamma"`
	// Type Must be 'gamma'.
	Type string `json:"type"`
}
type GradientColorStop struct {
	// HexColor Color at this gradient stop as a 6-digit hex code.
	HexColor string `json:"hex_color"`
	// Position Position along the gradient from 0 (start) to 100 (end).
	Position float64 `json:"position"`
}
type GradientLayer struct {
	// AngleInDegrees Direction angle for linear gradients, in degrees. Defaults to 0.
	AngleInDegrees float64 `json:"angle_in_degrees,omitempty"`
	// BorderBottomLeftRadius Bottom-left corner radius in pixels.
	BorderBottomLeftRadius int `json:"border_bottom_left_radius,omitempty"`
	// BorderBottomRightRadius Bottom-right corner radius in pixels.
	BorderBottomRightRadius int `json:"border_bottom_right_radius,omitempty"`
	// BorderRadius Uniform border radius in pixels. Cannot be combined with per-corner radii.
	BorderRadius int `json:"border_radius,omitempty"`
	// BorderTopLeftRadius Top-left corner radius in pixels.
	BorderTopLeftRadius int `json:"border_top_left_radius,omitempty"`
	// BorderTopRightRadius Top-right corner radius in pixels.
	BorderTopRightRadius int `json:"border_top_right_radius,omitempty"`
	// Colors defines the colors.
	Colors []GradientColorStop `json:"colors"`
	// Dimensions defines the dimensions.
	Dimensions Dimensions `json:"dimensions"`
	// GradientType Gradient type: 'linear' or 'radial'.
	GradientType string `json:"gradient_type"`
	// Index Z-order index of this layer. Must be >= 0.
	Index int `json:"index"`
	// Opacity Layer opacity from 0 (transparent) to 100 (opaque). Defaults to 100.
	Opacity int `json:"opacity,omitempty"`
	// Position defines the position.
	Position Position `json:"position"`
	// RotationInDegrees Rotation angle in degrees, from -180 to 180.
	RotationInDegrees float64 `json:"rotation_in_degrees,omitempty"`
	// Type Must be 'gradient'.
	Type string `json:"type"`
}
type GrayscaleOperation struct {
	// Type Must be 'grayscale'. Converts the image to grayscale.
	Type string `json:"type"`
}
type GridBlock struct {
	// Columns defines the columns.
	Columns []GridColumn `json:"columns"`
	// HorizontalAlignment Horizontal alignment of grid content.
	HorizontalAlignment string `json:"horizontal_alignment,omitempty"`
	// Styles defines the styles.
	Styles GridStyleOverrides `json:"styles,omitempty"`
	// Type Must be 'grid'.
	Type string `json:"type"`
	// VerticalAlignment Vertical alignment of grid content.
	VerticalAlignment string `json:"vertical_alignment,omitempty"`
}
type GridColumn struct {
	// Blocks Content blocks rendered inside this column.
	Blocks []any `json:"blocks"`
	// ColumnSpan Number of grid columns this column spans, from 1 to 12.
	ColumnSpan int `json:"column_span"`
	// HorizontalAlignment Horizontal alignment of content within this column.
	HorizontalAlignment string `json:"horizontal_alignment,omitempty"`
	// VerticalAlignment Vertical alignment of content within this column.
	VerticalAlignment string `json:"vertical_alignment,omitempty"`
}
type GridStyle struct {
	// BackgroundColor Grid background color as a 6-digit hex code.
	BackgroundColor string `json:"background_color"`
	// BorderColor Grid border color as a 6-digit hex code.
	BorderColor string `json:"border_color"`
	// BorderWidthInPt Border width in points. Must be >= 0.
	BorderWidthInPt float64 `json:"border_width_in_pt"`
	// GapInPt Gap between columns in points. Must be >= 0.
	GapInPt float64 `json:"gap_in_pt"`
}
type GridStyleOverrides struct {
	// BackgroundColor Override grid background color.
	BackgroundColor string `json:"background_color,omitempty"`
	// BorderColor Override grid border color.
	BorderColor string `json:"border_color,omitempty"`
	// BorderWidthInPt Override border width in points.
	BorderWidthInPt float64 `json:"border_width_in_pt,omitempty"`
	// GapInPt Override gap between columns in points.
	GapInPt float64 `json:"gap_in_pt,omitempty"`
}
type HeadlineBlock struct {
	// Level Headline level from h1 (largest) to h6 (smallest).
	Level string `json:"level"`
	// Styles defines the styles.
	Styles HeadlineStyleOverrides `json:"styles,omitempty"`
	// TableOfContents defines the table of contents.
	TableOfContents HeadlineTableOfContents `json:"table_of_contents,omitempty"`
	// Text Headline text content. Must not be empty.
	Text string `json:"text"`
	// Type Must be 'headline'.
	Type string `json:"type"`
}
type HeadlineStyle struct {
	// Color Headline text color as a 6-digit hex code.
	Color string `json:"color"`
	// FontFamily Font family name for headlines.
	FontFamily string `json:"font_family"`
	// FontSizeInPt Font size in points. Must be >= 1.
	FontSizeInPt float64 `json:"font_size_in_pt"`
	// FontWeight Font weight. Defaults to 'bold'.
	FontWeight string `json:"font_weight,omitempty"`
	// IsItalic Whether headlines are italic. Defaults to false.
	IsItalic bool `json:"is_italic,omitempty"`
	// SpacingAfterInPt Space after the headline in points.
	SpacingAfterInPt float64 `json:"spacing_after_in_pt"`
	// SpacingBeforeInPt Space before the headline in points.
	SpacingBeforeInPt float64 `json:"spacing_before_in_pt"`
}
type HeadlineStyleOverrides struct {
	// Color Override headline text color.
	Color string `json:"color,omitempty"`
	// FontFamily Override font family name.
	FontFamily string `json:"font_family,omitempty"`
	// FontSizeInPt Override font size in points.
	FontSizeInPt float64 `json:"font_size_in_pt,omitempty"`
	// FontWeight Override font weight.
	FontWeight string `json:"font_weight,omitempty"`
	// IsItalic Override italic style.
	IsItalic bool `json:"is_italic,omitempty"`
	// SpacingAfterInPt Override space after the headline.
	SpacingAfterInPt float64 `json:"spacing_after_in_pt,omitempty"`
	// SpacingBeforeInPt Override space before the headline.
	SpacingBeforeInPt float64 `json:"spacing_before_in_pt,omitempty"`
}
type HeadlineTableOfContents struct {
	// IsIncluded Whether to include this headline in the table of contents.
	IsIncluded bool `json:"is_included,omitempty"`
	// LabelOverride Custom label to show in the table of contents instead of the headline text.
	LabelOverride string `json:"label_override,omitempty"`
	// LevelOverride Override the headline level in the table of contents.
	LevelOverride string `json:"level_override,omitempty"`
}
type IbanFieldConfig struct {
	// DefaultValue Default IBAN value if the field is not found in the document.
	DefaultValue string `json:"default_value,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'IBAN' for an International Bank Account Number field.
	Type string `json:"type"`
}
type ImageBlock struct {
	// Buffer Base64-encoded image content.
	Buffer []byte `json:"buffer"`
	// Fit How the image fits the specified dimensions.
	Fit string `json:"fit,omitempty"`
	// HeightInPt Image height in points. Must be >= 1.
	HeightInPt float64 `json:"height_in_pt"`
	// Styles defines the styles.
	Styles ImageStyleOverrides `json:"styles,omitempty"`
	// Type Must be 'image'.
	Type string `json:"type"`
	// WidthInPt Image width in points. Must be >= 1.
	WidthInPt float64 `json:"width_in_pt"`
}
type ImageComposition struct {
	// Dimensions defines the dimensions.
	Dimensions Dimensions `json:"dimensions"`
	// Fonts defines the fonts.
	Fonts []FontDefinition `json:"fonts,omitempty"`
	// Layers defines the layers.
	Layers []any `json:"layers"`
	// OutputFormat defines the output format.
	OutputFormat string `json:"output_format,omitempty"`
	// WebhookUrl HTTPS URL to receive results asynchronously. If provided, returns 201 immediately.
	WebhookUrl string `json:"webhook_url,omitempty"`
}
type ImageLayer struct {
	// BorderBottomLeftRadius Bottom-left corner radius in pixels.
	BorderBottomLeftRadius int `json:"border_bottom_left_radius,omitempty"`
	// BorderBottomRightRadius Bottom-right corner radius in pixels.
	BorderBottomRightRadius int `json:"border_bottom_right_radius,omitempty"`
	// BorderRadius Uniform border radius in pixels. Cannot be combined with per-corner radii.
	BorderRadius int `json:"border_radius,omitempty"`
	// BorderTopLeftRadius Top-left corner radius in pixels.
	BorderTopLeftRadius int `json:"border_top_left_radius,omitempty"`
	// BorderTopRightRadius Top-right corner radius in pixels.
	BorderTopRightRadius int `json:"border_top_right_radius,omitempty"`
	// Dimensions defines the dimensions.
	Dimensions Dimensions `json:"dimensions,omitempty"`
	// File defines the file.
	File FileInput `json:"file"`
	// Index Z-order index of this layer. Must be >= 0.
	Index int `json:"index"`
	// Opacity Layer opacity from 0 (transparent) to 100 (opaque). Defaults to 100.
	Opacity int `json:"opacity,omitempty"`
	// Position defines the position.
	Position Position `json:"position,omitempty"`
	// RotationInDegrees Rotation angle in degrees, from -180 to 180.
	RotationInDegrees float64 `json:"rotation_in_degrees,omitempty"`
	// ShouldRemoveBackground Whether to remove the image background. Defaults to false.
	ShouldRemoveBackground bool `json:"should_remove_background,omitempty"`
	// ShouldUseSmartCropping Whether to auto-crop to the most interesting region. Defaults to false.
	ShouldUseSmartCropping bool `json:"should_use_smart_cropping,omitempty"`
	// Type Must be 'image'.
	Type string `json:"type"`
}
type ImageStyle struct {
	// BorderColor Image border color as a 6-digit hex code.
	BorderColor string `json:"border_color"`
	// BorderWidthInPt Image border width in points. Must be >= 0.
	BorderWidthInPt float64 `json:"border_width_in_pt"`
}
type ImageStyleOverrides struct {
	// BorderColor Override image border color.
	BorderColor string `json:"border_color,omitempty"`
	// BorderWidthInPt Override image border width in points.
	BorderWidthInPt float64 `json:"border_width_in_pt,omitempty"`
}
type InnerBorderStyle struct {
	// Horizontal defines the horizontal.
	Horizontal BorderStyle `json:"horizontal"`
	// Vertical defines the vertical.
	Vertical BorderStyle `json:"vertical"`
}
type InnerBorderStyleOverrides struct {
	// Horizontal defines the horizontal.
	Horizontal BorderStyleOverrides `json:"horizontal,omitempty"`
	// Vertical defines the vertical.
	Vertical BorderStyleOverrides `json:"vertical,omitempty"`
}
type IntegerFieldConfig struct {
	// DefaultValue Default integer value if the field is not found in the document.
	DefaultValue int `json:"default_value,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Max Maximum allowed value (inclusive).
	Max int `json:"max,omitempty"`
	// Min Minimum allowed value (inclusive).
	Min int `json:"min,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'INTEGER' for a whole-number field.
	Type string `json:"type"`
	// Unit Unit of measurement for context (e.g., 'kg', 'pcs').
	Unit string `json:"unit,omitempty"`
}
type InvertColorsOperation struct {
	// Type Must be 'invert_colors'. Inverts all color channels.
	Type string `json:"type"`
}
type LayoutLayer struct {
	// BackgroundColor Background color as a 6-digit hex code.
	BackgroundColor string `json:"background_color,omitempty"`
	// BackgroundLayers Layers rendered behind the main content as a background.
	BackgroundLayers []any `json:"background_layers,omitempty"`
	// BorderBottomLeftRadius Bottom-left corner radius in pixels.
	BorderBottomLeftRadius int `json:"border_bottom_left_radius,omitempty"`
	// BorderBottomRightRadius Bottom-right corner radius in pixels.
	BorderBottomRightRadius int `json:"border_bottom_right_radius,omitempty"`
	// BorderRadius Uniform border radius in pixels. Cannot be combined with per-corner radii.
	BorderRadius int `json:"border_radius,omitempty"`
	// BorderTopLeftRadius Top-left corner radius in pixels.
	BorderTopLeftRadius int `json:"border_top_left_radius,omitempty"`
	// BorderTopRightRadius Top-right corner radius in pixels.
	BorderTopRightRadius int `json:"border_top_right_radius,omitempty"`
	// Dimensions defines the dimensions.
	Dimensions Dimensions `json:"dimensions,omitempty"`
	// Direction Layout direction for child layers. Defaults to 'horizontal'.
	Direction string `json:"direction,omitempty"`
	// Gap Spacing between child layers in pixels. Defaults to 0.
	Gap int `json:"gap,omitempty"`
	// HorizontalAlignment Horizontal alignment of children: 'start', 'center', or 'end'. Defaults to 'start'.
	HorizontalAlignment string `json:"horizontal_alignment,omitempty"`
	// Index Z-order index of this layer. Must be >= 0.
	Index int `json:"index"`
	// Layers Child layers arranged according to the layout direction.
	Layers []any `json:"layers"`
	// Opacity Layer opacity from 0 (transparent) to 100 (opaque). Defaults to 100.
	Opacity int `json:"opacity,omitempty"`
	// Padding Uniform padding in pixels on all sides.
	Padding int `json:"padding,omitempty"`
	// PaddingBottom Bottom padding in pixels.
	PaddingBottom int `json:"padding_bottom,omitempty"`
	// PaddingLeft Left padding in pixels.
	PaddingLeft int `json:"padding_left,omitempty"`
	// PaddingRight Right padding in pixels.
	PaddingRight int `json:"padding_right,omitempty"`
	// PaddingTop Top padding in pixels.
	PaddingTop int `json:"padding_top,omitempty"`
	// Position defines the position.
	Position Position `json:"position,omitempty"`
	// Type Must be 'layout'.
	Type string `json:"type"`
	// VerticalAlignment Vertical alignment of children: 'start', 'center', or 'end'. Defaults to 'start'.
	VerticalAlignment string `json:"vertical_alignment,omitempty"`
}
type LinkStyle struct {
	// Color Link text color as a 6-digit hex code.
	Color string `json:"color"`
	// IsUnderlined Whether links are underlined. Defaults to true.
	IsUnderlined bool `json:"is_underlined,omitempty"`
}
type ListBlock struct {
	// Items defines the items.
	Items []ListItem `json:"items"`
	// Styles defines the styles.
	Styles ListStyleOverrides `json:"styles,omitempty"`
	// Type Must be 'list'.
	Type string `json:"type"`
	// Variant List variant: 'ordered' (numbered) or 'unordered' (bulleted).
	Variant string `json:"variant"`
}
type ListItem struct {
	// Text List item text content. Must not be empty.
	Text string `json:"text"`
}
type ListStyle struct {
	// MarkerColor Bullet or number color as a 6-digit hex code.
	MarkerColor string `json:"marker_color"`
	// MarkerGapInPt Gap between marker and text in points.
	MarkerGapInPt float64 `json:"marker_gap_in_pt"`
	// TextStyle defines the text style.
	TextStyle TextStyle `json:"text_style"`
}
type ListStyleOverrides struct {
	// MarkerColor Override bullet or number color.
	MarkerColor string `json:"marker_color,omitempty"`
	// MarkerGapInPt Override gap between marker and text.
	MarkerGapInPt float64 `json:"marker_gap_in_pt,omitempty"`
	// TextStyle defines the text style.
	TextStyle TextStyleOverrides `json:"text_style,omitempty"`
}
type Margins struct {
	// BottomInPt Bottom margin in points. Must be >= 0.
	BottomInPt float64 `json:"bottom_in_pt"`
	// LeftInPt Left margin in points. Must be >= 0.
	LeftInPt float64 `json:"left_in_pt"`
	// RightInPt Right margin in points. Must be >= 0.
	RightInPt float64 `json:"right_in_pt"`
	// TopInPt Top margin in points. Must be >= 0.
	TopInPt float64 `json:"top_in_pt"`
}
type Metadata struct {
	// Author Document author name.
	Author string `json:"author,omitempty"`
	// Language BCP 47 language tag (e.g., 'en', 'de'). Defaults to 'en'.
	Language string `json:"language,omitempty"`
	// Title Document title. Must not be empty.
	Title string `json:"title"`
}
type ModulateOperation struct {
	// Brightness Brightness multiplier. 1.0 is unchanged, > 1.0 brightens, < 1.0 darkens.
	Brightness float64 `json:"brightness,omitempty"`
	// Hue Hue rotation in degrees.
	Hue float64 `json:"hue,omitempty"`
	// Saturation Saturation multiplier. 1.0 is unchanged, 0.0 is fully desaturated.
	Saturation float64 `json:"saturation,omitempty"`
	// Type Must be 'modulate'.
	Type string `json:"type"`
}
type OpacityOperation struct {
	// OpacityInPercent Opacity percentage from 0 (fully transparent) to 100 (fully opaque).
	OpacityInPercent int `json:"opacity_in_percent"`
	// Type Must be 'opacity'.
	Type string `json:"type"`
}
type OuterBorderStyle struct {
	// Bottom defines the bottom.
	Bottom BorderStyle `json:"bottom"`
	// Left defines the left.
	Left BorderStyle `json:"left"`
	// Right defines the right.
	Right BorderStyle `json:"right"`
	// Top defines the top.
	Top BorderStyle `json:"top"`
}
type OuterBorderStyleOverrides struct {
	// Bottom defines the bottom.
	Bottom BorderStyleOverrides `json:"bottom,omitempty"`
	// Left defines the left.
	Left BorderStyleOverrides `json:"left,omitempty"`
	// Right defines the right.
	Right BorderStyleOverrides `json:"right,omitempty"`
	// Top defines the top.
	Top BorderStyleOverrides `json:"top,omitempty"`
}
type Page struct {
	// BackgroundColor Page background color as a 6-digit hex code.
	BackgroundColor string `json:"background_color,omitempty"`
	// Margins defines the margins.
	Margins Margins `json:"margins"`
	// Size defines the size.
	Size PageSize `json:"size"`
}
type PageBreakBlock struct {
	// Type Must be 'page-break'. Forces a page break.
	Type string `json:"type"`
}
type PageNumberBlock struct {
	// Styles defines the styles.
	Styles TextStyleOverrides `json:"styles,omitempty"`
	// TextAlignment Text alignment for the page number.
	TextAlignment string `json:"text_alignment,omitempty"`
	// Type Must be 'page-number'. Renders the current page number.
	Type string `json:"type"`
}
type PageSize struct {
	// HeightInPt Custom page height in points. Must be >= 1.
	HeightInPt float64 `json:"height_in_pt,omitempty"`
	// Preset Standard page size preset name.
	Preset string `json:"preset,omitempty"`
	// WidthInPt Custom page width in points. Must be >= 1.
	WidthInPt float64 `json:"width_in_pt,omitempty"`
}
type ParagraphBlock struct {
	// Markdown Markdown text content. Alternative to runs.
	Markdown string `json:"markdown,omitempty"`
	// Runs defines the runs.
	Runs []ParagraphRun `json:"runs,omitempty"`
	// Styles defines the styles.
	Styles TextStyleOverrides `json:"styles,omitempty"`
	// TextAlignment Text alignment within the paragraph.
	TextAlignment string `json:"text_alignment,omitempty"`
	// Type Must be 'paragraph'.
	Type string `json:"type"`
}
type ParagraphRun struct {
	// FontWeight Font weight override for this run.
	FontWeight string `json:"font_weight,omitempty"`
	// IsItalic Whether this run is italic.
	IsItalic bool `json:"is_italic,omitempty"`
	// LinkUrl URL to make this run a clickable link.
	LinkUrl string `json:"link_url,omitempty"`
	// Text Text content of the run. Must not be empty.
	Text string `json:"text"`
}
type Position struct {
	// XInPx Horizontal position in pixels from the left edge. Must be >= 0.
	XInPx float64 `json:"x_in_px"`
	// YInPx Vertical position in pixels from the top edge. Must be >= 0.
	YInPx float64 `json:"y_in_px"`
}
type QrCodeBlock struct {
	// BgHexColor Background color as a 6-digit hex code.
	BgHexColor string `json:"bg_hex_color,omitempty"`
	// FgHexColor Foreground color as a 6-digit hex code.
	FgHexColor string `json:"fg_hex_color,omitempty"`
	// HeightInPt QR code height in points. Must be >= 1.
	HeightInPt float64 `json:"height_in_pt"`
	// Type Must be 'qr-code'.
	Type string `json:"type"`
	// Value Data to encode in the QR code.
	Value string `json:"value"`
	// WidthInPt QR code width in points. Must be >= 1.
	WidthInPt float64 `json:"width_in_pt"`
}
type QrCodeLayer struct {
	// BackgroundHexColor QR code background color as a 6-digit hex code.
	BackgroundHexColor string `json:"background_hex_color"`
	// Dimensions defines the dimensions.
	Dimensions Dimensions `json:"dimensions"`
	// ForegroundHexColor QR code foreground color as a 6-digit hex code.
	ForegroundHexColor string `json:"foreground_hex_color"`
	// Index Z-order index of this layer. Must be >= 0.
	Index int `json:"index"`
	// Opacity Layer opacity from 0 (transparent) to 100 (opaque). Defaults to 100.
	Opacity int `json:"opacity,omitempty"`
	// Position defines the position.
	Position Position `json:"position"`
	// RotationInDegrees Rotation angle in degrees, from -180 to 180.
	RotationInDegrees float64 `json:"rotation_in_degrees,omitempty"`
	// Type Must be 'qr-code'.
	Type string `json:"type"`
	// Value Data to encode in the QR code.
	Value string `json:"value"`
}
type RemoveBackgroundOperation struct {
	// BackgroundHexColor Optional replacement background color as a 6-digit hex code.
	BackgroundHexColor string `json:"background_hex_color,omitempty"`
	// Type Must be 'remove_background'.
	Type string `json:"type"`
}
type RemoveTransparencyOperation struct {
	// HexColor Background color to replace transparent areas, as a 6-digit hex code.
	HexColor string `json:"hex_color"`
	// Type Must be 'remove_transparency'.
	Type string `json:"type"`
}
type ResizeOperation struct {
	// Fit How the image fits the target dimensions.
	Fit string `json:"fit"`
	// HeightInPx Target height in pixels. Must be > 0.
	HeightInPx int `json:"height_in_px"`
	// Type Must be 'resize'.
	Type string `json:"type"`
	// WidthInPx Target width in pixels. Must be > 0.
	WidthInPx int `json:"width_in_px"`
}
type RotateOperation struct {
	// AngleInDegrees Rotation angle in degrees. Positive rotates clockwise.
	AngleInDegrees float64 `json:"angle_in_degrees"`
	// HexColor Background fill color for uncovered areas after rotation, as a 6-digit hex code.
	HexColor string `json:"hex_color,omitempty"`
	// Type Must be 'rotate'.
	Type string `json:"type"`
}
type SeparatorBlock struct {
	// Styles defines the styles.
	Styles SeparatorStyleOverrides `json:"styles,omitempty"`
	// Type Must be 'separator'. Renders a horizontal line.
	Type string `json:"type"`
}
type SeparatorStyle struct {
	// Color Separator line color as a 6-digit hex code.
	Color string `json:"color"`
	// SpacingAfterInPt Space after the separator in points.
	SpacingAfterInPt float64 `json:"spacing_after_in_pt"`
	// SpacingBeforeInPt Space before the separator in points.
	SpacingBeforeInPt float64 `json:"spacing_before_in_pt"`
	// ThicknessInPt Line thickness in points. Must be >= 0.
	ThicknessInPt float64 `json:"thickness_in_pt"`
}
type SeparatorStyleOverrides struct {
	// Color Override separator color.
	Color string `json:"color,omitempty"`
	// SpacingAfterInPt Override space after the separator.
	SpacingAfterInPt float64 `json:"spacing_after_in_pt,omitempty"`
	// SpacingBeforeInPt Override space before the separator.
	SpacingBeforeInPt float64 `json:"spacing_before_in_pt,omitempty"`
	// ThicknessInPt Override line thickness in points.
	ThicknessInPt float64 `json:"thickness_in_pt,omitempty"`
}
type SharpenOperation struct {
	// Sigma Sharpening sigma. Must be > 0. Higher values sharpen more.
	Sigma float64 `json:"sigma"`
	// Type Must be 'sharpen'.
	Type string `json:"type"`
}
type SheetGenerationRequest struct {
	// Fonts defines the fonts.
	Fonts []FontDefinition `json:"fonts,omitempty"`
	// Format defines the format.
	Format string `json:"format"`
	// Sheets defines the sheets.
	Sheets []any `json:"sheets"`
	// Styles defines the styles.
	Styles SheetStyles `json:"styles,omitempty"`
	// WebhookUrl HTTPS URL to receive results asynchronously. If provided, returns 201 immediately.
	WebhookUrl string `json:"webhook_url,omitempty"`
}
type SheetStyles struct {
	// Body defines the body.
	Body CellStyle `json:"body,omitempty"`
	// Header defines the header.
	Header CellStyle `json:"header,omitempty"`
}
type SmartCropOperation struct {
	// HeightInPx Target crop height in pixels. Must be > 0.
	HeightInPx int `json:"height_in_px"`
	// Type Must be 'smart_crop'. Crops to the most interesting region.
	Type string `json:"type"`
	// WidthInPx Target crop width in pixels. Must be > 0.
	WidthInPx int `json:"width_in_px"`
}
type SolidColorLayer struct {
	// AngledEdges defines the angled edges.
	AngledEdges []AngledEdge `json:"angled_edges,omitempty"`
	// BorderBottomLeftRadius Bottom-left corner radius in pixels.
	BorderBottomLeftRadius int `json:"border_bottom_left_radius,omitempty"`
	// BorderBottomRightRadius Bottom-right corner radius in pixels.
	BorderBottomRightRadius int `json:"border_bottom_right_radius,omitempty"`
	// BorderRadius Uniform border radius in pixels. Cannot be combined with per-corner radii.
	BorderRadius int `json:"border_radius,omitempty"`
	// BorderTopLeftRadius Top-left corner radius in pixels.
	BorderTopLeftRadius int `json:"border_top_left_radius,omitempty"`
	// BorderTopRightRadius Top-right corner radius in pixels.
	BorderTopRightRadius int `json:"border_top_right_radius,omitempty"`
	// Dimensions defines the dimensions.
	Dimensions Dimensions `json:"dimensions,omitempty"`
	// HexColor Fill color as a 6-digit hex code.
	HexColor string `json:"hex_color"`
	// Index Z-order index of this layer. Must be >= 0.
	Index int `json:"index"`
	// Opacity Layer opacity from 0 (transparent) to 100 (opaque). Defaults to 100.
	Opacity int `json:"opacity,omitempty"`
	// Position defines the position.
	Position Position `json:"position,omitempty"`
	// RotationInDegrees Rotation angle in degrees, from -180 to 180.
	RotationInDegrees float64 `json:"rotation_in_degrees,omitempty"`
	// Type Must be 'solid-color'.
	Type string `json:"type"`
}
type TableBlock struct {
	// ColumnWidthsInPercent Relative column widths in percent. Each value must be >= 1.
	ColumnWidthsInPercent []float64 `json:"column_widths_in_percent,omitempty"`
	// Header defines the header.
	Header TableRow `json:"header,omitempty"`
	// Rows defines the rows.
	Rows []TableRow `json:"rows"`
	// Styles defines the styles.
	Styles TableStyleOverrides `json:"styles,omitempty"`
	// Type Must be 'table'.
	Type string `json:"type"`
}
type TableBodyStyle struct {
	// BackgroundColor Table body background color.
	BackgroundColor string `json:"background_color"`
	// FontSizeInPt Body text font size in points. Must be >= 1.
	FontSizeInPt float64 `json:"font_size_in_pt"`
	// TextColor Table body text color.
	TextColor string `json:"text_color"`
}
type TableBodyStyleOverrides struct {
	// BackgroundColor Override body background color.
	BackgroundColor string `json:"background_color,omitempty"`
	// FontSizeInPt Override body font size in points.
	FontSizeInPt float64 `json:"font_size_in_pt,omitempty"`
	// TextColor Override body text color.
	TextColor string `json:"text_color,omitempty"`
}
type TableBorderStyle struct {
	// Inner defines the inner.
	Inner InnerBorderStyle `json:"inner"`
	// Outer defines the outer.
	Outer OuterBorderStyle `json:"outer"`
}
type TableBorderStyleOverrides struct {
	// Inner defines the inner.
	Inner InnerBorderStyleOverrides `json:"inner,omitempty"`
	// Outer defines the outer.
	Outer OuterBorderStyleOverrides `json:"outer,omitempty"`
}
type TableCell struct {
	// ColumnSpan Number of columns this cell spans. Must be >= 1.
	ColumnSpan int `json:"column_span,omitempty"`
	// HorizontalAlignment Text alignment within the cell.
	HorizontalAlignment string `json:"horizontal_alignment,omitempty"`
	// RowSpan Number of rows this cell spans. Must be >= 1.
	RowSpan int `json:"row_span,omitempty"`
	// Styles defines the styles.
	Styles TableCellStyle `json:"styles,omitempty"`
	// Text Cell text content. Must not be empty.
	Text string `json:"text"`
}
type TableCellStyle struct {
	// BackgroundColor Cell background color.
	BackgroundColor string `json:"background_color,omitempty"`
	// FontSizeInPt Cell font size in points.
	FontSizeInPt float64 `json:"font_size_in_pt,omitempty"`
	// FontWeight Cell font weight.
	FontWeight string `json:"font_weight,omitempty"`
	// IsItalic Whether cell text is italic.
	IsItalic bool `json:"is_italic,omitempty"`
	// TextColor Cell text color.
	TextColor string `json:"text_color,omitempty"`
}
type TableHeaderStyle struct {
	// BackgroundColor Header row background color.
	BackgroundColor string `json:"background_color"`
	// FontSizeInPt Header font size in points.
	FontSizeInPt float64 `json:"font_size_in_pt"`
	// FontWeight Header font weight. Defaults to 'bold'.
	FontWeight string `json:"font_weight,omitempty"`
	// TextColor Header row text color.
	TextColor string `json:"text_color"`
}
type TableHeaderStyleOverrides struct {
	// BackgroundColor Override header background color.
	BackgroundColor string `json:"background_color,omitempty"`
	// FontSizeInPt Override header font size.
	FontSizeInPt float64 `json:"font_size_in_pt,omitempty"`
	// FontWeight Override header font weight.
	FontWeight string `json:"font_weight,omitempty"`
	// TextColor Override header text color.
	TextColor string `json:"text_color,omitempty"`
}
type TableOfContentsBlock struct {
	// Leader Leader style between entry and page number.
	Leader string `json:"leader"`
	// Levels Headline levels to include in the table of contents.
	Levels []string `json:"levels"`
	// Styles defines the styles.
	Styles TextStyleOverrides `json:"styles,omitempty"`
	// TextAlignment Text alignment for table of contents entries.
	TextAlignment string `json:"text_alignment,omitempty"`
	// Type Must be 'table-of-contents'.
	Type string `json:"type"`
}
type TableRow struct {
	// Cells defines the cells.
	Cells []TableCell `json:"cells"`
}
type TableStyle struct {
	// Body defines the body.
	Body TableBodyStyle `json:"body"`
	// Border defines the border.
	Border TableBorderStyle `json:"border"`
	// Header defines the header.
	Header TableHeaderStyle `json:"header"`
}
type TableStyleOverrides struct {
	// Body defines the body.
	Body TableBodyStyleOverrides `json:"body,omitempty"`
	// Border defines the border.
	Border TableBorderStyleOverrides `json:"border,omitempty"`
	// Header defines the header.
	Header TableHeaderStyleOverrides `json:"header,omitempty"`
}
type TextFieldConfig struct {
	// DefaultValue Default value if the field is not found in the document.
	DefaultValue string `json:"default_value,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// MaxLength Maximum character length for the extracted text. Must be > 0.
	MaxLength int `json:"max_length,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'TEXT' for a single-line text field.
	Type string `json:"type"`
}
type TextLayer struct {
	// Dimensions defines the dimensions.
	Dimensions Dimensions `json:"dimensions"`
	// FontName Font family name to use for rendering.
	FontName string `json:"font_name"`
	// FontSizeInPx Font size in pixels. Must be > 0.
	FontSizeInPx int `json:"font_size_in_px"`
	// FontStyle Font style: 'normal' or 'italic'. Defaults to 'normal'.
	FontStyle string `json:"font_style,omitempty"`
	// FontWeight Font weight. Defaults to 'regular'.
	FontWeight string `json:"font_weight,omitempty"`
	// Index Z-order index of this layer. Must be >= 0.
	Index int `json:"index"`
	// IsSplittingLines Whether to wrap text across multiple lines. Defaults to true.
	IsSplittingLines bool `json:"is_splitting_lines,omitempty"`
	// Opacity Layer opacity from 0 (transparent) to 100 (opaque). Defaults to 100.
	Opacity int `json:"opacity,omitempty"`
	// ParagraphSpacingInPx Extra spacing between paragraphs in pixels.
	ParagraphSpacingInPx int `json:"paragraph_spacing_in_px,omitempty"`
	// Position defines the position.
	Position Position `json:"position"`
	// RotationInDegrees Rotation angle in degrees, from -180 to 180.
	RotationInDegrees float64 `json:"rotation_in_degrees,omitempty"`
	// ShouldAutoScale Whether to auto-scale the font to fit the layer dimensions. Defaults to false.
	ShouldAutoScale bool `json:"should_auto_scale,omitempty"`
	// Text Text content to render. Must not be empty.
	Text string `json:"text"`
	// TextAlign Horizontal text alignment within the layer. Defaults to 'left'.
	TextAlign string `json:"text_align,omitempty"`
	// TextColor Text color as a 6-digit hex code.
	TextColor string `json:"text_color"`
	// Type Must be 'text'.
	Type string `json:"type"`
	// VerticalAlign Vertical text alignment within the layer. Defaults to 'top'.
	VerticalAlign string `json:"vertical_align,omitempty"`
}
type TextStyle struct {
	// Color Text color as a 6-digit hex code.
	Color string `json:"color"`
	// FontFamily Font family name for body text.
	FontFamily string `json:"font_family"`
	// FontSizeInPt Font size in points. Must be >= 1.
	FontSizeInPt float64 `json:"font_size_in_pt"`
	// FontWeight Font weight. Defaults to 'regular'.
	FontWeight string `json:"font_weight,omitempty"`
	// IsItalic Whether text is italic. Defaults to false.
	IsItalic bool `json:"is_italic,omitempty"`
	// LineHeight Line height multiplier. Must be >= 0.5.
	LineHeight float64 `json:"line_height"`
}
type TextStyleOverrides struct {
	// Color Override text color.
	Color string `json:"color,omitempty"`
	// FontFamily Override font family name.
	FontFamily string `json:"font_family,omitempty"`
	// FontSizeInPt Override font size in points.
	FontSizeInPt float64 `json:"font_size_in_pt,omitempty"`
	// FontWeight Override font weight.
	FontWeight string `json:"font_weight,omitempty"`
	// IsItalic Override italic style.
	IsItalic bool `json:"is_italic,omitempty"`
	// LineHeight Override line height multiplier.
	LineHeight float64 `json:"line_height,omitempty"`
}
type TextareaFieldConfig struct {
	// DefaultValue Default value if the field is not found in the document.
	DefaultValue string `json:"default_value,omitempty"`
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// MaxLength Maximum character length for the extracted text. Must be > 0.
	MaxLength int `json:"max_length,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'TEXTAREA' for a multi-line text field.
	Type string `json:"type"`
}
type ThresholdOperation struct {
	// IsGrayscale Whether to convert to grayscale before applying the threshold.
	IsGrayscale bool `json:"is_grayscale"`
	// Type Must be 'threshold'.
	Type string `json:"type"`
	// Value Threshold value from 0 to 255. Pixels above become white, below become black.
	Value int `json:"value"`
}
type TimeFieldConfig struct {
	// Description Human-readable description guiding the AI on what to extract for this field.
	Description string `json:"description"`
	// IsRequired Whether this field must be present in the extraction result. Defaults to false.
	IsRequired bool `json:"is_required,omitempty"`
	// Name Unique field name used as the key in extraction results.
	Name string `json:"name"`
	// Type Must be 'TIME' for a time field (HH:MM:SS).
	Type string `json:"type"`
}
type TintOperation struct {
	// HexColor Tint color as a 6-digit hex code (e.g., '#FF0000').
	HexColor string `json:"hex_color"`
	// Type Must be 'tint'.
	Type string `json:"type"`
}
type TransformOperation = any
type TrimOperation struct {
	// Threshold Color difference threshold for trimming. Must be >= 0. Lower is stricter.
	Threshold int `json:"threshold"`
	// Type Must be 'trim'. Removes uniform borders from the image.
	Type string `json:"type"`
}
type UpscaleOperation struct {
	// Factor Upscale multiplier. Allowed values: 2, 3, or 4.
	Factor string `json:"factor"`
	// Type Must be 'upscale'.
	Type string `json:"type"`
}
type UrlFileInput struct {
	// FetchOptions defines the fetch options.
	FetchOptions FetchOptions `json:"fetch_options,omitempty"`
	// Name Optional file name including extension (e.g., 'report.pdf'). Used for format detection when fetching direct files.
	Name string `json:"name,omitempty"`
	// Type Must be 'url' for remote file input.
	Type string `json:"type"`
	// Url URL to fetch the file from.
	Url string `json:"url"`
}
type WebsiteExtractionRequest struct {
	// File defines the file.
	File UrlFileInput `json:"file"`
	// Schema defines the schema.
	Schema ExtractionSchema `json:"schema"`
	// WebhookUrl HTTPS URL to receive results asynchronously. If provided, returns 201 immediately.
	WebhookUrl string `json:"webhook_url,omitempty"`
}
type WebsiteExtractionSuccessResponse struct {
	// Data Extracted fields keyed by requested field name.
	Data any `json:"data"`
	// Metadata defines the metadata.
	Metadata WebsiteMetadata `json:"metadata"`
	// Success defines the success.
	Success any `json:"success"`
}
type WebsiteMetadata struct {
	// Article defines the article.
	Article any `json:"article,omitempty"`
	// Assets defines the assets.
	Assets []any `json:"assets,omitempty"`
	// BodySize defines the body size.
	BodySize int `json:"body_size,omitempty"`
	// ContentType defines the content type.
	ContentType string `json:"content_type,omitempty"`
	// DublinCore defines the dublin core.
	DublinCore any `json:"dublin_core,omitempty"`
	// Favicons defines the favicons.
	Favicons []any `json:"favicons,omitempty"`
	// Feeds defines the feeds.
	Feeds []any `json:"feeds,omitempty"`
	// FinalUrl defines the final url.
	FinalUrl string `json:"final_url"`
	// Headings defines the headings.
	Headings []any `json:"headings,omitempty"`
	// Hreflang defines the hreflang.
	Hreflang []any `json:"hreflang,omitempty"`
	// Images defines the images.
	Images []any `json:"images,omitempty"`
	// JsonLd defines the json ld.
	JsonLd []any `json:"json_ld,omitempty"`
	// Links defines the links.
	Links []any `json:"links,omitempty"`
	// MetaRefreshUrl defines the meta refresh url.
	MetaRefreshUrl string `json:"meta_refresh_url,omitempty"`
	// NofollowDetected defines the nofollow detected.
	NofollowDetected bool `json:"nofollow_detected,omitempty"`
	// NoindexDetected defines the noindex detected.
	NoindexDetected bool `json:"noindex_detected,omitempty"`
	// OpenGraph defines the open graph.
	OpenGraph any `json:"open_graph,omitempty"`
	// Page defines the page.
	Page any `json:"page,omitempty"`
	// PageCount defines the page count.
	PageCount int `json:"page_count"`
	// RequestedUrl defines the requested url.
	RequestedUrl string `json:"requested_url"`
	// ResponseMeta defines the response meta.
	ResponseMeta any `json:"response_meta,omitempty"`
	// StatusCode defines the status code.
	StatusCode int `json:"status_code,omitempty"`
	// Twitter defines the twitter.
	Twitter any `json:"twitter,omitempty"`
	// WordCount defines the word count.
	WordCount int `json:"word_count,omitempty"`
}
type ExtractDocumentRequest struct {
	// Files defines the files.
	Files []FileInput `json:"files"`
	// Schema defines the schema.
	Schema ExtractionSchema `json:"schema"`
	// WebhookUrl HTTPS URL to receive results asynchronously. If provided, returns 201 immediately.
	WebhookUrl string `json:"webhook_url,omitempty"`
}

type ExtractDocumentAsyncRequest struct {
	// Files defines the files.
	Files []FileInput `json:"files"`
	// Schema defines the schema.
	Schema ExtractionSchema `json:"schema"`
	// WebhookURL Webhook URL for async processing.
	WebhookURL string `json:"webhook_url"`
}

type GenerateDocumentRequest struct {
	// Document defines the document.
	Document any `json:"document"`
	// Format defines the format.
	Format string `json:"format"`
	// WebhookUrl HTTPS URL to receive results asynchronously. If provided, returns 201 immediately.
	WebhookUrl string `json:"webhook_url,omitempty"`
}

type GenerateDocumentAsyncRequest struct {
	// Document defines the document.
	Document any `json:"document"`
	// Format defines the format.
	Format string `json:"format"`
	// WebhookURL Webhook URL for async processing.
	WebhookURL string `json:"webhook_url"`
}

type ConvertDocumentToMarkdownRequest struct {
	// File defines the file.
	File FileInput `json:"file"`
	// WebhookUrl HTTPS URL to receive results asynchronously. If provided, returns 201 immediately.
	WebhookUrl string `json:"webhook_url,omitempty"`
}

type ConvertDocumentToMarkdownAsyncRequest struct {
	// File defines the file.
	File FileInput `json:"file"`
	// WebhookURL Webhook URL for async processing.
	WebhookURL string `json:"webhook_url"`
}

type GenerateImageRequest struct {
	// Dimensions defines the dimensions.
	Dimensions Dimensions `json:"dimensions"`
	// Fonts defines the fonts.
	Fonts []FontDefinition `json:"fonts,omitempty"`
	// Layers defines the layers.
	Layers []any `json:"layers"`
	// OutputFormat defines the output format.
	OutputFormat string `json:"output_format,omitempty"`
	// WebhookUrl HTTPS URL to receive results asynchronously. If provided, returns 201 immediately.
	WebhookUrl string `json:"webhook_url,omitempty"`
}

type GenerateImageAsyncRequest struct {
	// Dimensions defines the dimensions.
	Dimensions Dimensions `json:"dimensions"`
	// Fonts defines the fonts.
	Fonts []FontDefinition `json:"fonts,omitempty"`
	// Layers defines the layers.
	Layers []any `json:"layers"`
	// OutputFormat defines the output format.
	OutputFormat string `json:"output_format,omitempty"`
	// WebhookURL Webhook URL for async processing.
	WebhookURL string `json:"webhook_url"`
}

type TransformImageRequest struct {
	// File defines the file.
	File FileInput `json:"file"`
	// Operations defines the operations.
	Operations []TransformOperation `json:"operations"`
	// WebhookUrl HTTPS URL to receive results asynchronously. If provided, returns 201 immediately.
	WebhookUrl string `json:"webhook_url,omitempty"`
}

type TransformImageAsyncRequest struct {
	// File defines the file.
	File FileInput `json:"file"`
	// Operations defines the operations.
	Operations []TransformOperation `json:"operations"`
	// WebhookURL Webhook URL for async processing.
	WebhookURL string `json:"webhook_url"`
}

type GenerateSheetRequest struct {
	// Fonts defines the fonts.
	Fonts []FontDefinition `json:"fonts,omitempty"`
	// Format defines the format.
	Format string `json:"format"`
	// Sheets defines the sheets.
	Sheets []any `json:"sheets"`
	// Styles defines the styles.
	Styles SheetStyles `json:"styles,omitempty"`
	// WebhookUrl HTTPS URL to receive results asynchronously. If provided, returns 201 immediately.
	WebhookUrl string `json:"webhook_url,omitempty"`
}

type GenerateSheetAsyncRequest struct {
	// Fonts defines the fonts.
	Fonts []FontDefinition `json:"fonts,omitempty"`
	// Format defines the format.
	Format string `json:"format"`
	// Sheets defines the sheets.
	Sheets []any `json:"sheets"`
	// Styles defines the styles.
	Styles SheetStyles `json:"styles,omitempty"`
	// WebhookURL Webhook URL for async processing.
	WebhookURL string `json:"webhook_url"`
}

type ExtractWebsiteRequest struct {
	// File defines the file.
	File UrlFileInput `json:"file"`
	// Schema defines the schema.
	Schema ExtractionSchema `json:"schema"`
	// WebhookUrl HTTPS URL to receive results asynchronously. If provided, returns 201 immediately.
	WebhookUrl string `json:"webhook_url,omitempty"`
}

type ExtractWebsiteAsyncRequest struct {
	// File defines the file.
	File UrlFileInput `json:"file"`
	// Schema defines the schema.
	Schema ExtractionSchema `json:"schema"`
	// WebhookURL Webhook URL for async processing.
	WebhookURL string `json:"webhook_url"`
}
