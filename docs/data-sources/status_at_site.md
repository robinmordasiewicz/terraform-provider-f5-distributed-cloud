---
page_title: "f5_distributed_cloud_status_at_site Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Any user configured object in F5XC Edge Cloud has a status object associated with that it. An object may be created in multiple sites and therefore it is desirable to have the ability to get the cu...
---

# f5_distributed_cloud_status_at_site (Data Source)

Any user configured object in F5XC Edge Cloud has a status object associated with that it. An object may be created in multiple sites and therefore it is desirable to have the ability to get the cu...

## Example Usage

```hcl
data "f5_distributed_cloud_status_at_site" "example" {
  name      = "example-status_at_site"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
