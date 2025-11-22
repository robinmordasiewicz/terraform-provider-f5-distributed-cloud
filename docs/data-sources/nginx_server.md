---
page_title: "f5xc_nginx_server Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  NGINX One Server Object configuration
---

# f5xc_nginx_server (Data Source)

NGINX One Server Object configuration

## Example Usage

```hcl
data "f5xc_nginx_server" "example" {
  name      = "example-nginx_server"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
