package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var _ provider.Provider = &punycodeProvider{}
var _ provider.ProviderWithFunctions = &punycodeProvider{}

// punycodeProvider is the provider implementation.
type punycodeProvider struct{}

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &punycodeProvider{}
}

// Metadata returns the provider type name.
func (p *punycodeProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "punycode"
}

// Schema defines the provider-level schema for configuration data.
func (p *punycodeProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provider for punycode encoding and decoding operations.",
	}
}

// Configure prepares a provider instance for data sources and resources.
func (p *punycodeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// No configuration needed for this provider
}

// DataSources defines the data sources implemented in the provider.
func (p *punycodeProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *punycodeProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

// Functions defines the functions implemented in the provider.
func (p *punycodeProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		NewEncodeFunction,
		NewDecodeFunction,
	}
}
