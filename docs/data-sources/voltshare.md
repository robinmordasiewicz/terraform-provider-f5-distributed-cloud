---
page_title: "f5_distributed_cloud_voltshare Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  F5XC VoltShare service serves APIs for users to securing the secrets to share it with each other.
---

# f5_distributed_cloud_voltshare (Data Source)

F5XC VoltShare service serves APIs for users to securing the secrets to share it with each other.

## Example Usage

```hcl
data "f5_distributed_cloud_voltshare" "example" {
  name      = "example-voltshare"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
