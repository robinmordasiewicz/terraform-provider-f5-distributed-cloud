---
page_title: "f5xc_cluster Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  cluster object represent common set of endpoints (providers of service) that can serve given route for virtual_host Cluster has list of references to Endpoint which forms the set Cluster has common...
---

# f5xc_cluster (Data Source)

cluster object represent common set of endpoints (providers of service) that can serve given route for virtual_host Cluster has list of references to Endpoint which forms the set Cluster has common...

## Example Usage

```hcl
data "f5xc_cluster" "example" {
  name      = "example-cluster"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
