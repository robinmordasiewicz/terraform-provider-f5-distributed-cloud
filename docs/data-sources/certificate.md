---
page_title: "f5_distributed_cloud_certificate Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Certificate represents a client or server certificate.
---

# f5_distributed_cloud_certificate (Data Source)

Certificate represents a client or server certificate.

## Example Usage

```hcl
data "f5_distributed_cloud_certificate" "example" {
  name      = "example-certificate"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
