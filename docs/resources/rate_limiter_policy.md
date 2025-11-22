---
page_title: "f5xc_rate_limiter_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Rate limiter policy defines parameters that can be used for fine-grained control over requests for a http load balancer that are subjected to rate limiting.  It will create the following child obje...
---

# f5xc_rate_limiter_policy (Resource)

Rate limiter policy defines parameters that can be used for fine-grained control over requests for a http load balancer that are subjected to rate limiting.  It will create the following child obje...

## Example Usage

```hcl
resource "f5xc_rate_limiter_policy" "example" {
  name        = "example-rate_limiter_policy"
  namespace   = "system"
  description = "Example RateLimiterPolicy resource"
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

RateLimiterPolicy can be imported using the namespace and name:

```shell
terraform import f5xc_rate_limiter_policy.example namespace/name
```
