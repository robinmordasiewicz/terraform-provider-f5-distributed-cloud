---
page_title: "f5_distributed_cloud_allowed_domain Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Allowed Domain Object defines which domains will be allowed by Client-Side Defense Client-Side Defense Objects is used to configure Client-Side Defense to detect/mitigate anomalous URLs for enterpr...
---

# f5_distributed_cloud_allowed_domain (Resource)

Allowed Domain Object defines which domains will be allowed by Client-Side Defense Client-Side Defense Objects is used to configure Client-Side Defense to detect/mitigate anomalous URLs for enterpr...

## Example Usage

```hcl
resource "f5_distributed_cloud_allowed_domain" "example" {
  name        = "example-allowed_domain"
  namespace   = "system"
  description = "Example AllowedDomain resource"
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

AllowedDomain can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_allowed_domain.example namespace/name
```
