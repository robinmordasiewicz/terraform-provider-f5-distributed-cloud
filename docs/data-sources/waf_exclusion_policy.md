---
page_title: "f5_distributed_cloud_waf_exclusion_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  WAF Exclusion Policy record
---

# f5_distributed_cloud_waf_exclusion_policy (Data Source)

WAF Exclusion Policy record

## Example Usage

```hcl
data "f5_distributed_cloud_waf_exclusion_policy" "example" {
  name      = "example-waf_exclusion_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
