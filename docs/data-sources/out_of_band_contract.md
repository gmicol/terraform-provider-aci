---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Contract"
layout: "aci"
page_title: "ACI: aci_out_of_band_contract"
sidebar_current: "docs-aci-data-source-aci_out_of_band_contract"
description: |-
  Data source for Out Of Band Contract
---

# aci_out_of_band_contract #

Data source for Out Of Band Contract

## API Information ##

* Class: [vzOOBBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzOOBBrCP/overview)

* Supported in ACI versions: 1.0(1e) and later.

* Distinguished Name Format: `uni/tn-mgmt/oobbrc-{name}`

## GUI Information ##

* Location: `Tenants (mgmt) -> Contracts -> Out-Of-Band Contracts`

## Example Usage ##

```hcl

data "aci_out_of_band_contract" "example" {
  name = "test_name"
}

```

## Schema ##

### Required ###

* `name` (name) - (string) The name of the Out Of Band Contract object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Out Of Band Contract object.
* `annotation` (annotation) - (string) The annotation of the Out Of Band Contract object.
* `description` (descr) - (string) The description of the Out Of Band Contract object.
* `intent` (intent) - (string) The Install Rules or Estimate Number of Rules.
* `name_alias` (nameAlias) - (string) The name alias of the Out Of Band Contract object.
* `owner_key` (ownerKey) - (string) The key for enabling clients to own their data for entity correlation.
* `owner_tag` (ownerTag) - (string) A tag for enabling clients to add their own data. For example, to indicate who created this object.
* `priority` (prio) - (string) The priority of the Out Of Band Contract object.
* `scope` (scope) - (string) Represents the scope of this contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile.
* `target_dscp` (targetDscp) - (string) The target DSCP value of the Out Of Band Contract object.

* `annotations` - (list) A list of Annotations objects ([tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.