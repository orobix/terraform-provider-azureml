package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceDatastore() *schema.Resource {
	return &schema.Resource{
		Description: "Use this resource to access the information of a specific Datastore of a certain Azure ML " +
			"Workspace. Authentication credentials are not included in the provided information.",

		ReadContext: dataSourceDatastoreRead,

		Schema: map[string]*schema.Schema{
			"resource_group_name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the resource group of the Azure ML Workspace to which the datastore belongs to.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"workspace_name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the Azure ML Workspace to which the datastore belongs to.",
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the datastore.",
				ValidateFunc: validation.StringIsNotEmpty,
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
	}
}

func dataSourceDatastoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient)
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	workspaceName := d.Get("workspace_name").(string)

	ds, err := client.ws.GetDatastore(resourceGroupName, workspaceName, name)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error retrieving datastore %s", name),
			Detail:   err.Error(),
		})
		return diags
	}

	if err := d.Set("description", ds.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_default", ds.IsDefault); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("storage_type", ds.StorageType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("storage_account_name", ds.StorageAccountName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("storage_container_name", ds.StorageContainerName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("credentials_type", ds.Auth.CredentialsType); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("creation_date", ds.SystemData.CreationDate.Format(defaultDateFormat)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("creation_user", ds.SystemData.CreationUser); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("creation_user_type", ds.SystemData.CreationUserType); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("last_modified_date", ds.SystemData.CreationDate.Format(defaultDateFormat)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_modified_user", ds.SystemData.LastModifiedUser); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_modified_user_type", ds.SystemData.LastModifiedUserType); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ds.Id)

	return diags
}
