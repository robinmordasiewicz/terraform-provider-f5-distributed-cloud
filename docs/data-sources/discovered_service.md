---
page_title: "f5_distributed_cloud_discovered_service Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Discovered Services represents the services (virtual-servers, k8s services, etc) which are discovered via the different discovery workflows.  package for Discovered Services
---

# f5_distributed_cloud_discovered_service (Data Source)

Discovered Services represents the services (virtual-servers, k8s services, etc) which are discovered via the different discovery workflows.  package for Discovered Services

## Example Usage

```hcl
data "f5_distributed_cloud_discovered_service" "example" {
  name      = "example-discovered_service"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
