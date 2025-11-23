---
page_title: "f5_distributed_cloud_site_mesh_group Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Site mesh group is a configuration tool to provide Site to Site connectivity  Set of sites in a site mesh group are defined by a virtual site.  The site mesh groups can be of type HUB or SPOKE. Whe...
---

# f5_distributed_cloud_site_mesh_group (Data Source)

Site mesh group is a configuration tool to provide Site to Site connectivity  Set of sites in a site mesh group are defined by a virtual site.  The site mesh groups can be of type HUB or SPOKE. Whe...

## Example Usage

```hcl
data "f5_distributed_cloud_site_mesh_group" "example" {
  name      = "example-site_mesh_group"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
