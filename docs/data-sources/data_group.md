---
page_title: "f5xc_data_group Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A data group is a group of related items - IP addresses/subnets, strings, or integers that can be referenced in iRules.
---

# f5xc_data_group (Data Source)

A data group is a group of related items - IP addresses/subnets, strings, or integers that can be referenced in iRules.

## Example Usage

```hcl
data "f5xc_data_group" "example" {
  name      = "example-data_group"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
