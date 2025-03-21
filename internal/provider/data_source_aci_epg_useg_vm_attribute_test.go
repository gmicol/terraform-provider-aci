// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceFvVmAttrWithFvCrtrn(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "1.1(1j)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvVmAttrDataSourceDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "name", "vm_attribute"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "category", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "label_name", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "operator", "equals"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "type", "vm-name"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "value", "default_value"),
				),
			},
			{
				Config:      testConfigFvVmAttrNotExistingFvCrtrn,
				ExpectError: regexp.MustCompile("Failed to read aci_epg_useg_vm_attribute data source"),
			},
		},
	})
}
func TestAccDataSourceFvVmAttrWithFvSCrtrn(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "1.1(1j)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             testConfigFvVmAttrDataSourceDependencyWithFvSCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "name", "vm_attribute"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "category", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "label_name", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "operator", "equals"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "type", "vm-name"),
					resource.TestCheckResourceAttr("data.aci_epg_useg_vm_attribute.test", "value", "default_value"),
				),
			},
			{
				Config:      testConfigFvVmAttrNotExistingFvSCrtrn,
				ExpectError: regexp.MustCompile("Failed to read aci_epg_useg_vm_attribute data source"),
			},
		},
	})
}

const testConfigFvVmAttrDataSourceDependencyWithFvCrtrn = testConfigFvVmAttrMinDependencyWithFvCrtrn + `
data "aci_epg_useg_vm_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  name = "vm_attribute"
  depends_on = [aci_epg_useg_vm_attribute.test]
}
`

const testConfigFvVmAttrNotExistingFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
data "aci_epg_useg_vm_attribute" "test_non_existing" {
  parent_dn = aci_epg_useg_block_statement.test.id
  name = "vm_attribute_non_existing"
}
`
const testConfigFvVmAttrDataSourceDependencyWithFvSCrtrn = testConfigFvVmAttrMinDependencyWithFvSCrtrn + `
data "aci_epg_useg_vm_attribute" "test" {
  parent_dn = aci_epg_useg_sub_block_statement.test.id
  name = "vm_attribute"
  depends_on = [aci_epg_useg_vm_attribute.test]
}
`

const testConfigFvVmAttrNotExistingFvSCrtrn = testConfigFvSCrtrnMinDependencyWithFvCrtrn + `
data "aci_epg_useg_vm_attribute" "test_non_existing" {
  parent_dn = aci_epg_useg_sub_block_statement.test.id
  name = "vm_attribute_non_existing"
}
`
