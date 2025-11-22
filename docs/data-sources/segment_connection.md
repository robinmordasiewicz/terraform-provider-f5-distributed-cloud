---
page_title: "f5xc_segment_connection Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Configure a Segment Connector to allow network traffic between Segments
---

# f5xc_segment_connection (Data Source)

Configure a Segment Connector to allow network traffic between Segments

## Example Usage

```hcl
data "f5xc_segment_connection" "example" {
  name      = "example-segment_connection"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
