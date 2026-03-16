package iterationlayer

// ── Marker Interfaces ──────────────────────────────────────────────────────
// These interfaces enable type-safe polymorphic slices in Go.
// Each variant struct implements the interface via the embedded marker method.

type FileInput interface {
  isFileInput()
}

type FieldConfig interface {
  isFieldConfig()
}

type TransformOperation interface {
  isTransformOperation()
}

type Layer interface {
  isLayer()
}

type ContentBlock interface {
  isContentBlock()
}

type HeaderFooterBlock interface {
  isHeaderFooterBlock()
}

// ── File Input ──────────────────────────────────────────────────────────────

type FileInputBase64 struct {
  Type   string `json:"type"`
  Name   string `json:"name"`
  Base64 string `json:"base64"`
}

func (FileInputBase64) isFileInput() {}

func NewFileFromBase64(name string, base64Content string) FileInputBase64 {
  return FileInputBase64{Type: "base64", Name: name, Base64: base64Content}
}

type FileInputURL struct {
  Type string `json:"type"`
  Name string `json:"name"`
  URL  string `json:"url"`
}

func (FileInputURL) isFileInput() {}

func NewFileFromURL(name string, url string) FileInputURL {
  return FileInputURL{Type: "url", Name: name, URL: url}
}

// ── Responses ──────────────────────────────────────────────────────────────

type ExtractionFieldResult struct {
  Value      any      `json:"value"`
  Confidence float64  `json:"confidence"`
  Citations  []string `json:"citations"`
  Source     string   `json:"source"`
  Type       string   `json:"type"`
}

type ExtractionResult map[string]ExtractionFieldResult

type BinaryResult struct {
  Buffer   string `json:"buffer"`
  MimeType string `json:"mime_type"`
}

type AsyncResult struct {
  Message string `json:"message"`
}

// ── Document Extraction ────────────────────────────────────────────────────

type TextFieldConfig struct {
  Type         string `json:"type"`
  Name         string `json:"name"`
  Description  string `json:"description"`
  IsRequired   *bool  `json:"is_required,omitempty"`
  MaxLength    *int   `json:"max_length,omitempty"`
  DefaultValue string `json:"default_value,omitempty"`
}

func (TextFieldConfig) isFieldConfig() {}

func NewTextFieldConfig(name string, description string) TextFieldConfig {
  return TextFieldConfig{Type: "TEXT", Name: name, Description: description}
}

type TextareaFieldConfig struct {
  Type         string `json:"type"`
  Name         string `json:"name"`
  Description  string `json:"description"`
  IsRequired   *bool  `json:"is_required,omitempty"`
  MaxLength    *int   `json:"max_length,omitempty"`
  DefaultValue string `json:"default_value,omitempty"`
}

func (TextareaFieldConfig) isFieldConfig() {}

func NewTextareaFieldConfig(name string, description string) TextareaFieldConfig {
  return TextareaFieldConfig{Type: "TEXTAREA", Name: name, Description: description}
}

type IntegerFieldConfig struct {
  Type         string `json:"type"`
  Name         string `json:"name"`
  Description  string `json:"description"`
  IsRequired   *bool  `json:"is_required,omitempty"`
  Min          *int   `json:"min,omitempty"`
  Max          *int   `json:"max,omitempty"`
  Unit         string `json:"unit,omitempty"`
  DefaultValue *int   `json:"default_value,omitempty"`
}

func (IntegerFieldConfig) isFieldConfig() {}

func NewIntegerFieldConfig(name string, description string) IntegerFieldConfig {
  return IntegerFieldConfig{Type: "INTEGER", Name: name, Description: description}
}

type DecimalFieldConfig struct {
  Type          string   `json:"type"`
  Name          string   `json:"name"`
  Description   string   `json:"description"`
  IsRequired    *bool    `json:"is_required,omitempty"`
  Min           *float64 `json:"min,omitempty"`
  Max           *float64 `json:"max,omitempty"`
  DecimalPoints *int     `json:"decimal_points,omitempty"`
  Unit          string   `json:"unit,omitempty"`
  DefaultValue  *float64 `json:"default_value,omitempty"`
}

func (DecimalFieldConfig) isFieldConfig() {}

func NewDecimalFieldConfig(name string, description string) DecimalFieldConfig {
  return DecimalFieldConfig{Type: "DECIMAL", Name: name, Description: description}
}

type DateFieldConfig struct {
  Type             string `json:"type"`
  Name             string `json:"name"`
  Description      string `json:"description"`
  IsRequired       *bool  `json:"is_required,omitempty"`
  AllowFutureDates *bool  `json:"allow_future_dates,omitempty"`
  AllowPastDates   *bool  `json:"allow_past_dates,omitempty"`
}

func (DateFieldConfig) isFieldConfig() {}

func NewDateFieldConfig(name string, description string) DateFieldConfig {
  return DateFieldConfig{Type: "DATE", Name: name, Description: description}
}

type DatetimeFieldConfig struct {
  Type             string `json:"type"`
  Name             string `json:"name"`
  Description      string `json:"description"`
  IsRequired       *bool  `json:"is_required,omitempty"`
  AllowFutureDates *bool  `json:"allow_future_dates,omitempty"`
  AllowPastDates   *bool  `json:"allow_past_dates,omitempty"`
}

func (DatetimeFieldConfig) isFieldConfig() {}

func NewDatetimeFieldConfig(name string, description string) DatetimeFieldConfig {
  return DatetimeFieldConfig{Type: "DATETIME", Name: name, Description: description}
}

type TimeFieldConfig struct {
  Type        string `json:"type"`
  Name        string `json:"name"`
  Description string `json:"description"`
  IsRequired  *bool  `json:"is_required,omitempty"`
}

func (TimeFieldConfig) isFieldConfig() {}

func NewTimeFieldConfig(name string, description string) TimeFieldConfig {
  return TimeFieldConfig{Type: "TIME", Name: name, Description: description}
}

type EnumFieldConfig struct {
  Type         string   `json:"type"`
  Name         string   `json:"name"`
  Description  string   `json:"description"`
  Values       []string `json:"values"`
  IsRequired   *bool    `json:"is_required,omitempty"`
  MinSelected  *int     `json:"min_selected,omitempty"`
  MaxSelected  *int     `json:"max_selected,omitempty"`
  DefaultValue []string `json:"default_value,omitempty"`
}

func (EnumFieldConfig) isFieldConfig() {}

func NewEnumFieldConfig(name string, description string, values []string) EnumFieldConfig {
  return EnumFieldConfig{Type: "ENUM", Name: name, Description: description, Values: values}
}

type BooleanFieldConfig struct {
  Type         string `json:"type"`
  Name         string `json:"name"`
  Description  string `json:"description"`
  IsRequired   *bool  `json:"is_required,omitempty"`
  DefaultValue *bool  `json:"default_value,omitempty"`
}

func (BooleanFieldConfig) isFieldConfig() {}

func NewBooleanFieldConfig(name string, description string) BooleanFieldConfig {
  return BooleanFieldConfig{Type: "BOOLEAN", Name: name, Description: description}
}

type EmailFieldConfig struct {
  Type         string `json:"type"`
  Name         string `json:"name"`
  Description  string `json:"description"`
  IsRequired   *bool  `json:"is_required,omitempty"`
  DefaultValue string `json:"default_value,omitempty"`
}

func (EmailFieldConfig) isFieldConfig() {}

func NewEmailFieldConfig(name string, description string) EmailFieldConfig {
  return EmailFieldConfig{Type: "EMAIL", Name: name, Description: description}
}

type IbanFieldConfig struct {
  Type         string `json:"type"`
  Name         string `json:"name"`
  Description  string `json:"description"`
  IsRequired   *bool  `json:"is_required,omitempty"`
  DefaultValue string `json:"default_value,omitempty"`
}

func (IbanFieldConfig) isFieldConfig() {}

func NewIbanFieldConfig(name string, description string) IbanFieldConfig {
  return IbanFieldConfig{Type: "IBAN", Name: name, Description: description}
}

type CountryFieldConfig struct {
  Type         string `json:"type"`
  Name         string `json:"name"`
  Description  string `json:"description"`
  IsRequired   *bool  `json:"is_required,omitempty"`
  DefaultValue string `json:"default_value,omitempty"`
}

func (CountryFieldConfig) isFieldConfig() {}

func NewCountryFieldConfig(name string, description string) CountryFieldConfig {
  return CountryFieldConfig{Type: "COUNTRY", Name: name, Description: description}
}

type CurrencyCodeFieldConfig struct {
  Type         string `json:"type"`
  Name         string `json:"name"`
  Description  string `json:"description"`
  IsRequired   *bool  `json:"is_required,omitempty"`
  DefaultValue string `json:"default_value,omitempty"`
}

func (CurrencyCodeFieldConfig) isFieldConfig() {}

func NewCurrencyCodeFieldConfig(name string, description string) CurrencyCodeFieldConfig {
  return CurrencyCodeFieldConfig{Type: "CURRENCY_CODE", Name: name, Description: description}
}

type CurrencyAmountFieldConfig struct {
  Type          string   `json:"type"`
  Name          string   `json:"name"`
  Description   string   `json:"description"`
  IsRequired    *bool    `json:"is_required,omitempty"`
  Min           *float64 `json:"min,omitempty"`
  Max           *float64 `json:"max,omitempty"`
  DecimalPoints *int     `json:"decimal_points,omitempty"`
  DefaultValue  *float64 `json:"default_value,omitempty"`
}

func (CurrencyAmountFieldConfig) isFieldConfig() {}

func NewCurrencyAmountFieldConfig(name string, description string) CurrencyAmountFieldConfig {
  return CurrencyAmountFieldConfig{Type: "CURRENCY_AMOUNT", Name: name, Description: description}
}

type AddressFieldConfig struct {
  Type                string   `json:"type"`
  Name                string   `json:"name"`
  Description         string   `json:"description"`
  IsRequired          *bool    `json:"is_required,omitempty"`
  AllowedCountryCodes []string `json:"allowed_country_codes,omitempty"`
}

func (AddressFieldConfig) isFieldConfig() {}

func NewAddressFieldConfig(name string, description string) AddressFieldConfig {
  return AddressFieldConfig{Type: "ADDRESS", Name: name, Description: description}
}

type ItemSchema struct {
  Fields []FieldConfig `json:"fields"`
}

type ArrayFieldConfig struct {
  Type        string     `json:"type"`
  Name        string     `json:"name"`
  Description string     `json:"description"`
  ItemSchema  ItemSchema `json:"item_schema"`
  IsRequired  *bool      `json:"is_required,omitempty"`
}

func (ArrayFieldConfig) isFieldConfig() {}

func NewArrayFieldConfig(name string, description string, itemFields []FieldConfig) ArrayFieldConfig {
  return ArrayFieldConfig{Type: "ARRAY", Name: name, Description: description, ItemSchema: ItemSchema{Fields: itemFields}}
}

type CalculatedFieldConfig struct {
  Type             string   `json:"type"`
  Name             string   `json:"name"`
  Description      string   `json:"description"`
  Operation        string   `json:"operation"`
  SourceFieldNames []string `json:"source_field_names"`
  IsRequired       *bool    `json:"is_required,omitempty"`
  Unit             string   `json:"unit,omitempty"`
}

func (CalculatedFieldConfig) isFieldConfig() {}

func NewCalculatedFieldConfig(name string, description string, operation string, sourceFieldNames []string) CalculatedFieldConfig {
  return CalculatedFieldConfig{Type: "CALCULATED", Name: name, Description: description, Operation: operation, SourceFieldNames: sourceFieldNames}
}

// ── Image Transformation ───────────────────────────────────────────────────

type ResizeOperation struct {
  Type       string `json:"type"`
  WidthInPx  int    `json:"width_in_px"`
  HeightInPx int    `json:"height_in_px"`
  Fit        string `json:"fit"`
}

func (ResizeOperation) isTransformOperation() {}

func NewResizeOperation(widthInPx int, heightInPx int, fit string) ResizeOperation {
  return ResizeOperation{Type: "resize", WidthInPx: widthInPx, HeightInPx: heightInPx, Fit: fit}
}

type CropOperation struct {
  Type       string `json:"type"`
  LeftInPx   int    `json:"left_in_px"`
  TopInPx    int    `json:"top_in_px"`
  WidthInPx  int    `json:"width_in_px"`
  HeightInPx int    `json:"height_in_px"`
}

func (CropOperation) isTransformOperation() {}

func NewCropOperation(leftInPx int, topInPx int, widthInPx int, heightInPx int) CropOperation {
  return CropOperation{Type: "crop", LeftInPx: leftInPx, TopInPx: topInPx, WidthInPx: widthInPx, HeightInPx: heightInPx}
}

type ExtendOperation struct {
  Type       string `json:"type"`
  TopInPx    int    `json:"top_in_px"`
  BottomInPx int    `json:"bottom_in_px"`
  LeftInPx   int    `json:"left_in_px"`
  RightInPx  int    `json:"right_in_px"`
  HexColor   string `json:"hex_color"`
}

func (ExtendOperation) isTransformOperation() {}

func NewExtendOperation(topInPx int, bottomInPx int, leftInPx int, rightInPx int, hexColor string) ExtendOperation {
  return ExtendOperation{Type: "extend", TopInPx: topInPx, BottomInPx: bottomInPx, LeftInPx: leftInPx, RightInPx: rightInPx, HexColor: hexColor}
}

type TrimOperation struct {
  Type      string `json:"type"`
  Threshold int    `json:"threshold"`
}

func (TrimOperation) isTransformOperation() {}

func NewTrimOperation(threshold int) TrimOperation {
  return TrimOperation{Type: "trim", Threshold: threshold}
}

type RotateOperation struct {
  Type           string  `json:"type"`
  AngleInDegrees float64 `json:"angle_in_degrees"`
  HexColor       string  `json:"hex_color,omitempty"`
}

func (RotateOperation) isTransformOperation() {}

func NewRotateOperation(angleInDegrees float64) RotateOperation {
  return RotateOperation{Type: "rotate", AngleInDegrees: angleInDegrees}
}

type FlipOperation struct {
  Type string `json:"type"`
}

func (FlipOperation) isTransformOperation() {}

func NewFlipOperation() FlipOperation {
  return FlipOperation{Type: "flip"}
}

type FlopOperation struct {
  Type string `json:"type"`
}

func (FlopOperation) isTransformOperation() {}

func NewFlopOperation() FlopOperation {
  return FlopOperation{Type: "flop"}
}

type BlurOperation struct {
  Type  string  `json:"type"`
  Sigma float64 `json:"sigma"`
}

func (BlurOperation) isTransformOperation() {}

func NewBlurOperation(sigma float64) BlurOperation {
  return BlurOperation{Type: "blur", Sigma: sigma}
}

type SharpenOperation struct {
  Type  string  `json:"type"`
  Sigma float64 `json:"sigma"`
}

func (SharpenOperation) isTransformOperation() {}

func NewSharpenOperation(sigma float64) SharpenOperation {
  return SharpenOperation{Type: "sharpen", Sigma: sigma}
}

type ModulateOperation struct {
  Type       string   `json:"type"`
  Brightness *float64 `json:"brightness,omitempty"`
  Saturation *float64 `json:"saturation,omitempty"`
  Hue        *float64 `json:"hue,omitempty"`
}

func (ModulateOperation) isTransformOperation() {}

func NewModulateOperation() ModulateOperation {
  return ModulateOperation{Type: "modulate"}
}

type TintOperation struct {
  Type     string `json:"type"`
  HexColor string `json:"hex_color"`
}

func (TintOperation) isTransformOperation() {}

func NewTintOperation(hexColor string) TintOperation {
  return TintOperation{Type: "tint", HexColor: hexColor}
}

type GrayscaleOperation struct {
  Type string `json:"type"`
}

func (GrayscaleOperation) isTransformOperation() {}

func NewGrayscaleOperation() GrayscaleOperation {
  return GrayscaleOperation{Type: "grayscale"}
}

type InvertColorsOperation struct {
  Type string `json:"type"`
}

func (InvertColorsOperation) isTransformOperation() {}

func NewInvertColorsOperation() InvertColorsOperation {
  return InvertColorsOperation{Type: "invert_colors"}
}

type AutoContrastOperation struct {
  Type string `json:"type"`
}

func (AutoContrastOperation) isTransformOperation() {}

func NewAutoContrastOperation() AutoContrastOperation {
  return AutoContrastOperation{Type: "auto_contrast"}
}

type GammaOperation struct {
  Type  string  `json:"type"`
  Gamma float64 `json:"gamma"`
}

func (GammaOperation) isTransformOperation() {}

func NewGammaOperation(gamma float64) GammaOperation {
  return GammaOperation{Type: "gamma", Gamma: gamma}
}

type RemoveTransparencyOperation struct {
  Type     string `json:"type"`
  HexColor string `json:"hex_color"`
}

func (RemoveTransparencyOperation) isTransformOperation() {}

func NewRemoveTransparencyOperation(hexColor string) RemoveTransparencyOperation {
  return RemoveTransparencyOperation{Type: "remove_transparency", HexColor: hexColor}
}

type ThresholdOperation struct {
  Type        string `json:"type"`
  Value       int    `json:"value"`
  IsGrayscale bool   `json:"is_grayscale"`
}

func (ThresholdOperation) isTransformOperation() {}

func NewThresholdOperation(value int, isGrayscale bool) ThresholdOperation {
  return ThresholdOperation{Type: "threshold", Value: value, IsGrayscale: isGrayscale}
}

type DenoiseOperation struct {
  Type string `json:"type"`
  Size int    `json:"size"`
}

func (DenoiseOperation) isTransformOperation() {}

func NewDenoiseOperation(size int) DenoiseOperation {
  return DenoiseOperation{Type: "denoise", Size: size}
}

type OpacityOperation struct {
  Type             string `json:"type"`
  OpacityInPercent int    `json:"opacity_in_percent"`
}

func (OpacityOperation) isTransformOperation() {}

func NewOpacityOperation(opacityInPercent int) OpacityOperation {
  return OpacityOperation{Type: "opacity", OpacityInPercent: opacityInPercent}
}

type ConvertOperation struct {
  Type    string `json:"type"`
  Format  string `json:"format"`
  Quality *int   `json:"quality,omitempty"`
}

func (ConvertOperation) isTransformOperation() {}

func NewConvertOperation(format string) ConvertOperation {
  return ConvertOperation{Type: "convert", Format: format}
}

type UpscaleOperation struct {
  Type   string `json:"type"`
  Factor int    `json:"factor"`
}

func (UpscaleOperation) isTransformOperation() {}

func NewUpscaleOperation(factor int) UpscaleOperation {
  return UpscaleOperation{Type: "upscale", Factor: factor}
}

type SmartCropOperation struct {
  Type       string `json:"type"`
  WidthInPx  int    `json:"width_in_px"`
  HeightInPx int    `json:"height_in_px"`
}

func (SmartCropOperation) isTransformOperation() {}

func NewSmartCropOperation(widthInPx int, heightInPx int) SmartCropOperation {
  return SmartCropOperation{Type: "smart_crop", WidthInPx: widthInPx, HeightInPx: heightInPx}
}

type CompressToSizeOperation struct {
  Type               string `json:"type"`
  MaxFileSizeInBytes int    `json:"max_file_size_in_bytes"`
}

func (CompressToSizeOperation) isTransformOperation() {}

func NewCompressToSizeOperation(maxFileSizeInBytes int) CompressToSizeOperation {
  return CompressToSizeOperation{Type: "compress_to_size", MaxFileSizeInBytes: maxFileSizeInBytes}
}

type RemoveBackgroundOperation struct {
  Type               string `json:"type"`
  BackgroundHexColor string `json:"background_hex_color,omitempty"`
}

func (RemoveBackgroundOperation) isTransformOperation() {}

func NewRemoveBackgroundOperation() RemoveBackgroundOperation {
  return RemoveBackgroundOperation{Type: "remove_background"}
}

// ── Image Generation ───────────────────────────────────────────────────────

type Dimensions struct {
  WidthInPx  int `json:"width_in_px"`
  HeightInPx int `json:"height_in_px"`
}

type Position struct {
  XInPx int `json:"x_in_px"`
  YInPx int `json:"y_in_px"`
}

type AngledEdge struct {
  Edge           string  `json:"edge"`
  AngleInDegrees float64 `json:"angle_in_degrees"`
}

type GradientColorStop struct {
  HexColor string  `json:"hex_color"`
  Position float64 `json:"position"`
}

type ImageFontDefinition struct {
  Name   string    `json:"name"`
  Weight string    `json:"weight"`
  Style  string    `json:"style"`
  File   FileInput `json:"file"`
}

type SolidColorLayer struct {
  Type     string `json:"type"`
  Index    int    `json:"index"`
  HexColor string `json:"hex_color"`
  Opacity  *int   `json:"opacity,omitempty"`
}

func (SolidColorLayer) isLayer() {}

func NewSolidColorLayer(index int, hexColor string) SolidColorLayer {
  return SolidColorLayer{Type: "solid-color", Index: index, HexColor: hexColor}
}

type SolidColorLayer struct {
  Type              string       `json:"type"`
  Index             int          `json:"index"`
  HexColor          string       `json:"hex_color"`
  Position          Position     `json:"position"`
  Dimensions        Dimensions   `json:"dimensions"`
  RotationInDegrees *float64     `json:"rotation_in_degrees,omitempty"`
  Opacity           *int         `json:"opacity,omitempty"`
  AngledEdges       []AngledEdge `json:"angled_edges,omitempty"`
}

func (SolidColorLayer) isLayer() {}

func NewSolidColorLayer(index int, hexColor string, position Position, dimensions Dimensions) SolidColorLayer {
  return SolidColorLayer{Type: "solid-color", Index: index, HexColor: hexColor, Position: position, Dimensions: dimensions}
}

type TextLayer struct {
  Type                 string     `json:"type"`
  Index                int        `json:"index"`
  Text                 string     `json:"text"`
  FontName             string     `json:"font_name"`
  FontSizeInPx         int        `json:"font_size_in_px"`
  TextColor            string     `json:"text_color"`
  Position             Position   `json:"position"`
  Dimensions           Dimensions `json:"dimensions"`
  FontWeight           string     `json:"font_weight,omitempty"`
  FontStyle            string     `json:"font_style,omitempty"`
  TextAlign            string     `json:"text_align,omitempty"`
  VerticalAlign        string     `json:"vertical_align,omitempty"`
  IsSplittingLines     *bool      `json:"is_splitting_lines,omitempty"`
  ParagraphSpacingInPx *int       `json:"paragraph_spacing_in_px,omitempty"`
  ShouldAutoScale      *bool      `json:"should_auto_scale,omitempty"`
  RotationInDegrees    *float64   `json:"rotation_in_degrees,omitempty"`
  Opacity              *int       `json:"opacity,omitempty"`
}

func (TextLayer) isLayer() {}

func NewTextLayer(index int, text string, fontName string, fontSizeInPx int, textColor string, position Position, dimensions Dimensions) TextLayer {
  return TextLayer{Type: "text", Index: index, Text: text, FontName: fontName, FontSizeInPx: fontSizeInPx, TextColor: textColor, Position: position, Dimensions: dimensions}
}

type ImageLayer struct {
  Type                   string     `json:"type"`
  Index                  int        `json:"index"`
  File                   FileInput  `json:"file"`
  Position               Position   `json:"position"`
  Dimensions             Dimensions `json:"dimensions"`
  RotationInDegrees      *float64   `json:"rotation_in_degrees,omitempty"`
  Opacity                *int       `json:"opacity,omitempty"`
  ShouldUseSmartCropping *bool      `json:"should_use_smart_cropping,omitempty"`
  ShouldRemoveBackground *bool      `json:"should_remove_background,omitempty"`
}

func (ImageLayer) isLayer() {}

func NewImageLayer(index int, file FileInput, position Position, dimensions Dimensions) ImageLayer {
  return ImageLayer{Type: "image", Index: index, File: file, Position: position, Dimensions: dimensions}
}

type ImageLayer struct {
  Type                   string    `json:"type"`
  Index                  int       `json:"index"`
  File                   FileInput `json:"file"`
  Opacity                *int      `json:"opacity,omitempty"`
  ShouldUseSmartCropping *bool     `json:"should_use_smart_cropping,omitempty"`
}

func (ImageLayer) isLayer() {}

func NewImageLayer(index int, file FileInput) ImageLayer {
  return ImageLayer{Type: "image", Index: index, File: file}
}

type QrCodeLayer struct {
  Type               string     `json:"type"`
  Index              int        `json:"index"`
  Value              string     `json:"value"`
  ForegroundHexColor string     `json:"foreground_hex_color"`
  BackgroundHexColor string     `json:"background_hex_color"`
  Position           Position   `json:"position"`
  Dimensions         Dimensions `json:"dimensions"`
  RotationInDegrees  *float64   `json:"rotation_in_degrees,omitempty"`
  Opacity            *int       `json:"opacity,omitempty"`
}

func (QrCodeLayer) isLayer() {}

func NewQrCodeLayer(index int, value string, fgColor string, bgColor string, position Position, dimensions Dimensions) QrCodeLayer {
  return QrCodeLayer{Type: "qr-code", Index: index, Value: value, ForegroundHexColor: fgColor, BackgroundHexColor: bgColor, Position: position, Dimensions: dimensions}
}

type BarcodeLayer struct {
  Type               string     `json:"type"`
  Index              int        `json:"index"`
  Value              string     `json:"value"`
  Format             string     `json:"format"`
  ForegroundHexColor string     `json:"foreground_hex_color"`
  BackgroundHexColor string     `json:"background_hex_color"`
  Position           Position   `json:"position"`
  Dimensions         Dimensions `json:"dimensions"`
  RotationInDegrees  *float64   `json:"rotation_in_degrees,omitempty"`
  Opacity            *int       `json:"opacity,omitempty"`
}

func (BarcodeLayer) isLayer() {}

func NewBarcodeLayer(index int, value string, format string, fgColor string, bgColor string, position Position, dimensions Dimensions) BarcodeLayer {
  return BarcodeLayer{Type: "barcode", Index: index, Value: value, Format: format, ForegroundHexColor: fgColor, BackgroundHexColor: bgColor, Position: position, Dimensions: dimensions}
}

type GradientLayer struct {
  Type              string              `json:"type"`
  Index             int                 `json:"index"`
  GradientType      string              `json:"gradient_type"`
  Colors            []GradientColorStop `json:"colors"`
  Position          Position            `json:"position"`
  Dimensions        Dimensions          `json:"dimensions"`
  AngleInDegrees    *float64            `json:"angle_in_degrees,omitempty"`
  RotationInDegrees *float64            `json:"rotation_in_degrees,omitempty"`
  Opacity           *int                `json:"opacity,omitempty"`
}

func (GradientLayer) isLayer() {}

func NewGradientLayer(index int, gradientType string, colors []GradientColorStop, position Position, dimensions Dimensions) GradientLayer {
  return GradientLayer{Type: "gradient", Index: index, GradientType: gradientType, Colors: colors, Position: position, Dimensions: dimensions}
}

// ── Document Generation ────────────────────────────────────────────────────

type DocumentMetadata struct {
  Title    string `json:"title"`
  Author   string `json:"author,omitempty"`
  Language string `json:"language,omitempty"`
}

type DocPageSize struct {
  Preset     string   `json:"preset,omitempty"`
  WidthInPt  *float64 `json:"width_in_pt,omitempty"`
  HeightInPt *float64 `json:"height_in_pt,omitempty"`
}

type DocMargins struct {
  TopInPt    float64 `json:"top_in_pt"`
  RightInPt  float64 `json:"right_in_pt"`
  BottomInPt float64 `json:"bottom_in_pt"`
  LeftInPt   float64 `json:"left_in_pt"`
}

type DocumentPage struct {
  Size            DocPageSize `json:"size"`
  Margins         DocMargins  `json:"margins"`
  BackgroundColor string      `json:"background_color,omitempty"`
}

type DocumentFontDefinition struct {
  Name   string `json:"name"`
  Weight string `json:"weight"`
  Style  string `json:"style"`
  Buffer string `json:"buffer"`
}

type TextStyle struct {
  FontFamily   string  `json:"font_family"`
  FontSizeInPt float64 `json:"font_size_in_pt"`
  LineHeight   float64 `json:"line_height"`
  Color        string  `json:"color"`
  IsBold       *bool   `json:"is_bold,omitempty"`
  IsItalic     *bool   `json:"is_italic,omitempty"`
}

type HeadlineStyle struct {
  FontFamily        string  `json:"font_family"`
  FontSizeInPt      float64 `json:"font_size_in_pt"`
  Color             string  `json:"color"`
  SpacingBeforeInPt float64 `json:"spacing_before_in_pt"`
  SpacingAfterInPt  float64 `json:"spacing_after_in_pt"`
  IsBold            *bool   `json:"is_bold,omitempty"`
  IsItalic          *bool   `json:"is_italic,omitempty"`
}

type LinkStyle struct {
  Color        string `json:"color"`
  IsUnderlined *bool  `json:"is_underlined,omitempty"`
}

type ListStyle struct {
  IndentInPt              float64 `json:"indent_in_pt"`
  SpacingBetweenItemsInPt float64 `json:"spacing_between_items_in_pt"`
}

type BorderStyle struct {
  WidthInPt float64 `json:"width_in_pt"`
  Color     string  `json:"color"`
}

type TableHeaderStyle struct {
  BackgroundColor string       `json:"background_color"`
  FontFamily      string       `json:"font_family"`
  FontSizeInPt    float64      `json:"font_size_in_pt"`
  Color           string       `json:"color"`
  PaddingInPt     float64      `json:"padding_in_pt"`
  IsBold          *bool        `json:"is_bold,omitempty"`
  Border          *BorderStyle `json:"border,omitempty"`
}

type TableBodyStyle struct {
  FontFamily               string       `json:"font_family"`
  FontSizeInPt             float64      `json:"font_size_in_pt"`
  Color                    string       `json:"color"`
  PaddingInPt              float64      `json:"padding_in_pt"`
  AlternateBackgroundColor string       `json:"alternate_background_color,omitempty"`
  Border                   *BorderStyle `json:"border,omitempty"`
}

type TableStyle struct {
  Header TableHeaderStyle `json:"header"`
  Body   TableBodyStyle   `json:"body"`
}

type GridStyle struct {
  GapInPt float64 `json:"gap_in_pt"`
}

type SeparatorStyle struct {
  Color            string  `json:"color"`
  ThicknessInPt    float64 `json:"thickness_in_pt"`
  MarginTopInPt    float64 `json:"margin_top_in_pt"`
  MarginBottomInPt float64 `json:"margin_bottom_in_pt"`
}

type ImageStyle struct {
  Alignment        string  `json:"alignment"`
  MarginTopInPt    float64 `json:"margin_top_in_pt"`
  MarginBottomInPt float64 `json:"margin_bottom_in_pt"`
}

type DocumentStyles struct {
  Text      TextStyle      `json:"text"`
  Headline  HeadlineStyle  `json:"headline"`
  Link      LinkStyle      `json:"link"`
  List      ListStyle      `json:"list"`
  Table     TableStyle     `json:"table"`
  Grid      GridStyle      `json:"grid"`
  Separator SeparatorStyle `json:"separator"`
  Image     ImageStyle     `json:"image"`
}

// ── Content Blocks ─────────────────────────────────────────────────────────

type ParagraphRun struct {
  Text     string `json:"text"`
  IsBold   *bool  `json:"is_bold,omitempty"`
  IsItalic *bool  `json:"is_italic,omitempty"`
  LinkURL  string `json:"link_url,omitempty"`
}

type ParagraphBlock struct {
  Type          string         `json:"type"`
  Markdown      string         `json:"markdown,omitempty"`
  TextAlignment string         `json:"text_alignment,omitempty"`
  Runs          []ParagraphRun `json:"runs,omitempty"`
}

func (ParagraphBlock) isContentBlock()      {}
func (ParagraphBlock) isHeaderFooterBlock() {}

func NewParagraphBlock() ParagraphBlock {
  return ParagraphBlock{Type: "paragraph"}
}

type HeadlineBlock struct {
  Type  string `json:"type"`
  Level string `json:"level"`
  Text  string `json:"text"`
}

func (HeadlineBlock) isContentBlock()      {}
func (HeadlineBlock) isHeaderFooterBlock() {}

func NewHeadlineBlock(level string, text string) HeadlineBlock {
  return HeadlineBlock{Type: "headline", Level: level, Text: text}
}

type ImageBlock struct {
  Type       string  `json:"type"`
  Buffer     string  `json:"buffer"`
  WidthInPt  float64 `json:"width_in_pt"`
  HeightInPt float64 `json:"height_in_pt"`
  Fit        string  `json:"fit,omitempty"`
}

func (ImageBlock) isContentBlock()      {}
func (ImageBlock) isHeaderFooterBlock() {}

func NewImageBlock(buffer string, widthInPt float64, heightInPt float64) ImageBlock {
  return ImageBlock{Type: "image", Buffer: buffer, WidthInPt: widthInPt, HeightInPt: heightInPt}
}

type TableCell struct {
  Text                string `json:"text"`
  HorizontalAlignment string `json:"horizontal_alignment,omitempty"`
  ColumnSpan          *int   `json:"column_span,omitempty"`
  RowSpan             *int   `json:"row_span,omitempty"`
}

type TableRow struct {
  Cells []TableCell `json:"cells"`
}

type TableBlock struct {
  Type                  string     `json:"type"`
  Rows                  []TableRow `json:"rows"`
  Header                *TableRow  `json:"header,omitempty"`
  ColumnWidthsInPercent []float64  `json:"column_widths_in_percent,omitempty"`
}

func (TableBlock) isContentBlock()      {}
func (TableBlock) isHeaderFooterBlock() {}

func NewTableBlock(rows []TableRow) TableBlock {
  return TableBlock{Type: "table", Rows: rows}
}

type GridColumn struct {
  ColumnSpan          int            `json:"column_span"`
  Blocks              []ContentBlock `json:"blocks"`
  HorizontalAlignment string         `json:"horizontal_alignment,omitempty"`
  VerticalAlignment   string         `json:"vertical_alignment,omitempty"`
}

type GridBlock struct {
  Type                string       `json:"type"`
  Columns             []GridColumn `json:"columns"`
  HorizontalAlignment string       `json:"horizontal_alignment,omitempty"`
  VerticalAlignment   string       `json:"vertical_alignment,omitempty"`
}

func (GridBlock) isContentBlock()      {}
func (GridBlock) isHeaderFooterBlock() {}

func NewGridBlock(columns []GridColumn) GridBlock {
  return GridBlock{Type: "grid", Columns: columns}
}

type ListItem struct {
  Text string `json:"text"`
}

type ListBlock struct {
  Type    string     `json:"type"`
  Variant string     `json:"variant"`
  Items   []ListItem `json:"items"`
}

func (ListBlock) isContentBlock()      {}
func (ListBlock) isHeaderFooterBlock() {}

func NewListBlock(variant string, items []ListItem) ListBlock {
  return ListBlock{Type: "list", Variant: variant, Items: items}
}

type TableOfContentsBlock struct {
  Type          string   `json:"type"`
  Levels        []string `json:"levels"`
  Leader        string   `json:"leader"`
  TextAlignment string   `json:"text_alignment,omitempty"`
}

func (TableOfContentsBlock) isContentBlock() {}

func NewTableOfContentsBlock(levels []string, leader string) TableOfContentsBlock {
  return TableOfContentsBlock{Type: "table-of-contents", Levels: levels, Leader: leader}
}

type PageBreakBlock struct {
  Type string `json:"type"`
}

func (PageBreakBlock) isContentBlock() {}

func NewPageBreakBlock() PageBreakBlock {
  return PageBreakBlock{Type: "page-break"}
}

type SeparatorBlock struct {
  Type string `json:"type"`
}

func (SeparatorBlock) isContentBlock()      {}
func (SeparatorBlock) isHeaderFooterBlock() {}

func NewSeparatorBlock() SeparatorBlock {
  return SeparatorBlock{Type: "separator"}
}

type DocumentQrCodeBlock struct {
  Type       string  `json:"type"`
  Value      string  `json:"value"`
  WidthInPt  float64 `json:"width_in_pt"`
  HeightInPt float64 `json:"height_in_pt"`
  FgHexColor string  `json:"fg_hex_color,omitempty"`
  BgHexColor string  `json:"bg_hex_color,omitempty"`
}

func (DocumentQrCodeBlock) isContentBlock() {}

func NewDocumentQrCodeBlock(value string, widthInPt float64, heightInPt float64) DocumentQrCodeBlock {
  return DocumentQrCodeBlock{Type: "qr-code", Value: value, WidthInPt: widthInPt, HeightInPt: heightInPt}
}

type DocumentBarcodeBlock struct {
  Type       string  `json:"type"`
  Value      string  `json:"value"`
  Format     string  `json:"format"`
  WidthInPt  float64 `json:"width_in_pt"`
  HeightInPt float64 `json:"height_in_pt"`
  FgHexColor string  `json:"fg_hex_color,omitempty"`
  BgHexColor string  `json:"bg_hex_color,omitempty"`
}

func (DocumentBarcodeBlock) isContentBlock() {}

func NewDocumentBarcodeBlock(value string, format string, widthInPt float64, heightInPt float64) DocumentBarcodeBlock {
  return DocumentBarcodeBlock{Type: "barcode", Value: value, Format: format, WidthInPt: widthInPt, HeightInPt: heightInPt}
}

type PageNumberBlock struct {
  Type          string `json:"type"`
  TextAlignment string `json:"text_alignment,omitempty"`
}

func (PageNumberBlock) isHeaderFooterBlock() {}

func NewPageNumberBlock() PageNumberBlock {
  return PageNumberBlock{Type: "page-number"}
}

type DocumentDefinition struct {
  Metadata                   DocumentMetadata         `json:"metadata"`
  Page                       DocumentPage             `json:"page"`
  Styles                     DocumentStyles           `json:"styles"`
  Content                    []ContentBlock           `json:"content"`
  Fonts                      []DocumentFontDefinition `json:"fonts,omitempty"`
  Header                     []HeaderFooterBlock      `json:"header,omitempty"`
  Footer                     []HeaderFooterBlock      `json:"footer,omitempty"`
  HeaderDistanceFromEdgeInPt *float64                 `json:"header_distance_from_edge_in_pt,omitempty"`
  FooterDistanceFromEdgeInPt *float64                 `json:"footer_distance_from_edge_in_pt,omitempty"`
}

// ── Request Types ──────────────────────────────────────────────────────────

type ExtractionSchema map[string]FieldConfig

type ExtractRequest struct {
  Files  []FileInput      `json:"files"`
  Schema ExtractionSchema `json:"schema"`
}

type ExtractAsyncRequest struct {
  Files      []FileInput      `json:"files"`
  Schema     ExtractionSchema `json:"schema"`
  WebhookURL string           `json:"webhook_url"`
}

type TransformRequest struct {
  File       FileInput            `json:"file"`
  Operations []TransformOperation `json:"operations"`
}

type TransformAsyncRequest struct {
  File       FileInput            `json:"file"`
  Operations []TransformOperation `json:"operations"`
  WebhookURL string               `json:"webhook_url"`
}

type GenerateImageRequest struct {
  Dimensions   Dimensions            `json:"dimensions"`
  Layers       []Layer               `json:"layers"`
  Fonts        []ImageFontDefinition `json:"fonts,omitempty"`
  OutputFormat string                `json:"output_format,omitempty"`
}

type GenerateImageAsyncRequest struct {
  Dimensions   Dimensions            `json:"dimensions"`
  Layers       []Layer               `json:"layers"`
  Fonts        []ImageFontDefinition `json:"fonts,omitempty"`
  OutputFormat string                `json:"output_format,omitempty"`
  WebhookURL   string                `json:"webhook_url"`
}

type GenerateDocumentRequest struct {
  Format   string             `json:"format"`
  Document DocumentDefinition `json:"document"`
}

type GenerateDocumentAsyncRequest struct {
  Format     string             `json:"format"`
  Document   DocumentDefinition `json:"document"`
  WebhookURL string             `json:"webhook_url"`
}
