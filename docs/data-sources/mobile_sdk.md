---
page_title: "f5_distributed_cloud_mobile_sdk Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Manages an F5 Distributed Cloud resource.
---

# f5_distributed_cloud_mobile_sdk (Data Source)

Manages an F5 Distributed Cloud resource.

## Example Usage

```hcl
data "f5_distributed_cloud_mobile_sdk" "example" {
  name      = "example-mobile_sdk"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
