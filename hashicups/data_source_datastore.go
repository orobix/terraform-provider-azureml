package hashicups

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDatastores() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDatastoresRead,
		Schema: map[string]*schema.Schema{
			"coffees": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"teaser": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"price": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"image": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ingredients": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ingredient_id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceDatastoresRead(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics {

}
