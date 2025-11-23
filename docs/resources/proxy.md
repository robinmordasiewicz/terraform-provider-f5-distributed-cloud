---
page_title: "f5_distributed_cloud_proxy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Proxy view defines a required parameters that can be used in CRUD, to create and manage a Proxy.  View will create following child objects.  * virtual_host * advertise_policy * service_policy_set
---

# f5_distributed_cloud_proxy (Resource)

Proxy view defines a required parameters that can be used in CRUD, to create and manage a Proxy.  View will create following child objects.  * virtual_host * advertise_policy * service_policy_set

## Example Usage

```hcl
resource "f5_distributed_cloud_proxy" "example" {
  name        = "example-proxy"
  namespace   = "system"
  description = "Example Proxy resource"
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

Proxy can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_proxy.example namespace/name
```
