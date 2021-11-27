package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct {
	configured bool
	client     interface{} // TODO
}

func (p provider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	panic("implement me")
}

func (p provider) Configure(ctx context.Context, request tfsdk.ConfigureProviderRequest, response *tfsdk.ConfigureProviderResponse) {
	panic("implement me")
}

func (p provider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	panic("implement me")
}

func (p provider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	panic("implement me")
}
