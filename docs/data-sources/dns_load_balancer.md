---
page_title: "f5xc_dns_load_balancer Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Load Balancer Record is used for configuring DNS Load Balancer for a record.
---

# f5xc_dns_load_balancer (Data Source)

DNS Load Balancer Record is used for configuring DNS Load Balancer for a record.

## Example Usage

```hcl
data "f5xc_dns_load_balancer" "example" {
  name      = "example-dns_load_balancer"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
