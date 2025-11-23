---
page_title: "f5_distributed_cloud_proxy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Proxy view defines a required parameters that can be used in CRUD, to create and manage a Proxy.  View will create following child objects.  * virtual_host * advertise_policy * service_policy_set
---

# f5_distributed_cloud_proxy (Data Source)

Proxy view defines a required parameters that can be used in CRUD, to create and manage a Proxy.  View will create following child objects.  * virtual_host * advertise_policy * service_policy_set

## Example Usage

```hcl
data "f5_distributed_cloud_proxy" "example" {
  name      = "example-proxy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
