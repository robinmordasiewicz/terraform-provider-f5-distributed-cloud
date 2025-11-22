---
page_title: "f5xc_ike2 Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  IKE Phase2 profile mainly includes the following 1. Encryption protocols to be used for IKE SA 2. Authentication Protocols to be used for IKE SA 3. DH group (if PFS is enabled) 4. Key lifetime etc
---

# f5xc_ike2 (Resource)

IKE Phase2 profile mainly includes the following 1. Encryption protocols to be used for IKE SA 2. Authentication Protocols to be used for IKE SA 3. DH group (if PFS is enabled) 4. Key lifetime etc

## Example Usage

```hcl
resource "f5xc_ike2" "example" {
  name        = "example-ike2"
  namespace   = "system"
  description = "Example Ike2 resource"
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

Ike2 can be imported using the namespace and name:

```shell
terraform import f5xc_ike2.example namespace/name
```
