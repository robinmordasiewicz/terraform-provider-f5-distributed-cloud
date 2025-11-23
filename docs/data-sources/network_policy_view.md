---
page_title: "f5_distributed_cloud_network_policy_view Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network policy site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in Network policy. It can be used to either automatically create or Manually as...
---

# f5_distributed_cloud_network_policy_view (Data Source)

Network policy site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in Network policy. It can be used to either automatically create or Manually as...

## Example Usage

```hcl
data "f5_distributed_cloud_network_policy_view" "example" {
  name      = "example-network_policy_view"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
