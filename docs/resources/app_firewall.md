---
page_title: "f5xc_app_firewall Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  WAF Configuration
---

# f5xc_app_firewall (Resource)

WAF Configuration

## Example Usage

```hcl
resource "f5xc_app_firewall" "example" {
  name        = "example-app_firewall"
  namespace   = "system"
  description = "Example AppFirewall resource"
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

AppFirewall can be imported using the namespace and name:

```shell
terraform import f5xc_app_firewall.example namespace/name
```
