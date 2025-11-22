---
page_title: "f5xc_cloud_connect Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Cloud Connect Represents connection endpoint for cloud.
---

# f5xc_cloud_connect (Resource)

Cloud Connect Represents connection endpoint for cloud.

## Example Usage

```hcl
resource "f5xc_cloud_connect" "example" {
  name        = "example-cloud_connect"
  namespace   = "system"
  description = "Example CloudConnect resource"
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

CloudConnect can be imported using the namespace and name:

```shell
terraform import f5xc_cloud_connect.example namespace/name
```
