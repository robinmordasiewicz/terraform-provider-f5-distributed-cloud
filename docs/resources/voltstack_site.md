---
page_title: "f5_distributed_cloud_voltstack_site Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  App Stack site defines a required parameters that can be used in CRUD, to create and manage an App Stack site.
---

# f5_distributed_cloud_voltstack_site (Resource)

App Stack site defines a required parameters that can be used in CRUD, to create and manage an App Stack site.

## Example Usage

```hcl
resource "f5_distributed_cloud_voltstack_site" "example" {
  name        = "example-voltstack_site"
  namespace   = "system"
  description = "Example VoltstackSite resource"
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

VoltstackSite can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_voltstack_site.example namespace/name
```
