// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceQosDppPolWithFvTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.2(2g)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigQosDppPolMinDependencyWithFvTenantAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "admin_state", "disabled"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "admin_state", "disabled"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "conform_action", "transmit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "conform_action", "transmit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "conform_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "conform_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "conform_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "conform_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "exceed_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "exceed_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "exceed_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "exceed_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "exceed_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "exceed_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "excessive_burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "excessive_burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "excessive_burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "excessive_burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "mode", "bit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "mode", "bit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "peak_rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "peak_rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "peak_rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "peak_rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "sharing_mode", "dedicated"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "sharing_mode", "dedicated"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "type", "1R2C"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "type", "1R2C"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "violate_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "violate_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "violate_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "violate_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "violate_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "violate_mark_dscp", "unspecified"),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.2(2g)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigQosDppPolMinDependencyWithFvTenantAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.2(2g)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigQosDppPolMinDependencyWithFvTenantAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "admin_state", "disabled"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "admin_state", "disabled"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "conform_action", "transmit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "conform_action", "transmit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "conform_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "conform_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "conform_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "conform_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "exceed_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "exceed_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "exceed_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "exceed_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "exceed_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "exceed_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "excessive_burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "excessive_burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "excessive_burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "excessive_burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "mode", "bit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "mode", "bit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "peak_rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "peak_rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "peak_rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "peak_rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "sharing_mode", "dedicated"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "sharing_mode", "dedicated"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "type", "1R2C"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "type", "1R2C"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "violate_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "violate_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "violate_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "violate_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test", "violate_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.allow_test_2", "violate_mark_dscp", "unspecified"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "1.2(2g)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigQosDppPolMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "admin_state", "disabled"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_action", "transmit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "excessive_burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "excessive_burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "mode", "bit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "peak_rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "peak_rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "sharing_mode", "dedicated"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "type", "1R2C"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_mark_dscp", "unspecified"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigQosDppPolAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "admin_state", "disabled"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "burst_unit", "giga"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "excessive_burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "excessive_burst_unit", "giga"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "mode", "bit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "owner_tag", "owner_tag_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "peak_rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "peak_rate_unit", "giga"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "rate_unit", "giga"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "sharing_mode", "dedicated"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "type", "1R2C"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_mark_dscp", "unspecified"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigQosDppPolMinDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "name", "test_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigQosDppPolResetDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "admin_state", "disabled"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_action", "transmit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "excessive_burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "excessive_burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "mode", "bit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "peak_rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "peak_rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "sharing_mode", "dedicated"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "type", "1R2C"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_mark_dscp", "unspecified"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_data_plane_policing_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigQosDppPolChildrenDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "admin_state", "disabled"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_action", "transmit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "description", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "excessive_burst", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "excessive_burst_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "mode", "bit"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "peak_rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "peak_rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "rate", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "rate_unit", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "sharing_mode", "dedicated"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "type", "1R2C"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_action", "drop"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_mark_cos", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_mark_dscp", "unspecified"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_data_plane_policing_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigQosDppPolChildrenRemoveFromConfigDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigQosDppPolChildrenRemoveOneDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigQosDppPolChildrenRemoveAllDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "tags.#", "0"),
				),
			},
			// Update with minimum config and custom type semantic equivalent values
			{
				Config:             testConfigQosDppPolCustomTypeDependencyWithFvTenant,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "burst", "0xffffffffffffffff"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_mark_cos", "0xffff"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "conform_mark_dscp", "0xffff"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_mark_cos", "0xffff"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "exceed_mark_dscp", "0xffff"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "excessive_burst", "0xffffffffffffffff"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "name", "test_name"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_mark_cos", "0xffff"),
					resource.TestCheckResourceAttr("aci_data_plane_policing_policy.test", "violate_mark_dscp", "0xffff"),
				),
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}

const testConfigQosDppPolMinDependencyWithFvTenantAllowExisting = testConfigFvTenantMin + `
resource "aci_data_plane_policing_policy" "allow_test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
resource "aci_data_plane_policing_policy" "allow_test_2" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  depends_on = [aci_data_plane_policing_policy.allow_test]
}
`

const testConfigQosDppPolMinDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_data_plane_policing_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
`

const testConfigQosDppPolAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_data_plane_policing_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  admin_state = "disabled"
  annotation = "annotation"
  burst = "unspecified"
  burst_unit = "giga"
  conform_action = "drop"
  conform_mark_cos = "unspecified"
  conform_mark_dscp = "unspecified"
  description = "description_1"
  exceed_action = "drop"
  exceed_mark_cos = "unspecified"
  exceed_mark_dscp = "unspecified"
  excessive_burst = "unspecified"
  excessive_burst_unit = "giga"
  mode = "bit"
  name_alias = "name_alias_1"
  owner_key = "owner_key_1"
  owner_tag = "owner_tag_1"
  peak_rate = "0"
  peak_rate_unit = "giga"
  rate = "0"
  rate_unit = "giga"
  sharing_mode = "dedicated"
  type = "1R2C"
  violate_action = "drop"
  violate_mark_cos = "unspecified"
  violate_mark_dscp = "unspecified"
}
`

const testConfigQosDppPolResetDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_data_plane_policing_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  admin_state = "disabled"
  annotation = "orchestrator:terraform"
  burst = "unspecified"
  burst_unit = "unspecified"
  conform_action = "transmit"
  conform_mark_cos = "unspecified"
  conform_mark_dscp = "unspecified"
  description = ""
  exceed_action = "drop"
  exceed_mark_cos = "unspecified"
  exceed_mark_dscp = "unspecified"
  excessive_burst = "unspecified"
  excessive_burst_unit = "unspecified"
  mode = "bit"
  name_alias = ""
  owner_key = ""
  owner_tag = ""
  peak_rate = "0"
  peak_rate_unit = "unspecified"
  rate = "0"
  rate_unit = "unspecified"
  sharing_mode = "dedicated"
  type = "1R2C"
  violate_action = "drop"
  violate_mark_cos = "unspecified"
  violate_mark_dscp = "unspecified"
}
`
const testConfigQosDppPolChildrenDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_data_plane_policing_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
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

const testConfigQosDppPolChildrenRemoveFromConfigDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_data_plane_policing_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
}
`

const testConfigQosDppPolChildrenRemoveOneDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_data_plane_policing_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
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

const testConfigQosDppPolChildrenRemoveAllDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_data_plane_policing_policy" "test" {
  parent_dn = aci_tenant.test.id
  name = "test_name"
  annotations = []
  tags = []
}
`

const testConfigQosDppPolCustomTypeDependencyWithFvTenant = testConfigFvTenantMin + `
resource "aci_data_plane_policing_policy" "test" {
  parent_dn = aci_tenant.test.id
  burst = "0xffffffffffffffff"
  conform_mark_cos = "0xffff"
  conform_mark_dscp = "0xffff"
  exceed_mark_cos = "0xffff"
  exceed_mark_dscp = "0xffff"
  excessive_burst = "0xffffffffffffffff"
  name = "test_name"
  violate_mark_cos = "0xffff"
  violate_mark_dscp = "0xffff"
}
`
