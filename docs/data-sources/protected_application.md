---
page_title: "f5_distributed_cloud_protected_application Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Configures application protected by Bot Defense
---

# f5_distributed_cloud_protected_application (Data Source)

Configures application protected by Bot Defense

## Example Usage

```hcl
data "f5_distributed_cloud_protected_application" "example" {
  name      = "example-protected_application"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
