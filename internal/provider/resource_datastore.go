package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/orobix/azureml-go-sdk/workspace"
	"time"
)

func resourceDatastore() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a Datastore.",

		CreateContext: resourceDatastoreCreate,
		ReadContext:   resourceDatastoreRead,
		UpdateContext: resourceDatastoreUpdate,
		DeleteContext: resourceDatastoreDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"resource_group_name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the resource group of the Azure ML Workspace to which the datastore belongs to.",
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"workspace_name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the Azure ML Workspace to which the datastore belongs to.",
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the datastore.",
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the datastore.",
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The description of the datastore.",
			},
			"is_default": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Is the datastore the default datastore of the Azure ML Workspace?",
			},
			"storage_type": {
				Type:     schema.TypeString,
				Required: true,
				Description: fmt.Sprintf(
					"The type of the storage to which the datstore is linked to. Possible values are: %+q",
					GetAllowedStorageTypes(),
				),
				ForceNew:     true,
				ValidateFunc: IsValidStorageType,
			},
			"storage_account_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The name of the Storage Account to which the datastore is linked to.",
				ForceNew:     true,
				ValidateFunc: IsValidStorageAccountName,
			},
			"storage_container_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The name of the Storage Container to which the datastore is linked to.",
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
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
			"auth": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"credentials_type": {
							Type:     schema.TypeString,
							Required: true,
							Description: fmt.Sprintf(
								"The type of credentials used for authenticating with the underlying storage. Possible values are: %+q.",
								GetAllowedCredentialTypes(),
							),
							ForceNew:     true,
							ValidateFunc: IsValidCredentialsType,
						},
						"tenant_id": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "The ID of the tenant to which the Service Principal used for authenticating " +
								"belongs to.",
						},
						"client_id": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "The application ID of the service principal used for authenticating with the " +
								"underlying storage of the datastore.",
						},
						"client_secret": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							Description: "The client secret of the service principal used for authenticating with the " +
								"underlying storage of the datastore.",
						},
						"account_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
							Description: "The primary key of the Storage Account linked to the datastore.",
						},
						"sql_user_name": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "The username of the identity used for authenticating with the SQL database linked " +
								"to the storage account.",
						},
						"sql_user_password": {
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
							Description: "The password of the identity used for authenticating with the SQL database linked " +
								"to the storage account.",
						},
					},
				},
			},
		},
	}
}

func resourceDatastoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)
	resourceGroupName := d.Get("resource_group_name").(string)
	workspaceName := d.Get("workspace_name").(string)

	datastore, err := resourceDatastoreGetResourceData(d)
	if err != nil {
		return diag.FromErr(err)
	}

	createdDatastore, err := client.ws.CreateOrUpdateDatastore(resourceGroupName, workspaceName, datastore)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(createdDatastore.Id)
	return resourceDatastoreSetResourceData(d, createdDatastore)
}

func resourceDatastoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient)
	resourceGroupName := d.Get("resource_group_name").(string)
	workspaceName := d.Get("workspace_name").(string)
	datastoreName := d.Get("name").(string)

	ds, err := client.ws.GetDatastore(resourceGroupName, workspaceName, datastoreName)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error reading datastore %s", datastoreName),
			Detail:   err.Error(),
		})
	}

	d.SetId(ds.Id)
	return resourceDatastoreSetResourceData(d, ds)
}

func resourceDatastoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)
	resourceGroupName := d.Get("resource_group_name").(string)
	workspaceName := d.Get("workspace_name").(string)

	datastore, err := resourceDatastoreGetResourceData(d)
	if err != nil {
		return diag.FromErr(err)
	}

	createdDatastore, err := client.ws.CreateOrUpdateDatastore(resourceGroupName, workspaceName, datastore)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(createdDatastore.Id)
	return resourceDatastoreSetResourceData(d, createdDatastore)
}

func resourceDatastoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*apiClient)
	resourceGroupName := d.Get("resource_group_name").(string)
	workspaceName := d.Get("workspace_name").(string)
	datastoreName := d.Get("name").(string)

	err := client.ws.DeleteDatastore(resourceGroupName, workspaceName, datastoreName)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Error deleting datastore %s.", datastoreName),
			Detail:   err.Error(),
		})
	}
	return diags
}

func resourceDatastoreGetResourceData(d *schema.ResourceData) (*workspace.Datastore, error) {
	var creationDate time.Time
	var lastModifiedDate time.Time
	var err error

	if d.Get("creation_date") != "" {
		creationDate, err = time.Parse(defaultDateFormat, d.Get("creation_date").(string))
		if err != nil {
			return nil, err
		}
	}
	if d.Get("last_modified_date") != "" {
		lastModifiedDate, err = time.Parse(defaultDateFormat, d.Get("last_modified_date").(string))
		if err != nil {
			return nil, err
		}
	}

	auth, err := schemaSetToDatastoreAuth(d.Get("auth").(*schema.Set))
	if err != nil {
		return nil, err
	}

	return &workspace.Datastore{
		Id:                   d.Get("id").(string),
		Name:                 d.Get("name").(string),
		IsDefault:            d.Get("is_default").(bool),
		Description:          d.Get("description").(string),
		StorageType:          d.Get("storage_type").(string),
		StorageAccountName:   d.Get("storage_account_name").(string),
		StorageContainerName: d.Get("storage_container_name").(string),
		SystemData: &workspace.SystemData{
			CreationDate:         creationDate,
			CreationUser:         d.Get("creation_user").(string),
			CreationUserType:     d.Get("creation_user_type").(string),
			LastModifiedDate:     lastModifiedDate,
			LastModifiedUser:     d.Get("last_modified_user").(string),
			LastModifiedUserType: d.Get("last_modified_user_type").(string),
		},
		Auth: auth,
	}, nil
}

func resourceDatastoreSetResourceData(d *schema.ResourceData, datastore *workspace.Datastore) diag.Diagnostics {
	if err := d.Set("description", datastore.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_default", datastore.IsDefault); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("storage_type", datastore.StorageType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("storage_account_name", datastore.StorageAccountName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("storage_container_name", datastore.StorageContainerName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("creation_date", datastore.SystemData.CreationDate.Format(defaultDateFormat)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("creation_user", datastore.SystemData.CreationUser); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("creation_user_type", datastore.SystemData.CreationUserType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_modified_date", datastore.SystemData.CreationDate.Format(defaultDateFormat)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_modified_user", datastore.SystemData.LastModifiedUser); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_modified_user_type", datastore.SystemData.LastModifiedUserType); err != nil {
		return diag.FromErr(err)
	}

	currentAuthSet := d.Get("auth").(*schema.Set)
	currentAuth := currentAuthSet.List()[0].(map[string]interface{})
	currentAuth["credentials_type"] = datastore.Auth.CredentialsType
	if err := d.Set("auth", currentAuthSet); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func schemaSetToDatastoreAuth(set *schema.Set) (*workspace.DatastoreAuth, error) {
	data := set.List()[0].(map[string]interface{})
	auth := new(workspace.DatastoreAuth)
	auth.CredentialsType = data["credentials_type"].(string)
	if data["client_id"] != nil {
		auth.ClientId = data["client_id"].(string)
	}
	if data["tenant_id"] != nil {
		auth.TenantId = data["tenant_id"].(string)
	}
	if data["client_secret"] != nil {
		auth.ClientSecret = data["client_secret"].(string)
	}
	if data["account_key"] != nil {
		auth.AccountKey = data["account_key"].(string)
	}
	if data["sql_user_password"] != nil {
		auth.SqlUserPassword = data["sql_user_password"].(string)
	}
	if data["sql_user_name"] != nil {
		auth.SqlUserName = data["sql_user_name"].(string)
	}

	return auth, nil
}
