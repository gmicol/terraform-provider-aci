// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvRsFcPathAttWithFvAEPg(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "2.0(1m)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsFcPathAttMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test_2", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test", "vsan", "unknown"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test_2", "vsan", "unknown"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test", "vsan_mode", "regular"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test_2", "vsan_mode", "regular"),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "2.0(1m)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigFvRsFcPathAttMinDependencyWithFvAEPgAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "2.0(1m)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsFcPathAttMinDependencyWithFvAEPgAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test_2", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test", "vsan", "unknown"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test_2", "vsan", "unknown"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test", "vsan_mode", "regular"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.allow_test_2", "vsan_mode", "regular"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "apic", "2.0(1m)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsFcPathAttMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "vsan", "unknown"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "vsan_mode", "regular"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRsFcPathAttAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "vsan", "vsan-10"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "vsan_mode", "native"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRsFcPathAttMinDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRsFcPathAttResetDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "vsan", "unknown"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "vsan_mode", "regular"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_to_fibre_channel_path.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvRsFcPathAttChildrenDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "target_dn", "topology/pod-1/paths-101/pathep-[eth1/1]"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "description", ""),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "vsan", "unknown"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "vsan_mode", "regular"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_to_fibre_channel_path.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvRsFcPathAttChildrenRemoveFromConfigDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvRsFcPathAttChildrenRemoveOneDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvRsFcPathAttChildrenRemoveAllDependencyWithFvAEPg,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_to_fibre_channel_path.test", "tags.#", "0"),
				),
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}

const testDependencyConfigFvRsFcPathAtt = `
`

const testConfigFvRsFcPathAttMinDependencyWithFvAEPgAllowExisting = testDependencyConfigFvRsFcPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_fibre_channel_path" "allow_test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
}
resource "aci_relation_to_fibre_channel_path" "allow_test_2" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
  depends_on = [aci_relation_to_fibre_channel_path.allow_test]
}
`

const testConfigFvRsFcPathAttMinDependencyWithFvAEPg = testDependencyConfigFvRsFcPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_fibre_channel_path" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
}
`

const testConfigFvRsFcPathAttAllDependencyWithFvAEPg = testDependencyConfigFvRsFcPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_fibre_channel_path" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
  annotation = "annotation"
  description = "description_1"
  vsan = "vsan-10"
  vsan_mode = "native"
}
`

const testConfigFvRsFcPathAttResetDependencyWithFvAEPg = testDependencyConfigFvRsFcPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_fibre_channel_path" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
  annotation = "orchestrator:terraform"
  description = ""
  vsan = "unknown"
  vsan_mode = "regular"
}
`
const testConfigFvRsFcPathAttChildrenDependencyWithFvAEPg = testDependencyConfigFvRsFcPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_fibre_channel_path" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
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

const testConfigFvRsFcPathAttChildrenRemoveFromConfigDependencyWithFvAEPg = testDependencyConfigFvRsFcPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_fibre_channel_path" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
}
`

const testConfigFvRsFcPathAttChildrenRemoveOneDependencyWithFvAEPg = testDependencyConfigFvRsFcPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_fibre_channel_path" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
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

const testConfigFvRsFcPathAttChildrenRemoveAllDependencyWithFvAEPg = testDependencyConfigFvRsFcPathAtt + testConfigFvAEPgMinDependencyWithFvAp + `
resource "aci_relation_to_fibre_channel_path" "test" {
  parent_dn = aci_application_epg.test.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
  annotations = []
  tags = []
}
`
