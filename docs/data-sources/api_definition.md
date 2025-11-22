---
page_title: "f5xc_api_definition Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  The api_definition construct provides a mechanism to create api_groups based on referred OpenAPI specs. Default api_groups, which are built automatically, include a group containing all the operati...
---

# f5xc_api_definition (Data Source)

The api_definition construct provides a mechanism to create api_groups based on referred OpenAPI specs. Default api_groups, which are built automatically, include a group containing all the operati...

## Example Usage

```hcl
data "f5xc_api_definition" "example" {
  name      = "example-api_definition"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
