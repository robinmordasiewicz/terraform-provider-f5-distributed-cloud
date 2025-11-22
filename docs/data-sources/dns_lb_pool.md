---
page_title: "f5xc_dns_lb_pool Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Load Balancer Pool  object is used for configuring DNS Load Balancer Pool
---

# f5xc_dns_lb_pool (Data Source)

DNS Load Balancer Pool  object is used for configuring DNS Load Balancer Pool

## Example Usage

```hcl
data "f5xc_dns_lb_pool" "example" {
  name      = "example-dns_lb_pool"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
