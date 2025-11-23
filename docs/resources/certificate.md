---
page_title: "f5_distributed_cloud_certificate Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Certificate represents a client or server certificate.
---

# f5_distributed_cloud_certificate (Resource)

Certificate represents a client or server certificate.

## Example Usage

```hcl
resource "f5_distributed_cloud_certificate" "example" {
  name        = "example-certificate"
  namespace   = "system"
  description = "Example Certificate resource"
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

Certificate can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_certificate.example namespace/name
```
