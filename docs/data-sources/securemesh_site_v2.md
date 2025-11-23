---
page_title: "f5_distributed_cloud_securemesh_site_v2 Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Secure Mesh site defines a required parameters that can be used in CRUD, to create and manage an Secure Mesh site.
---

# f5_distributed_cloud_securemesh_site_v2 (Data Source)

Secure Mesh site defines a required parameters that can be used in CRUD, to create and manage an Secure Mesh site.

## Example Usage

```hcl
data "f5_distributed_cloud_securemesh_site_v2" "example" {
  name      = "example-securemesh_site_v2"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
