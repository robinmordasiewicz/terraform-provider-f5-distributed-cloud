---
page_title: "f5_distributed_cloud_nat_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  NAT Policy object represents the configuration of Network Address Translation parameters on F5XC Customer Edge sites. The following properties can be configured in this object:      Site where the ...
---

# f5_distributed_cloud_nat_policy (Resource)

NAT Policy object represents the configuration of Network Address Translation parameters on F5XC Customer Edge sites. The following properties can be configured in this object:      Site where the ...

## Example Usage

```hcl
resource "f5_distributed_cloud_nat_policy" "example" {
  name        = "example-nat_policy"
  namespace   = "system"
  description = "Example NATPolicy resource"
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

NATPolicy can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_nat_policy.example namespace/name
```
