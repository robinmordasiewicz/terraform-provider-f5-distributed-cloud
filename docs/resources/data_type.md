---
page_title: "f5xc_data_type Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A data_type is defined by a set of rules. these rules include the patterns for which that data type will be detected. A data type also includes information like it's related compliances weather is ...
---

# f5xc_data_type (Resource)

A data_type is defined by a set of rules. these rules include the patterns for which that data type will be detected. A data type also includes information like it's related compliances weather is ...

## Example Usage

```hcl
resource "f5xc_data_type" "example" {
  name        = "example-data_type"
  namespace   = "system"
  description = "Example DataType resource"
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

DataType can be imported using the namespace and name:

```shell
terraform import f5xc_data_type.example namespace/name
```
