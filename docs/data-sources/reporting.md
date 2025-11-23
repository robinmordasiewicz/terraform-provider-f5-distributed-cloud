---
page_title: "f5_distributed_cloud_reporting Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Use this API to generate reports for Shape Bot Defense
---

# f5_distributed_cloud_reporting (Data Source)

Use this API to generate reports for Shape Bot Defense

## Example Usage

```hcl
data "f5_distributed_cloud_reporting" "example" {
  name      = "example-reporting"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
