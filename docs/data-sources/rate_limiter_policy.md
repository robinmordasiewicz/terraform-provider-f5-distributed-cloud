---
page_title: "f5xc_rate_limiter_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Rate limiter policy defines parameters that can be used for fine-grained control over requests for a http load balancer that are subjected to rate limiting.  It will create the following child obje...
---

# f5xc_rate_limiter_policy (Data Source)

Rate limiter policy defines parameters that can be used for fine-grained control over requests for a http load balancer that are subjected to rate limiting.  It will create the following child obje...

## Example Usage

```hcl
data "f5xc_rate_limiter_policy" "example" {
  name      = "example-rate_limiter_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
