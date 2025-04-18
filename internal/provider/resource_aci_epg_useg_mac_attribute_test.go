// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvMacAttrWithFvCrtrn(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "1.1(1j)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvMacAttrMinDependencyWithFvCrtrnAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "name", "mac_attr"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "name", "mac_attr"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "mac", "AA:BB:CC:DD:EE:FF"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "mac", "AA:BB:CC:DD:EE:FF"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "owner_tag", ""),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "1.1(1j)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigFvMacAttrMinDependencyWithFvCrtrnAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "1.1(1j)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvMacAttrMinDependencyWithFvCrtrnAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "name", "mac_attr"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "name", "mac_attr"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "mac", "AA:BB:CC:DD:EE:FF"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "mac", "AA:BB:CC:DD:EE:FF"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.allow_test_2", "owner_tag", ""),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "1.1(1j)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvMacAttrMinDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "name", "mac_attr"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "mac", "AA:BB:CC:DD:EE:FF"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "owner_tag", ""),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvMacAttrAllDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "name", "mac_attr"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "mac", "AA:BB:CC:BB:BB:EE"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "owner_tag", "owner_tag_1"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvMacAttrMinDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "mac", "AA:BB:CC:DD:EE:FF"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "name", "mac_attr"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvMacAttrResetDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "mac", "AA:BB:CC:DD:EE:FF"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "name", "mac_attr"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "owner_tag", ""),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_epg_useg_mac_attribute.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvMacAttrChildrenDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "mac", "AA:BB:CC:DD:EE:FF"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "name", "mac_attr"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "description", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_epg_useg_mac_attribute.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvMacAttrChildrenRemoveFromConfigDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvMacAttrChildrenRemoveOneDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvMacAttrChildrenRemoveAllDependencyWithFvCrtrn,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_epg_useg_mac_attribute.test", "tags.#", "0"),
				),
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}

const testConfigFvMacAttrMinDependencyWithFvCrtrnAllowExisting = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_mac_attribute" "allow_test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  mac = "AA:BB:CC:DD:EE:FF"
  name = "mac_attr"
}
resource "aci_epg_useg_mac_attribute" "allow_test_2" {
  parent_dn = aci_epg_useg_block_statement.test.id
  mac = "AA:BB:CC:DD:EE:FF"
  name = "mac_attr"
  depends_on = [aci_epg_useg_mac_attribute.allow_test]
}
`

const testConfigFvMacAttrMinDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_mac_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  mac = "AA:BB:CC:DD:EE:FF"
  name = "mac_attr"
}
`

const testConfigFvMacAttrAllDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_mac_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  name = "mac_attr"
  annotation = "annotation"
  description = "description_1"
  mac = "AA:BB:CC:BB:BB:EE"
  name_alias = "name_alias_1"
  owner_key = "owner_key_1"
  owner_tag = "owner_tag_1"
}
`

const testConfigFvMacAttrResetDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_mac_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  name = "mac_attr"
  annotation = "orchestrator:terraform"
  description = ""
  mac = "AA:BB:CC:DD:EE:FF"
  name_alias = ""
  owner_key = ""
  owner_tag = ""
}
`
const testConfigFvMacAttrChildrenDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_mac_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  mac = "AA:BB:CC:DD:EE:FF"
  name = "mac_attr"
  annotations = [
    {
      key = "key_0"
      value = "value_1"
    },
    {
      key = "key_1"
      value = "test_value"
    },
  ]
  tags = [
    {
      key = "key_0"
      value = "value_1"
    },
    {
      key = "key_1"
      value = "test_value"
    },
  ]
}
`

const testConfigFvMacAttrChildrenRemoveFromConfigDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_mac_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  mac = "AA:BB:CC:DD:EE:FF"
  name = "mac_attr"
}
`

const testConfigFvMacAttrChildrenRemoveOneDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_mac_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  mac = "AA:BB:CC:DD:EE:FF"
  name = "mac_attr"
  annotations = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  tags = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
}
`

const testConfigFvMacAttrChildrenRemoveAllDependencyWithFvCrtrn = testConfigFvCrtrnMinDependencyWithFvAEPg + `
resource "aci_epg_useg_mac_attribute" "test" {
  parent_dn = aci_epg_useg_block_statement.test.id
  mac = "AA:BB:CC:DD:EE:FF"
  name = "mac_attr"
  annotations = []
  tags = []
}
`
