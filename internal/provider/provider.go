package provider

import (
	"context"
	"github.com/Telemaco019/azureml-go-sdk/workspace"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

const (
	defaultDateFormat = time.RFC1123
)

func New() tfsdk.Provider {
	return &provider{}
}

type provider struct {
	configured bool
	client     *workspace.Workspace
}

// Provider schema struct
type providerData struct {
	ClientId       types.String `tfsdk:"client_id"`
	ClientSecret   types.String `tfsdk:"client_secret"`
	TenantId       types.String `tfsdk:"tenant_id"`
	SubscriptionId types.String `tfsdk:"subscription_id"`
}

func (p provider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"client_id": {
				Type:     types.StringType,
				Required: true,
			},
			"client_secret": {
				Type:      types.StringType,
				Required:  true,
				Sensitive: true,
			},
			"tenant_id": {
				Type:     types.StringType,
				Required: true,
			},
			"subscription_id": {
				Type:     types.StringType,
				Required: true,
			},
		},
	}, nil
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// Retrieve provider data from configuration
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	//// User must provide a user to the provider
	//var username string
	//if config.Username.Unknown {
	//	// Cannot connect to client with an unknown value
	//	resp.Diagnostics.AddWarning(
	//		"Unable to create client",
	//		"Cannot use unknown value as username",
	//	)
	//	return
	//}
	//
	//if config.Username.Null {
	//	username = os.Getenv("HASHICUPS_USERNAME")
	//} else {
	//	username = config.Username.Value
	//}
	//
	//if username == "" {
	//	// Error vs warning - empty value must stop execution
	//	resp.Diagnostics.AddError(
	//		"Unable to find username",
	//		"Username cannot be an empty string",
	//	)
	//	return
	//}
	//
	//// User must provide a password to the provider
	//var password string
	//if config.Password.Unknown {
	//	// Cannot connect to client with an unknown value
	//	resp.Diagnostics.AddError(
	//		"Unable to create client",
	//		"Cannot use unknown value as password",
	//	)
	//	return
	//}
	//
	//if config.Password.Null {
	//	password = os.Getenv("HASHICUPS_PASSWORD")
	//} else {
	//	password = config.Password.Value
	//}
	//
	//if password == "" {
	//	// Error vs warning - empty value must stop execution
	//	resp.Diagnostics.AddError(
	//		"Unable to find password",
	//		"password cannot be an empty string",
	//	)
	//	return
	//}
	//
	//// User must specify a host
	//var host string
	//if config.Host.Unknown {
	//	// Cannot connect to client with an unknown value
	//	resp.Diagnostics.AddError(
	//		"Unable to create client",
	//		"Cannot use unknown value as host",
	//	)
	//	return
	//}
	//
	//if config.Host.Null {
	//	host = os.Getenv("HASHICUPS_HOST")
	//} else {
	//	host = config.Host.Value
	//}
	//
	//if host == "" {
	//	// Error vs warning - empty value must stop execution
	//	resp.Diagnostics.AddError(
	//		"Unable to find host",
	//		"Host cannot be an empty string",
	//	)
	//	return
	//}

	// Create a new HashiCups client and set it to the provider client
	ws, err := workspace.New(workspace.Config{
		ClientId:       config.ClientId.Value,
		ClientSecret:   config.ClientSecret.Value,
		TenantId:       config.TenantId.Value,
		SubscriptionId: config.SubscriptionId.Value,
	}, false)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create hashicups client:\n\n"+err.Error(),
		)
		return
	}

	p.client = ws
	p.configured = true
}

func (p *provider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{}, nil
}

// GetDataSources - Defines provider data sources
func (p *provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"azureml_datastores": dataSourceDatastoresType{},
	}, nil
}
