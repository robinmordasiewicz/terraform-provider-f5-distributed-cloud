---
page_title: "f5_distributed_cloud_policer Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  * Policer objects enforces traffic rate limits * network_policy_rule and fast_acl_rule can refer to a policer. Packets   matching those rules will be subjected to the traffic contract specified in ...
---

# f5_distributed_cloud_policer (Resource)

* Policer objects enforces traffic rate limits * network_policy_rule and fast_acl_rule can refer to a policer. Packets   matching those rules will be subjected to the traffic contract specified in ...

## Example Usage

```hcl
resource "f5_distributed_cloud_policer" "example" {
  name        = "example-policer"
  namespace   = "system"
  description = "Example Policer resource"
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

Policer can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_policer.example namespace/name
```
