---
page_title: "f5xc_network_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network Policy is applied to all IP packets to and from a given endpoint (called 'local_endpoint').  Local endpoint can be,   prefix : list of ip prefixes representing local endpoint   prefix_selec...
---

# f5xc_network_policy (Data Source)

Network Policy is applied to all IP packets to and from a given endpoint (called 'local_endpoint').  Local endpoint can be,   prefix : list of ip prefixes representing local endpoint   prefix_selec...

## Example Usage

```hcl
data "f5xc_network_policy" "example" {
  name      = "example-network_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
