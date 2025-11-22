---
page_title: "f5xc_ike1 Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  IKE Phase1 profile mainly includes the following 1. Encryption protocols to be used for IKE SA 2. Authentication Protocols to be used for IKE SA 3. DH group 4. PRF 5. Key lifetime etc
---

# f5xc_ike1 (Data Source)

IKE Phase1 profile mainly includes the following 1. Encryption protocols to be used for IKE SA 2. Authentication Protocols to be used for IKE SA 3. DH group 4. PRF 5. Key lifetime etc

## Example Usage

```hcl
data "f5xc_ike1" "example" {
  name      = "example-ike1"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
