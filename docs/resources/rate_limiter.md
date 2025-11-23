---
page_title: "f5_distributed_cloud_rate_limiter Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A rate_limiter specifies a list of rate limit unit periods and the corresponding value of the total number of requests to be allowed in that period. It also contains an optional reference to a user...
---

# f5_distributed_cloud_rate_limiter (Resource)

A rate_limiter specifies a list of rate limit unit periods and the corresponding value of the total number of requests to be allowed in that period. It also contains an optional reference to a user...

## Example Usage

```hcl
resource "f5_distributed_cloud_rate_limiter" "example" {
  name        = "example-rate_limiter"
  namespace   = "system"
  description = "Example RateLimiter resource"
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

RateLimiter can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_rate_limiter.example namespace/name
```
