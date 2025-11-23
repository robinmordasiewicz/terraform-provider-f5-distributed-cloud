---
page_title: "f5_distributed_cloud_wifi Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Proto definitions for runtime WIFI info on sites.
---

# f5_distributed_cloud_wifi (Data Source)

Proto definitions for runtime WIFI info on sites.

## Example Usage

```hcl
data "f5_distributed_cloud_wifi" "example" {
  name      = "example-wifi"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
