---
page_title: "f5xc_token Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  token object is used to manage site admission. User must generate token before provisioning and pass this token to site during it's registration. Invalid tokens are refused and site with invalid to...
---

# f5xc_token (Resource)

token object is used to manage site admission. User must generate token before provisioning and pass this token to site during it's registration. Invalid tokens are refused and site with invalid to...

## Example Usage

```hcl
resource "f5xc_token" "example" {
  name        = "example-token"
  namespace   = "system"
  description = "Example Token resource"
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

Token can be imported using the namespace and name:

```shell
terraform import f5xc_token.example namespace/name
```
