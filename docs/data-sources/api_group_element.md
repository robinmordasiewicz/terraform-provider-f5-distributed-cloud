---
page_title: "f5_distributed_cloud_api_group_element Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A api_group_element object consists of an unordered list of HTTP methods and a path regular expression. The method and path from the input request API are matched against all elements in an api_gro...
---

# f5_distributed_cloud_api_group_element (Data Source)

A api_group_element object consists of an unordered list of HTTP methods and a path regular expression. The method and path from the input request API are matched against all elements in an api_gro...

## Example Usage

```hcl
data "f5_distributed_cloud_api_group_element" "example" {
  name      = "example-api_group_element"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
