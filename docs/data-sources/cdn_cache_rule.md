---
page_title: "f5xc_cdn_cache_rule Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  CDN cache rule view defines a required parameters that can be used in CRUD, to create and manage CDN cache rule. It can be used to create CDN cache rule.
---

# f5xc_cdn_cache_rule (Data Source)

CDN cache rule view defines a required parameters that can be used in CRUD, to create and manage CDN cache rule. It can be used to create CDN cache rule.

## Example Usage

```hcl
data "f5xc_cdn_cache_rule" "example" {
  name      = "example-cdn_cache_rule"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
