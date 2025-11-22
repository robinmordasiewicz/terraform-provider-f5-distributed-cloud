---
page_title: "f5xc_app_api_group Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  The app_api_group construct provides a mechanism to classify the universal set of request APIs into a much smaller number of logical groups in order to make it easier to author and maintain API lev...
---

# f5xc_app_api_group (Data Source)

The app_api_group construct provides a mechanism to classify the universal set of request APIs into a much smaller number of logical groups in order to make it easier to author and maintain API lev...

## Example Usage

```hcl
data "f5xc_app_api_group" "example" {
  name      = "example-app_api_group"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
