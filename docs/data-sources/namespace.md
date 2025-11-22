---
page_title: "f5xc_namespace Data Source - F5 Distributed Cloud"
subcategory: ""
description: |-
  Fetches information about an existing F5 Distributed Cloud Namespace.
---

# f5xc_namespace (Data Source)

Fetches information about an existing F5 Distributed Cloud Namespace.

## Example Usage

```terraform
data "f5xc_namespace" "existing" {
  name = "my-namespace"
}

output "namespace_id" {
  value = data.f5xc_namespace.existing.id
}

# Use in a resource
resource "f5xc_origin_pool" "example" {
  name      = "my-pool"
  namespace = data.f5xc_namespace.existing.name
  # ...
}
```

## Argument Reference

- `name` - (Required) Name of the namespace to look up.

## Attribute Reference

- `id` - The unique identifier of the namespace.
- `description` - Description of the namespace.
- `labels` - Map of labels applied to the namespace.
