---
page_title: "f5xc_api_discovery Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  The api_discovery contains settings for API discovery
---

# f5xc_api_discovery (Data Source)

The api_discovery contains settings for API discovery

## Example Usage

```hcl
data "f5xc_api_discovery" "example" {
  name      = "example-api_discovery"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
