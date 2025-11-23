---
page_title: "f5_distributed_cloud_api_crawler Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  This is the api crawler type
---

# f5_distributed_cloud_api_crawler (Data Source)

This is the api crawler type

## Example Usage

```hcl
data "f5_distributed_cloud_api_crawler" "example" {
  name      = "example-api_crawler"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
