---
page_title: "f5xc_site_mesh_group Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Site mesh group is a configuration tool to provide Site to Site connectivity  Set of sites in a site mesh group are defined by a virtual site.  The site mesh groups can be of type HUB or SPOKE. Whe...
---

# f5xc_site_mesh_group (Resource)

Site mesh group is a configuration tool to provide Site to Site connectivity  Set of sites in a site mesh group are defined by a virtual site.  The site mesh groups can be of type HUB or SPOKE. Whe...

## Example Usage

```hcl
resource "f5xc_site_mesh_group" "example" {
  name        = "example-site_mesh_group"
  namespace   = "system"
  description = "Example SiteMeshGroup resource"
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

SiteMeshGroup can be imported using the namespace and name:

```shell
terraform import f5xc_site_mesh_group.example namespace/name
```
