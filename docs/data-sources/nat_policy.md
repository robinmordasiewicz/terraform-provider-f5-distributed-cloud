---
page_title: "f5xc_nat_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  NAT Policy object represents the configuration of Network Address Translation parameters on F5XC Customer Edge sites. The following properties can be configured in this object:      Site where the ...
---

# f5xc_nat_policy (Data Source)

NAT Policy object represents the configuration of Network Address Translation parameters on F5XC Customer Edge sites. The following properties can be configured in this object:      Site where the ...

## Example Usage

```hcl
data "f5xc_nat_policy" "example" {
  name      = "example-nat_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
