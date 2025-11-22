---
page_title: "f5xc_endpoint Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Endpoint object represent the actual endpoint that provides the service (Origin Server). Sometimes due to dynamic discovery of the endpoints, single endpoint object may result in multiple actual di...
---

# f5xc_endpoint (Data Source)

Endpoint object represent the actual endpoint that provides the service (Origin Server). Sometimes due to dynamic discovery of the endpoints, single endpoint object may result in multiple actual di...

## Example Usage

```hcl
data "f5xc_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
