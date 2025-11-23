---
page_title: "f5_distributed_cloud_ike_phase2_profile Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  IKE Phase2 profile mainly includes the following 1. Encryption protocols to be used for IKE SA 2. Authentication Protocols to be used for IKE SA 3. DH group (if PFS is enabled) 4. Key lifetime etc
---

# f5_distributed_cloud_ike_phase2_profile (Resource)

IKE Phase2 profile mainly includes the following 1. Encryption protocols to be used for IKE SA 2. Authentication Protocols to be used for IKE SA 3. DH group (if PFS is enabled) 4. Key lifetime etc

## Example Usage

```hcl
resource "f5_distributed_cloud_ike_phase2_profile" "example" {
  name        = "example-ike_phase2_profile"
  namespace   = "system"
  description = "Example IkePhase2Profile resource"
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

IkePhase2Profile can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_ike_phase2_profile.example namespace/name
```
