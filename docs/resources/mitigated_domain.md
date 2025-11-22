---
page_title: "f5xc_mitigated_domain Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Mitigated Domain Object defines which domains will be mitigated by Client-Side Defense. Client-Side Defense Objects is used to configure Client-Side Defense to detect/mitigate anomalous URLs for en...
---

# f5xc_mitigated_domain (Resource)

Mitigated Domain Object defines which domains will be mitigated by Client-Side Defense. Client-Side Defense Objects is used to configure Client-Side Defense to detect/mitigate anomalous URLs for en...

## Example Usage

```hcl
resource "f5xc_mitigated_domain" "example" {
  name        = "example-mitigated_domain"
  namespace   = "system"
  description = "Example MitigatedDomain resource"
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

MitigatedDomain can be imported using the namespace and name:

```shell
terraform import f5xc_mitigated_domain.example namespace/name
```
