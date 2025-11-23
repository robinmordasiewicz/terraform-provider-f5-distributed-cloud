---
page_title: "f5_distributed_cloud_endpoint Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Endpoint object represent the actual endpoint that provides the service (Origin Server). Sometimes due to dynamic discovery of the endpoints, single endpoint object may result in multiple actual di...
---

# f5_distributed_cloud_endpoint (Resource)

Endpoint object represent the actual endpoint that provides the service (Origin Server). Sometimes due to dynamic discovery of the endpoints, single endpoint object may result in multiple actual di...

## Example Usage

```hcl
resource "f5_distributed_cloud_endpoint" "example" {
  name        = "example-endpoint"
  namespace   = "system"
  description = "Example Endpoint resource"
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

Endpoint can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_endpoint.example namespace/name
```
