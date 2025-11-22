---
page_title: "f5xc_ike_phase1_profile Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  IKE Phase1 profile mainly includes the following 1. Encryption protocols to be used for IKE SA 2. Authentication Protocols to be used for IKE SA 3. DH group 4. PRF 5. Key lifetime etc
---

# f5xc_ike_phase1_profile (Resource)

IKE Phase1 profile mainly includes the following 1. Encryption protocols to be used for IKE SA 2. Authentication Protocols to be used for IKE SA 3. DH group 4. PRF 5. Key lifetime etc

## Example Usage

```hcl
resource "f5xc_ike_phase1_profile" "example" {
  name        = "example-ike_phase1_profile"
  namespace   = "system"
  description = "Example IkePhase1Profile resource"
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

IkePhase1Profile can be imported using the namespace and name:

```shell
terraform import f5xc_ike_phase1_profile.example namespace/name
```
