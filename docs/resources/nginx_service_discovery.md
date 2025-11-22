---
page_title: "f5xc_nginx_service_discovery Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  NGINX Service discovery in F5XC
---

# f5xc_nginx_service_discovery (Resource)

NGINX Service discovery in F5XC

## Example Usage

```hcl
resource "f5xc_nginx_service_discovery" "example" {
  name        = "example-nginx_service_discovery"
  namespace   = "system"
  description = "Example NginxServiceDiscovery resource"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace where the resource will be created.
- `description` - (Optional) Description of the resource.
- `labels` - (Optional) Labels for the resource.
- `annotations` - (Optional) Annotations for the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.

## Import

NginxServiceDiscovery can be imported using the namespace and name:

```shell
terraform import f5xc_nginx_service_discovery.example namespace/name
```
