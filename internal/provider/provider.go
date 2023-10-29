package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

const reolinkVersion = "v0.6.3"

// reolink is the provider implementation.
type reolinkProvider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &reolinkProvider{
			version: version,
		}
	}
}

// Metadata returns the provider type name.
func (p *reolinkProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "reolink"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *reolinkProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *reolinkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *reolinkProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewManifestDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *reolinkProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
