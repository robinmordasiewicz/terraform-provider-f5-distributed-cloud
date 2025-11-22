---
page_title: "f5xc_cloud_connect Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Cloud Connect Represents connection endpoint for cloud.
---

# f5xc_cloud_connect (Data Source)

Cloud Connect Represents connection endpoint for cloud.

## Example Usage

```hcl
data "f5xc_cloud_connect" "example" {
  name      = "example-cloud_connect"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
