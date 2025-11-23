---
page_title: "f5_distributed_cloud_nginx_csg Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  NGINX One CSG configuration
---

# f5_distributed_cloud_nginx_csg (Data Source)

NGINX One CSG configuration

## Example Usage

```hcl
data "f5_distributed_cloud_nginx_csg" "example" {
  name      = "example-nginx_csg"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
