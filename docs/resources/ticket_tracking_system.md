---
page_title: "f5_distributed_cloud_ticket_tracking_system Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Public Custom APIs for Ticket Tracking System related operations
---

# f5_distributed_cloud_ticket_tracking_system (Resource)

Public Custom APIs for Ticket Tracking System related operations

## Example Usage

```hcl
resource "f5_distributed_cloud_ticket_tracking_system" "example" {
  name        = "example-ticket_tracking_system"
  namespace   = "system"
  description = "Example TicketTrackingSystem resource"
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

TicketTrackingSystem can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_ticket_tracking_system.example namespace/name
```
