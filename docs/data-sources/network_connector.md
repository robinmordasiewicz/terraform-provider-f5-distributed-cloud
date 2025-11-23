---
page_title: "f5_distributed_cloud_network_connector Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network Connector is used to create connection between two virtual networks on a given site.     Outside Network is the unsecured network     Inside Network is secured network  Inside network can h...
---

# f5_distributed_cloud_network_connector (Data Source)

Network Connector is used to create connection between two virtual networks on a given site.     Outside Network is the unsecured network     Inside Network is secured network  Inside network can h...

## Example Usage

```hcl
data "f5_distributed_cloud_network_connector" "example" {
  name      = "example-network_connector"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
