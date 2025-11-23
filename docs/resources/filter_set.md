---
page_title: "f5_distributed_cloud_filter_set Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Filter Set is a set of saved filtering criteria used in the Console. This allows users to declare named sets of filters so that they can be consistently used and shared to quickly reactivate a part...
---

# f5_distributed_cloud_filter_set (Resource)

Filter Set is a set of saved filtering criteria used in the Console. This allows users to declare named sets of filters so that they can be consistently used and shared to quickly reactivate a part...

## Example Usage

```hcl
resource "f5_distributed_cloud_filter_set" "example" {
  name        = "example-filter_set"
  namespace   = "system"
  description = "Example FilterSet resource"
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

FilterSet can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_filter_set.example namespace/name
```
