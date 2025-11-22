---
page_title: "f5xc_rate_limiter Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A rate_limiter specifies a list of rate limit unit periods and the corresponding value of the total number of requests to be allowed in that period. It also contains an optional reference to a user...
---

# f5xc_rate_limiter (Data Source)

A rate_limiter specifies a list of rate limit unit periods and the corresponding value of the total number of requests to be allowed in that period. It also contains an optional reference to a user...

## Example Usage

```hcl
data "f5xc_rate_limiter" "example" {
  name      = "example-rate_limiter"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
