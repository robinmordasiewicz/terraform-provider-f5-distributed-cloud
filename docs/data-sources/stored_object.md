---
page_title: "f5_distributed_cloud_stored_object Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Manages an F5 Distributed Cloud resource.
---

# f5_distributed_cloud_stored_object (Data Source)

Manages an F5 Distributed Cloud resource.

## Example Usage

```hcl
data "f5_distributed_cloud_stored_object" "example" {
  name      = "example-stored_object"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
