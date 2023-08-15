package aci

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciPrivateLinkLabelfortheserviceEPg() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceAciPrivateLinkLabelfortheserviceEPgRead,
		SchemaVersion: 1,
		Schema: AppendBaseAttrSchema(AppendNameAliasAttrSchema(map[string]*schema.Schema{
			"cloud_service_epg_dn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		})),
	}
}

func dataSourceAciPrivateLinkLabelfortheserviceEPgRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)
	name := d.Get("name").(string)
	CloudServiceEPgDn := d.Get("cloud_service_epg_dn").(string)
	rn := fmt.Sprintf(models.RnCloudPrivateLinkLabel, name)
	dn := fmt.Sprintf("%s/%s", CloudServiceEPgDn, rn)

	cloudPrivateLinkLabel, err := getRemotePrivateLinkLabelfortheserviceEPg(aciClient, dn)
	if err != nil {
		return nil
	}

	d.SetId(dn)

	_, err = setPrivateLinkLabelfortheserviceEPgAttributes(cloudPrivateLinkLabel, d)
	if err != nil {
		return nil
	}

	return nil
}