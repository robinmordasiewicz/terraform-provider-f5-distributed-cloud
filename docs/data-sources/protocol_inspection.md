---
page_title: "f5_distributed_cloud_protocol_inspection Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Protocol Inspection view defines the required parameters that can be used in CRUD, to create and manage Protocol Inspection. It can be used to create Protocol Inspection.  View will create the foll...
---

# f5_distributed_cloud_protocol_inspection (Data Source)

Protocol Inspection view defines the required parameters that can be used in CRUD, to create and manage Protocol Inspection. It can be used to create Protocol Inspection.  View will create the foll...

## Example Usage

```hcl
data "f5_distributed_cloud_protocol_inspection" "example" {
  name      = "example-protocol_inspection"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
