package otelcolmetadataschema

import "github.com/invopop/jsonschema"

// cf.)
// - https://github.com/open-telemetry/opentelemetry-collector/blob/main/cmd/mdatagen/metadata-schema.yaml
// - https://github.com/open-telemetry/opentelemetry-collector/blob/main/cmd/mdatagen/loader.go

type Metadata struct {
	Type               string                              `json:"type" jsonschema_description:"The type of the component - Usually the name. The type and class combined uniquely identify the component (eg. receiver/otlp) or subcomponent (eg. receiver/hostmetricsreceiver/cpu)"`
	Parent             string                              `json:"parent,omitempty" jsonschema_description:"Required for subcomponents: The type of the parent component."`
	Status             *Status                             `json:"status,omitempty" jsonschema_description:"Required for components (Optional for subcomponents): A high-level view of the development status and use of this component"`
	SemConvVersion     string                              `json:"sem_conv_version,omitempty" jsonschema_description:"OTel Semantic Conventions version that will be associated with the scraped metrics. This attribute should be set for metrics compliant with OTel Semantic Conventions."`
	ResourceAttributes map[AttributeName]ResourceAttribute `json:"resource_attributes,omitempty" jsonschema_description:"map of resource attribute definitions with the key being the attribute name."`
	Attributes         map[AttributeName]Attribute         `json:"attributes,omitempty" jsonschema_description:"map of attribute definitions with the key being the attribute name and value being described below."`
	Metrics            map[MetricName]Metric               `json:"metrics" jsonschema_description:"map of metric names with the key being the metric name and valuem being described below."`
	Telemetry          map[MetricName]Metric               `json:"telemetry,omitempty" jsonschema_description:"map of metric names with the key being the metric name and valuem being described below."`
	ScopeName          string                              `json:"scope_name,omitempty"`
	// ShortFolderName    string               `json:"-"`
	Tests Tests `json:"tests,omitempty" jsonschema_description:"Lifecycle tests generated for this component."`
}

type Status struct {
	Stability            Stability   `json:"stability" jsonschema_description:"The stability of the component - See https://github.com/open-telemetry/opentelemetry-collector#stability-levels"`
	Distributions        []string    `json:"distributions,omitempty" jsonschema_description:"The distributions that this component is bundled with (For example core or contrib). See statusdata.go for a list of common distros."`
	Class                string      `json:"class" jsonschema:"enum=receiver,enum=processor,enum=exporter,enum=connector,enum=extension,enum=cmd,enum=pkg" jsonschema_description:"The class of the component (For example receiver)"`
	Warnings             []string    `json:"warnings,omitempty" jsonschema_description:"A list of warnings that should be brought to the attention of users looking to use this component"`
	Codeowners           *Codeowners `json:"codeowners,omitempty" jsonschema_description:"Metadata related to codeowners of the component"`
	UnsupportedPlatforms []string    `json:"unsupported_platforms,omitempty"`
}

type Stability struct {
	Development  []StabilityItem `json:"development,omitempty"`
	Alpha        []StabilityItem `json:"alpha,omitempty"`
	Beta         []StabilityItem `json:"beta,omitempty"`
	Stable       []StabilityItem `json:"stable,omitempty"`
	Deprecated   []StabilityItem `json:"deprecated,omitempty"`
	Unmaintained []StabilityItem `json:"unmaintained,omitempty"`
}

type StabilityItem string

func (StabilityItem) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type: "string",
		Enum: []any{
			"metrics", "traces", "logs", "traces_to_metrics",
			"metrics_to_metrics", "logs_to_metrics", "extension",
		},
	}
}

type Codeowners struct {
	Active     []string `json:"active,omitempty"`
	Emeritus   []string `json:"emeritus,omitempty"`
	SeekingNew bool     `json:"seeking_new,omitempty"`
}

type AttributeName string

type ResourceAttribute struct {
	Description string         `json:"description" jsonschema_description:"description of the attribute."`
	Enabled     bool           `json:"enabled" jsonschema_description:"whether the resource attribute is added the emitted metrics by default."`
	Include     []FilterConfig `json:"include,omitempty"`
	Exclude     []FilterConfig `json:"exclude,omitempty"`
	Enum        []string       `json:"enum,omitempty" jsonschema_description:"array of attribute values if they are static values (currently, only string type is supported)."`
	Type        string         `json:"type" jsonschema:"enum=string,enum=int,enum=double,enum=bool,enum=bytes,enum=slice,enum=map" jsonschema_description:"attribute value type."`
	// FullName     attributeName   `json:"-"`
	Warnings Warnings `json:"warnings,omitempty" jsonschema_description:"warnings that will be shown to user under specified conditions."`
}

type Attribute struct {
	Description  string         `json:"description" jsonschema_description:"description of the attribute."`
	NameOverride string         `json:"name_override,omitempty" jsonschema_description:"this field can be used to override the actual attribute name defined by the key. It should be used if multiple metrics have different attributes with the same name."`
	Include      []FilterConfig `json:"include,omitempty"`
	Exclude      []FilterConfig `json:"exclude,omitempty"`
	Enum         []string       `json:"enum,omitempty" jsonschema_description:" array of attribute values if they are static values (currently, only string type is supported)."`
	Type         string         `json:"type" jsonschema:"enum=string,enum=int,enum=double,enum=bool,enum=bytes,enum=slice,enum=map" jsonschema_description:"attribute value type."`
	// FullName     attributeName   `json:"-"`
}

type FilterConfig struct {
	Strict string `json:"strict,omitempty" jsonschema:"oneof_required=filter"`
	Regex  string `json:"regexp,omitempty" jsonschema:"oneof_required=filter"`
}

type MetricName string

type Metric struct {
	Enabled               bool            `json:"enabled" jsonschema_description:"whether the metric is collected by default."`
	Warnings              Warnings        `json:"warnings,omitempty" jsonschema_description:"warnings that will be shown to user under specified conditions."`
	Description           string          `json:"description" jsonschema_description:"metric description."`
	ExtendedDocumentation string          `json:"extended_documentation,omitempty" jsonschema_description:"extended documentation of the metric."`
	Unit                  *string         `json:"unit" jsonschema:"oneof_type=string;number" jsonschema_description:"metric unit as defined by https://ucum.org/ucum.html."`
	Sum                   *Sum            `json:"sum,omitempty" jsonschema:"oneof_required:metrictype" jsonschema_description:"metric type with its settings."`
	Gauge                 *Gauge          `json:"gauge,omitempty" jsonschema:"oneof_required:metrictype" jsonschema_description:"metric type with its settings."`
	Attributes            []AttributeName `json:"attributes,omitempty" jsonschema_description:"array of attributes that were defined in the attributes section that are emitted by this metric."`
}

type Sum struct {
	AggregationTemporality string `json:"aggregation_temporality" jsonschema_description:"whether reported values incorporate previous measurements (cumulative) or not (delta)."`
	Mono                   bool   `json:"monotonic" jsonschema_description:"whether the metric is monotonic (no negative delta values)."`
	MetricValueType        string `json:"value_type" jsonschema_description:"type of number data point values."`
	MetricInputType        string `json:"input_type,omitempty" jsonschema_description:"Indicates the type the metric needs to be parsed from. If set, the generated functions will parse the value from string to value_type."`
}

type Gauge struct {
	MetricValueType string `json:"value_type" jsonschema_description:"type of number data point values."`
	MetricInputType string `json:"input_type,omitempty" jsonschema_description:"Indicates the type the metric needs to be parsed from. If set, the generated functions will parse the value from string to value_type."`
}

type Warnings struct {
	// TODO: each Warnings should have different tags for descriptions

	IfEnabled       string `json:"if_enabled,omitempty"`
	IfEnabledNotSet string `json:"if_enabled_not_set,omitempty"`
	IfConfigured    string `json:"if_configured,omitempty"`
}

type Tests struct {
	Config              any    `json:"config,omitempty" jsonschema_description:"{} by default, specific testing configuration for lifecycle tests."`
	SkipLifecycle       bool   `json:"skip_lifecycle,omitempty" jsonschema_description:"false by default Skip lifecycle tests for this component. Not recommended for components that are not in development."`
	SkipShutdown        bool   `json:"skip_shutdown,omitempty" jsonschema_description:"false by default Skip shutdown tests for this component. Not recommended for components that are not in development."`
	GoLeak              GoLeak `json:"goleak,omitempty" jsonschema_description:"{} by default generates a package_test to enable check for leaks"`
	ExpectConsumerError bool   `json:"expect_consumer_error,omitempty" jsonschema_description:"Whether it's expected that the Consume[Logs|Metrics|Traces] method will return an error with the given configuration."`
}

type GoLeak struct {
	Skip     bool   `json:"skip" jsonschema_description:"set to true if goleak tests should be skipped"`
	Ignore   Ignore `json:"ignore,omitempty"`
	Setup    string `json:"setup,omitempty" jsonschema_description:"upports configuring a setup function that runs before goleak checks"`
	Teardown string `json:"teardown,omitempty" jsonschema_description:"supports configuring a teardown function that runs before goleak checks"`
}

type Ignore struct {
	Top []string `json:"top,omitempty" jsonschema_description:"array of strings representing functions that should be ignore via IgnoreTopFunction"`
	Any []string `json:"any,omitempty" jsonschema_description:"array of strings representing functions that should be ignore via IgnoreAnyFunction"`
}
