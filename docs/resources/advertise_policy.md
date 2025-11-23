---
page_title: "f5_distributed_cloud_advertise_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers. Basic concept is a service is made available to consumers in a given ...
---

# f5_distributed_cloud_advertise_policy (Resource)

advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers. Basic concept is a service is made available to consumers in a given ...

## Example Usage

```hcl
resource "f5_distributed_cloud_advertise_policy" "example" {
  name        = "example-advertise_policy"
  namespace   = "system"
  description = "Example AdvertisePolicy resource"
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

AdvertisePolicy can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_advertise_policy.example namespace/name
```
