---
page_title: "f5xc_dns_lb_health_check Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Load Balancer Health Check object is used for configuring DNS Load Balancer Health Checks
---

# f5xc_dns_lb_health_check (Data Source)

DNS Load Balancer Health Check object is used for configuring DNS Load Balancer Health Checks

## Example Usage

```hcl
data "f5xc_dns_lb_health_check" "example" {
  name      = "example-dns_lb_health_check"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
