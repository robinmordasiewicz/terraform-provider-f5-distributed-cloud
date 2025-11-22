---
page_title: "f5xc_api_definition Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  The api_definition construct provides a mechanism to create api_groups based on referred OpenAPI specs. Default api_groups, which are built automatically, include a group containing all the operati...
---

# f5xc_api_definition (Resource)

The api_definition construct provides a mechanism to create api_groups based on referred OpenAPI specs. Default api_groups, which are built automatically, include a group containing all the operati...

## Example Usage

```hcl
resource "f5xc_api_definition" "example" {
  name        = "example-api_definition"
  namespace   = "system"
  description = "Example APIDefinition resource"
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

APIDefinition can be imported using the namespace and name:

```shell
terraform import f5xc_api_definition.example namespace/name
```
