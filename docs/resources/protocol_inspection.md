---
page_title: "f5xc_protocol_inspection Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Protocol Inspection view defines the required parameters that can be used in CRUD, to create and manage Protocol Inspection. It can be used to create Protocol Inspection.  View will create the foll...
---

# f5xc_protocol_inspection (Resource)

Protocol Inspection view defines the required parameters that can be used in CRUD, to create and manage Protocol Inspection. It can be used to create Protocol Inspection.  View will create the foll...

## Example Usage

```hcl
resource "f5xc_protocol_inspection" "example" {
  name        = "example-protocol_inspection"
  namespace   = "system"
  description = "Example ProtocolInspection resource"
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

ProtocolInspection can be imported using the namespace and name:

```shell
terraform import f5xc_protocol_inspection.example namespace/name
```
