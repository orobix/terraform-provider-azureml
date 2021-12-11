package provider

import (
	"context"
	"fmt"
	"github.com/Telemaco019/azureml-go-sdk/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceDatastores() *schema.Resource {
	return &schema.Resource{
		Description: "Use this resource to retrieve the list of Datastores of a certain Azure ML " +
			"Workspace. Authentication credentials are not included in the provided information.",

		ReadContext: dataSourceDatastoresRead,

		Schema: map[string]*schema.Schema{
			"resource_group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the resource group of the Azure ML Workspace to which the datastore belongs to.",
			},
			"workspace_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Azure ML Workspace to which the datastore belongs to.",
			},
			"datastores": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the resource group of the Azure ML Workspace to which the datastore belongs to.",
						},
						"workspace_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Azure ML Workspace to which the datastore belongs to.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the datastore.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the datastore.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the datastore.",
						},
						"is_default": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is the datastore the default datastore of the Azure ML Workspace?",
						},
						"storage_type": {
							Type:     schema.TypeString,
							Computed: true,
							Description: fmt.Sprintf(
								"The type of the storage to which the datstore is linked to. Possible values are: %v",
								//NewStorageTypeValidator().allowedTypes,
							),
						},
						"storage_account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Storage Account to which the datastore is linked to.",
						},
						"storage_container_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Storage Container to which the datastore is linked to.",
						},
						"credentials_type": {
							Type:     schema.TypeString,
							Computed: true,
							Description: fmt.Sprintf(
								"The type of credentials used for authenticating with the underlying storage. ",
								//NewDatastoreCredentialsTypeValidator().allowedTypes,
							),
						},
						"creation_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp corresponding to the creation of the datastore.",
						},
						"creation_user": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user that created the datastore.",
						},
						"creation_user_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The kind of user that created the datastore (Service Principal or User).",
						},
						"last_modified_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The timestamp corresponding to the last update of the datastore.",
						},
						"last_modified_user": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user that last updated the datastore.",
						},
						"last_modified_user_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The kind of user that last updated the datastore (Service Principal or User).",
						},
					},
				},
			},
		},
	}
}

func dataSourceDatastoresRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient)
	resourceGroupName := d.Get("resource_group_name").(string)
	workspaceName := d.Get("workspace_name").(string)

	dsl, err := client.ws.GetDatastores(resourceGroupName, workspaceName)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error retrieving datastores.",
			Detail:   err.Error(),
		})
		return diags
	}

	values, err := listFromDatastores(dsl)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error retrieving datastores.",
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("datastores", values); err != nil {
		return diag.FromErr(err)
	}

	id, err := hash(fmt.Sprintf("%s%s%d", resourceGroupName, workspaceName, len(values)))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.Itoa(int(id)))

	return diags
}

func listFromDatastores(dsl []workspace.Datastore) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, len(dsl))
	for i, ds := range dsl {
		v, err := fromDatastore(ds)
		if err != nil {
			return nil, fmt.Errorf("unable to parse datastore: %w", err)
		}
		result[i] = v
	}
	return result, nil
}

func fromDatastore(ds workspace.Datastore) (map[string]interface{}, error) {
	return map[string]interface{}{
		"name":                    ds.Name,
		"description":             ds.Description,
		"is_default":              ds.IsDefault,
		"storage_type":            ds.StorageType,
		"storage_account_name":    ds.StorageAccountName,
		"storage_container_name":  ds.StorageContainerName,
		"credentials_type":        ds.Auth.CredentialsType,
		"creation_date":           ds.SystemData.CreationDate.Format(defaultDateFormat),
		"creation_user":           ds.SystemData.CreationUser,
		"creation_user_type":      ds.SystemData.CreationUserType,
		"last_modified_date":      ds.SystemData.LastModifiedDate.Format(defaultDateFormat),
		"last_modified_user":      ds.SystemData.LastModifiedUser,
		"last_modified_user_type": ds.SystemData.LastModifiedUserType,
	}, nil
}
