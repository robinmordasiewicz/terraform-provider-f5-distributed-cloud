---
page_title: "f5xc_dns_load_balancer Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Load Balancer Record is used for configuring DNS Load Balancer for a record.
---

# f5xc_dns_load_balancer (Resource)

DNS Load Balancer Record is used for configuring DNS Load Balancer for a record.

## Example Usage

```hcl
resource "f5xc_dns_load_balancer" "example" {
  name        = "example-dns_load_balancer"
  namespace   = "system"
  description = "Example DNSLoadBalancer resource"
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

DNSLoadBalancer can be imported using the namespace and name:

```shell
terraform import f5xc_dns_load_balancer.example namespace/name
```
