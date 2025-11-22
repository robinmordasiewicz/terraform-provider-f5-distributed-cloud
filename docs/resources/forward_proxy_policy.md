---
page_title: "f5xc_forward_proxy_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Forward Proxy policy defines access control rules for connections going via forward Proxy. It is view type of config object that uses service policy mechanism to achieve Required functionality.  Vi...
---

# f5xc_forward_proxy_policy (Resource)

Forward Proxy policy defines access control rules for connections going via forward Proxy. It is view type of config object that uses service policy mechanism to achieve Required functionality.  Vi...

## Example Usage

```hcl
resource "f5xc_forward_proxy_policy" "example" {
  name        = "example-forward_proxy_policy"
  namespace   = "system"
  description = "Example ForwardProxyPolicy resource"
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

ForwardProxyPolicy can be imported using the namespace and name:

```shell
terraform import f5xc_forward_proxy_policy.example namespace/name
```
