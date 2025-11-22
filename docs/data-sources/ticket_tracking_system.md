---
page_title: "f5xc_ticket_tracking_system Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Public Custom APIs for Ticket Tracking System related operations
---

# f5xc_ticket_tracking_system (Data Source)

Public Custom APIs for Ticket Tracking System related operations

## Example Usage

```hcl
data "f5xc_ticket_tracking_system" "example" {
  name      = "example-ticket_tracking_system"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
