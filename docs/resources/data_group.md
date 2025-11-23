---
page_title: "f5_distributed_cloud_data_group Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A data group is a group of related items - IP addresses/subnets, strings, or integers that can be referenced in iRules.
---

# f5_distributed_cloud_data_group (Resource)

A data group is a group of related items - IP addresses/subnets, strings, or integers that can be referenced in iRules.

## Example Usage

```hcl
resource "f5_distributed_cloud_data_group" "example" {
  name        = "example-data_group"
  namespace   = "system"
  description = "Example DataGroup resource"
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

DataGroup can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_data_group.example namespace/name
```
