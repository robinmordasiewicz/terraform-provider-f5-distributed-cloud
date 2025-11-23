---
page_title: "f5_distributed_cloud_network_connector Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Network Connector is used to create connection between two virtual networks on a given site.     Outside Network is the unsecured network     Inside Network is secured network  Inside network can h...
---

# f5_distributed_cloud_network_connector (Resource)

Network Connector is used to create connection between two virtual networks on a given site.     Outside Network is the unsecured network     Inside Network is secured network  Inside network can h...

## Example Usage

```hcl
resource "f5_distributed_cloud_network_connector" "example" {
  name        = "example-network_connector"
  namespace   = "system"
  description = "Example NetworkConnector resource"
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

NetworkConnector can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_network_connector.example namespace/name
```
