---
page_title: "f5_distributed_cloud_nginx_service_discovery Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  NGINX Service discovery in F5XC
---

# f5_distributed_cloud_nginx_service_discovery (Data Source)

NGINX Service discovery in F5XC

## Example Usage

```hcl
data "f5_distributed_cloud_nginx_service_discovery" "example" {
  name      = "example-nginx_service_discovery"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
