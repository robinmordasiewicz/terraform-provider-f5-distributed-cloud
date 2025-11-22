---
page_title: "f5xc_ike1 Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  IKE Phase1 profile mainly includes the following 1. Encryption protocols to be used for IKE SA 2. Authentication Protocols to be used for IKE SA 3. DH group 4. PRF 5. Key lifetime etc
---

# f5xc_ike1 (Resource)

IKE Phase1 profile mainly includes the following 1. Encryption protocols to be used for IKE SA 2. Authentication Protocols to be used for IKE SA 3. DH group 4. PRF 5. Key lifetime etc

## Example Usage

```hcl
resource "f5xc_ike1" "example" {
  name        = "example-ike1"
  namespace   = "system"
  description = "Example Ike1 resource"
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

Ike1 can be imported using the namespace and name:

```shell
terraform import f5xc_ike1.example namespace/name
```
