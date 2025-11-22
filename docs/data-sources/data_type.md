---
page_title: "f5xc_data_type Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A data_type is defined by a set of rules. these rules include the patterns for which that data type will be detected. A data type also includes information like it's related compliances weather is ...
---

# f5xc_data_type (Data Source)

A data_type is defined by a set of rules. these rules include the patterns for which that data type will be detected. A data type also includes information like it's related compliances weather is ...

## Example Usage

```hcl
data "f5xc_data_type" "example" {
  name      = "example-data_type"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
