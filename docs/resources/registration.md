---
page_title: "f5_distributed_cloud_registration Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  registration API(s) are used by Customer edge site to register itself with volterra edge cloud. Every node in given site is represented by it's registration. registration must have unique hostname ...
---

# f5_distributed_cloud_registration (Resource)

registration API(s) are used by Customer edge site to register itself with volterra edge cloud. Every node in given site is represented by it's registration. registration must have unique hostname ...

## Example Usage

```hcl
resource "f5_distributed_cloud_registration" "example" {
  name        = "example-registration"
  namespace   = "system"
  description = "Example Registration resource"
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

Registration can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_registration.example namespace/name
```
