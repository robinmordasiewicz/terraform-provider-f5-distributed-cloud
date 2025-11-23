---
page_title: "f5_distributed_cloud_log Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Two types of logs are supported, viz, access logs and audit logs.   * Access logs are sampled records of API calls made to a virtual host. It contains     both the request and the response data wit...
---

# f5_distributed_cloud_log (Data Source)

Two types of logs are supported, viz, access logs and audit logs.   * Access logs are sampled records of API calls made to a virtual host. It contains     both the request and the response data wit...

## Example Usage

```hcl
data "f5_distributed_cloud_log" "example" {
  name      = "example-log"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
