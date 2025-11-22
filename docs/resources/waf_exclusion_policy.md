---
page_title: "f5xc_waf_exclusion_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  WAF Exclusion Policy record
---

# f5xc_waf_exclusion_policy (Resource)

WAF Exclusion Policy record

## Example Usage

```hcl
resource "f5xc_waf_exclusion_policy" "example" {
  name        = "example-waf_exclusion_policy"
  namespace   = "system"
  description = "Example WAFExclusionPolicy resource"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace where the resource will be created.
- `description` - (Optional) Description of the resource.
- `labels` - (Optional) Labels for the resource.
- `annotations` - (Optional) Annotations for the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.

## Import

WAFExclusionPolicy can be imported using the namespace and name:

```shell
terraform import f5xc_waf_exclusion_policy.example namespace/name
```
