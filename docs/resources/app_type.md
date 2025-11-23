---
page_title: "f5_distributed_cloud_app_type Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  App Type object defines a application profile type from an advanced monitoring/security point of view. An App type is a set of (micro) services that interact with one another and function as an app...
---

# f5_distributed_cloud_app_type (Resource)

App Type object defines a application profile type from an advanced monitoring/security point of view. An App type is a set of (micro) services that interact with one another and function as an app...

## Example Usage

```hcl
resource "f5_distributed_cloud_app_type" "example" {
  name        = "example-app_type"
  namespace   = "system"
  description = "Example AppType resource"
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

AppType can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_app_type.example namespace/name
```
