---
page_title: "f5xc_synthetic_monitor Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Custom handler for DNS Monitor and HTTP Monitor
---

# f5xc_synthetic_monitor (Data Source)

Custom handler for DNS Monitor and HTTP Monitor

## Example Usage

```hcl
data "f5xc_synthetic_monitor" "example" {
  name      = "example-synthetic_monitor"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
