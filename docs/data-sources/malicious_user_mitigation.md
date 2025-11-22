---
page_title: "f5xc_malicious_user_mitigation Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A malicious_user_mitigation object consists of settings that specify the actions to be taken when malicious users are determined to be at different threat levels. User's activity is monitored and c...
---

# f5xc_malicious_user_mitigation (Data Source)

A malicious_user_mitigation object consists of settings that specify the actions to be taken when malicious users are determined to be at different threat levels. User's activity is monitored and c...

## Example Usage

```hcl
data "f5xc_malicious_user_mitigation" "example" {
  name      = "example-malicious_user_mitigation"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
