---
page_title: "f5_distributed_cloud_virtual_site Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Virtual site object is mechanism to create arbitrary set of sites It selects all the sites for which label selector expression return true. Selector is goes thru all customer edge sites for a given...
---

# f5_distributed_cloud_virtual_site (Resource)

Virtual site object is mechanism to create arbitrary set of sites It selects all the sites for which label selector expression return true. Selector is goes thru all customer edge sites for a given...

## Example Usage

```hcl
resource "f5_distributed_cloud_virtual_site" "example" {
  name        = "example-virtual_site"
  namespace   = "system"
  description = "Example VirtualSite resource"
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

VirtualSite can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_virtual_site.example namespace/name
```
