---
page_title: "f5xc_app_api_group Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  The app_api_group construct provides a mechanism to classify the universal set of request APIs into a much smaller number of logical groups in order to make it easier to author and maintain API lev...
---

# f5xc_app_api_group (Resource)

The app_api_group construct provides a mechanism to classify the universal set of request APIs into a much smaller number of logical groups in order to make it easier to author and maintain API lev...

## Example Usage

```hcl
resource "f5xc_app_api_group" "example" {
  name        = "example-app_api_group"
  namespace   = "system"
  description = "Example AppAPIGroup resource"
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

AppAPIGroup can be imported using the namespace and name:

```shell
terraform import f5xc_app_api_group.example namespace/name
```
