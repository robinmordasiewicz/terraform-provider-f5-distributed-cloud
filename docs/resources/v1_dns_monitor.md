---
page_title: "f5_distributed_cloud_v1_dns_monitor Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Monitor defines a DNS synthetic monitor.
---

# f5_distributed_cloud_v1_dns_monitor (Resource)

DNS Monitor defines a DNS synthetic monitor.

## Example Usage

```hcl
resource "f5_distributed_cloud_v1_dns_monitor" "example" {
  name        = "example-v1_dns_monitor"
  namespace   = "system"
  description = "Example V1DNSMonitor resource"
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

V1DNSMonitor can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_v1_dns_monitor.example namespace/name
```
