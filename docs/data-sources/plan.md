---
page_title: "f5_distributed_cloud_plan Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Usage plan related RPCs. Used for billing and onboarding.
---

# f5_distributed_cloud_plan (Data Source)

Usage plan related RPCs. Used for billing and onboarding.

## Example Usage

```hcl
data "f5_distributed_cloud_plan" "example" {
  name      = "example-plan"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
