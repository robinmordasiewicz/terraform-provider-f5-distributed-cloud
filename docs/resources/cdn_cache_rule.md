---
page_title: "f5xc_cdn_cache_rule Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  CDN cache rule view defines a required parameters that can be used in CRUD, to create and manage CDN cache rule. It can be used to create CDN cache rule.
---

# f5xc_cdn_cache_rule (Resource)

CDN cache rule view defines a required parameters that can be used in CRUD, to create and manage CDN cache rule. It can be used to create CDN cache rule.

## Example Usage

```hcl
resource "f5xc_cdn_cache_rule" "example" {
  name        = "example-cdn_cache_rule"
  namespace   = "system"
  description = "Example CDNCacheRule resource"
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

CDNCacheRule can be imported using the namespace and name:

```shell
terraform import f5xc_cdn_cache_rule.example namespace/name
```
