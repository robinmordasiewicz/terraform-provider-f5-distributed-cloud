---
page_title: "f5xc_safe Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Use this API to interact with SAFE endpoints. All calls which not include user interaction allow a blob as the payload and return a blob as the return parameter.
---

# f5xc_safe (Data Source)

Use this API to interact with SAFE endpoints. All calls which not include user interaction allow a blob as the payload and return a blob as the return parameter.

## Example Usage

```hcl
data "f5xc_safe" "example" {
  name      = "example-safe"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
