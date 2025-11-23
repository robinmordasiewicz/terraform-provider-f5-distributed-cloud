---
page_title: "f5_distributed_cloud_dns_lb_health_check Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Load Balancer Health Check object is used for configuring DNS Load Balancer Health Checks
---

# f5_distributed_cloud_dns_lb_health_check (Resource)

DNS Load Balancer Health Check object is used for configuring DNS Load Balancer Health Checks

## Example Usage

```hcl
resource "f5_distributed_cloud_dns_lb_health_check" "example" {
  name        = "example-dns_lb_health_check"
  namespace   = "system"
  description = "Example DNSLBHealthCheck resource"
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

DNSLBHealthCheck can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_dns_lb_health_check.example namespace/name
```
