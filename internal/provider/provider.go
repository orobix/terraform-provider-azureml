package provider

import (
	"context"
	"github.com/Telemaco019/azureml-go-sdk/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

const (
	defaultDateFormat = time.RFC1123
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"client_id": {
					Type:     schema.TypeString,
					Required: true,
				},
				"client_secret": {
					Type:      schema.TypeString,
					Required:  true,
					Sensitive: true,
				},
				"tenant_id": {
					Type:     schema.TypeString,
					Required: true,
				},
				"subscription_id": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"azureml_datastore":  dataSourceDatastore(),
				"azureml_datastores": dataSourceDatastores(),
			},
			ResourcesMap: map[string]*schema.Resource{},
		}
		p.ConfigureContextFunc = configure(version, p)
		return p
	}
}

type apiClient struct {
	ws *workspace.Workspace
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(_ context.Context, r *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var apiClient = new(apiClient)
		var diags diag.Diagnostics

		ws, err := workspace.New(workspace.Config{
			ClientId:       r.Get("client_id").(string),
			ClientSecret:   r.Get("client_secret").(string),
			TenantId:       r.Get("tenant_id").(string),
			SubscriptionId: r.Get("subscription_id").(string),
		}, false)

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create client",
				Detail:   "Unable to create azureml client:\n\n" + err.Error(),
			})
			return nil, diags
		}

		apiClient.ws = ws
		return apiClient, diags
	}
}
