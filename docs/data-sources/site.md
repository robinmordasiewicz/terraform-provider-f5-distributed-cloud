---
page_title: "f5_distributed_cloud_site Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Site represent physical/cloud cluster of volterra processing elements. There are two types of sites currently.     Regional Edge (RE)      Regional Edge sites are network edge sites owned and opera...
---

# f5_distributed_cloud_site (Data Source)

Site represent physical/cloud cluster of volterra processing elements. There are two types of sites currently.     Regional Edge (RE)      Regional Edge sites are network edge sites owned and opera...

## Example Usage

```hcl
data "f5_distributed_cloud_site" "example" {
  name      = "example-site"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
