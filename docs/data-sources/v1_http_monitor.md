---
page_title: "f5_distributed_cloud_v1_http_monitor Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  HTTP Monitor defines an HTTP synthetic monitor.
---

# f5_distributed_cloud_v1_http_monitor (Data Source)

HTTP Monitor defines an HTTP synthetic monitor.

## Example Usage

```hcl
data "f5_distributed_cloud_v1_http_monitor" "example" {
  name      = "example-v1_http_monitor"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
