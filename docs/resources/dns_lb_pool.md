---
page_title: "f5xc_dns_lb_pool Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Load Balancer Pool  object is used for configuring DNS Load Balancer Pool
---

# f5xc_dns_lb_pool (Resource)

DNS Load Balancer Pool  object is used for configuring DNS Load Balancer Pool

## Example Usage

```hcl
resource "f5xc_dns_lb_pool" "example" {
  name        = "example-dns_lb_pool"
  namespace   = "system"
  description = "Example DNSLBPool resource"
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

DNSLBPool can be imported using the namespace and name:

```shell
terraform import f5xc_dns_lb_pool.example namespace/name
```
