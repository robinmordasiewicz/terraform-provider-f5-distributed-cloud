---
page_title: "f5_distributed_cloud_virtual_site Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Virtual site object is mechanism to create arbitrary set of sites It selects all the sites for which label selector expression return true. Selector is goes thru all customer edge sites for a given...
---

# f5_distributed_cloud_virtual_site (Data Source)

Virtual site object is mechanism to create arbitrary set of sites It selects all the sites for which label selector expression return true. Selector is goes thru all customer edge sites for a given...

## Example Usage

```hcl
data "f5_distributed_cloud_virtual_site" "example" {
  name      = "example-virtual_site"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
