---
page_title: "f5_distributed_cloud_recognize Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Use this API to forward API calls into the Shape APIs
---

# f5_distributed_cloud_recognize (Data Source)

Use this API to forward API calls into the Shape APIs

## Example Usage

```hcl
data "f5_distributed_cloud_recognize" "example" {
  name      = "example-recognize"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
