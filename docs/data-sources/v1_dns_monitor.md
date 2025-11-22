---
page_title: "f5xc_v1_dns_monitor Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Monitor defines a DNS synthetic monitor.
---

# f5xc_v1_dns_monitor (Data Source)

DNS Monitor defines a DNS synthetic monitor.

## Example Usage

```hcl
data "f5xc_v1_dns_monitor" "example" {
  name      = "example-v1_dns_monitor"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
