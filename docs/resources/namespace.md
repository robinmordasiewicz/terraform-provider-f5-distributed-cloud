---
page_title: "f5_distributed_cloud_namespace Resource - F5 Distributed Cloud"
subcategory: ""
description: |-
  Manages an F5 Distributed Cloud Namespace.
---

# f5_distributed_cloud_namespace (Resource)

Manages an F5 Distributed Cloud Namespace.

Namespaces provide logical isolation for resources within F5 XC. Resources like origin pools, HTTP load balancers, and application firewalls are created within namespaces.

## Example Usage

```terraform
resource "f5_distributed_cloud_namespace" "example" {
  name        = "my-namespace"
  description = "Example namespace for my application"
}
```

### Namespace with Labels

```terraform
resource "f5_distributed_cloud_namespace" "production" {
  name        = "production"
  description = "Production environment namespace"

  labels = {
    environment = "production"
    team        = "platform"
  }
}
```

## Argument Reference

- `name` - (Required) Name of the namespace. Must be unique across the tenant. Changing this forces a new resource.
- `description` - (Optional) Description of the namespace.
- `labels` - (Optional) Map of labels to apply to the namespace.

## Attribute Reference

- `id` - The unique identifier of the namespace.

## Import

Namespaces can be imported using the name:

```shell
terraform import f5_distributed_cloud_namespace.example my-namespace
```
