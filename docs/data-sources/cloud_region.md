---
page_title: "f5xc_cloud_region Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Cloud Region contains tenant specific configuration object. Users cannot create/delete these objects. They will be internally created or deleted whenever the corresponding cloud_region_region objec...
---

# f5xc_cloud_region (Data Source)

Cloud Region contains tenant specific configuration object. Users cannot create/delete these objects. They will be internally created or deleted whenever the corresponding cloud_region_region objec...

## Example Usage

```hcl
data "f5xc_cloud_region" "example" {
  name      = "example-cloud_region"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
