---
page_title: "f5_distributed_cloud_fleet Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Fleet is used to configure infrastructure components (like nodes) in one or more F5XC customer edge sites homogeneously. Fleet configuration has following information,   * Software image release to...
---

# f5_distributed_cloud_fleet (Data Source)

Fleet is used to configure infrastructure components (like nodes) in one or more F5XC customer edge sites homogeneously. Fleet configuration has following information,   * Software image release to...

## Example Usage

```hcl
data "f5_distributed_cloud_fleet" "example" {
  name      = "example-fleet"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
