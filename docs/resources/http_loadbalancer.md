---
page_title: "f5xc_http_loadbalancer Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  HTTP Load Balancer view defines a required parameters that can be used in CRUD, to create and manage HTTP Load Balancer. It can be used to create HTTP Load Balancer and HTTPS Load Balancer.  View w...
---

# f5xc_http_loadbalancer (Resource)

HTTP Load Balancer view defines a required parameters that can be used in CRUD, to create and manage HTTP Load Balancer. It can be used to create HTTP Load Balancer and HTTPS Load Balancer.  View w...

## Example Usage

```hcl
resource "f5xc_http_loadbalancer" "example" {
  name        = "example-http_loadbalancer"
  namespace   = "system"
  description = "Example HTTPLoadbalancer resource"
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

HTTPLoadbalancer can be imported using the namespace and name:

```shell
terraform import f5xc_http_loadbalancer.example namespace/name
```
