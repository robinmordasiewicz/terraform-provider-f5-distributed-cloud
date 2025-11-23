---
page_title: "f5_distributed_cloud_nginx_instance Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  NGINX One Instance configuration
---

# f5_distributed_cloud_nginx_instance (Data Source)

NGINX One Instance configuration

## Example Usage

```hcl
data "f5_distributed_cloud_nginx_instance" "example" {
  name      = "example-nginx_instance"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
