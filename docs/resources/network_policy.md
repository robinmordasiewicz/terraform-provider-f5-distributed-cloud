---
page_title: "f5xc_network_policy Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network Policy is applied to all IP packets to and from a given endpoint (called 'local_endpoint').  Local endpoint can be,   prefix : list of ip prefixes representing local endpoint   prefix_selec...
---

# f5xc_network_policy (Resource)

Network Policy is applied to all IP packets to and from a given endpoint (called 'local_endpoint').  Local endpoint can be,   prefix : list of ip prefixes representing local endpoint   prefix_selec...

## Example Usage

```hcl
resource "f5xc_network_policy" "example" {
  name        = "example-network_policy"
  namespace   = "system"
  description = "Example NetworkPolicy resource"
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

NetworkPolicy can be imported using the namespace and name:

```shell
terraform import f5xc_network_policy.example namespace/name
```
