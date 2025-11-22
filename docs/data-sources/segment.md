---
page_title: "f5xc_segment Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network Segment.
---

# f5xc_segment (Data Source)

Network Segment.

## Example Usage

```hcl
data "f5xc_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
