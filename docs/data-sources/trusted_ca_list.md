---
page_title: "f5_distributed_cloud_trusted_ca_list Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A Root CA Certificate represents list of trusted root CAs
---

# f5_distributed_cloud_trusted_ca_list (Data Source)

A Root CA Certificate represents list of trusted root CAs

## Example Usage

```hcl
data "f5_distributed_cloud_trusted_ca_list" "example" {
  name      = "example-trusted_ca_list"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
