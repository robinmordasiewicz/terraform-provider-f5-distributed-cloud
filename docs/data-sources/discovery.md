---
page_title: "f5_distributed_cloud_discovery Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Service discovery in F5XC performs following  Dynamic service discovery: Resolving the load balancer endpoints for a ADC cluster Dynamic VIP discovery: Publishing virtual IP to attract traffic from...
---

# f5_distributed_cloud_discovery (Data Source)

Service discovery in F5XC performs following  Dynamic service discovery: Resolving the load balancer endpoints for a ADC cluster Dynamic VIP discovery: Publishing virtual IP to attract traffic from...

## Example Usage

```hcl
data "f5_distributed_cloud_discovery" "example" {
  name      = "example-discovery"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
