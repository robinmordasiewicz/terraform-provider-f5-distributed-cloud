---
page_title: "f5xc_app_type Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  App Type object defines a application profile type from an advanced monitoring/security point of view. An App type is a set of (micro) services that interact with one another and function as an app...
---

# f5xc_app_type (Data Source)

App Type object defines a application profile type from an advanced monitoring/security point of view. An App type is a set of (micro) services that interact with one another and function as an app...

## Example Usage

```hcl
data "f5xc_app_type" "example" {
  name      = "example-app_type"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
