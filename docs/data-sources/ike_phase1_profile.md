---
page_title: "f5_distributed_cloud_ike_phase1_profile Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  IKE Phase1 profile mainly includes the following 1. Encryption protocols to be used for IKE SA 2. Authentication Protocols to be used for IKE SA 3. DH group 4. PRF 5. Key lifetime etc
---

# f5_distributed_cloud_ike_phase1_profile (Data Source)

IKE Phase1 profile mainly includes the following 1. Encryption protocols to be used for IKE SA 2. Authentication Protocols to be used for IKE SA 3. DH group 4. PRF 5. Key lifetime etc

## Example Usage

```hcl
data "f5_distributed_cloud_ike_phase1_profile" "example" {
  name      = "example-ike_phase1_profile"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
