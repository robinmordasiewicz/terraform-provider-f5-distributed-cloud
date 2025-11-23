---
page_title: "f5_distributed_cloud_malicious_user_mitigation Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  A malicious_user_mitigation object consists of settings that specify the actions to be taken when malicious users are determined to be at different threat levels. User's activity is monitored and c...
---

# f5_distributed_cloud_malicious_user_mitigation (Resource)

A malicious_user_mitigation object consists of settings that specify the actions to be taken when malicious users are determined to be at different threat levels. User's activity is monitored and c...

## Example Usage

```hcl
resource "f5_distributed_cloud_malicious_user_mitigation" "example" {
  name        = "example-malicious_user_mitigation"
  namespace   = "system"
  description = "Example MaliciousUserMitigation resource"
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

MaliciousUserMitigation can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_malicious_user_mitigation.example namespace/name
```
