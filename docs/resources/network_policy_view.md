---
page_title: "f5xc_network_policy_view Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network policy site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in Network policy. It can be used to either automatically create or Manually as...
---

# f5xc_network_policy_view (Resource)

Network policy site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in Network policy. It can be used to either automatically create or Manually as...

## Example Usage

```hcl
resource "f5xc_network_policy_view" "example" {
  name        = "example-network_policy_view"
  namespace   = "system"
  description = "Example NetworkPolicyView resource"
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

NetworkPolicyView can be imported using the namespace and name:

```shell
terraform import f5xc_network_policy_view.example namespace/name
```
