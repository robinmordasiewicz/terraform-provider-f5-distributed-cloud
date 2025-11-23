---
page_title: "f5_distributed_cloud_authentication Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Authentication Object contains authentication specific config . This includes authentication type config, config specific to each authentication type and cookie specific paramters
---

# f5_distributed_cloud_authentication (Resource)

Authentication Object contains authentication specific config . This includes authentication type config, config specific to each authentication type and cookie specific paramters

## Example Usage

```hcl
resource "f5_distributed_cloud_authentication" "example" {
  name        = "example-authentication"
  namespace   = "system"
  description = "Example Authentication resource"
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

Authentication can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_authentication.example namespace/name
```
