---
page_title: "f5xc_v1_http_monitor Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  HTTP Monitor defines an HTTP synthetic monitor.
---

# f5xc_v1_http_monitor (Resource)

HTTP Monitor defines an HTTP synthetic monitor.

## Example Usage

```hcl
resource "f5xc_v1_http_monitor" "example" {
  name        = "example-v1_http_monitor"
  namespace   = "system"
  description = "Example V1HTTPMonitor resource"
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

V1HTTPMonitor can be imported using the namespace and name:

```shell
terraform import f5xc_v1_http_monitor.example namespace/name
```
