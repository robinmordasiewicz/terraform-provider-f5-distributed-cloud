---
page_title: "f5xc_rrset Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  x-required APIs to create, update or delete individual records of a DNS zone without having to send the whole DNS zone information.
---

# f5xc_rrset (Data Source)

x-required APIs to create, update or delete individual records of a DNS zone without having to send the whole DNS zone information.

## Example Usage

```hcl
data "f5xc_rrset" "example" {
  name      = "example-rrset"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
