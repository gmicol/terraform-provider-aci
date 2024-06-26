package aci

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciCloudTemplateRegion() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceAciCloudTemplateRegionRead,
		SchemaVersion: 1,
		Schema: AppendBaseAttrSchema(AppendNameAliasAttrSchema(map[string]*schema.Schema{
			"parent_dn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hub_networking": {
				Type:     schema.TypeString,
				Computed: true,
			},
		})),
	}
}

func dataSourceAciCloudTemplateRegionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)
	cloudProviderandRegionNamesDn := d.Get("parent_dn").(string)
	dn := fmt.Sprintf("%s/%s", cloudProviderandRegionNamesDn, models.RnCloudtemplateRegionDetail)

	cloudTemplateRegionDetail, err := getRemoteCloudTemplateRegion(aciClient, dn)
	if err != nil {
		return nil
	}

	d.SetId(dn)

	_, err = setCloudTemplateRegionAttributes(cloudTemplateRegionDetail, d)
	if err != nil {
		return nil
	}

	return nil
}
