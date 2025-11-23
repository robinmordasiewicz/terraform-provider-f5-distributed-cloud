---
page_title: "f5_distributed_cloud_fleet Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Fleet is used to configure infrastructure components (like nodes) in one or more F5XC customer edge sites homogeneously. Fleet configuration has following information,   * Software image release to...
---

# f5_distributed_cloud_fleet (Resource)

Fleet is used to configure infrastructure components (like nodes) in one or more F5XC customer edge sites homogeneously. Fleet configuration has following information,   * Software image release to...

## Example Usage

```hcl
resource "f5_distributed_cloud_fleet" "example" {
  name        = "example-fleet"
  namespace   = "system"
  description = "Example Fleet resource"
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

Fleet can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_fleet.example namespace/name
```
