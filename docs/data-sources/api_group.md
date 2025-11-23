---
page_title: "f5_distributed_cloud_api_group Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  The api_group construct provides a mechanism to classify the universal set of request APIs into a much smaller number of logical groups in order to make it easier to author and maintain API level a...
---

# f5_distributed_cloud_api_group (Data Source)

The api_group construct provides a mechanism to classify the universal set of request APIs into a much smaller number of logical groups in order to make it easier to author and maintain API level a...

## Example Usage

```hcl
data "f5_distributed_cloud_api_group" "example" {
  name      = "example-api_group"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
