---
page_title: "f5xc_infraprotect_deny_list_rule Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DDoS transit Deny List Rule information
---

# f5xc_infraprotect_deny_list_rule (Resource)

DDoS transit Deny List Rule information

## Example Usage

```hcl
resource "f5xc_infraprotect_deny_list_rule" "example" {
  name        = "example-infraprotect_deny_list_rule"
  namespace   = "system"
  description = "Example InfraprotectDenyListRule resource"
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

InfraprotectDenyListRule can be imported using the namespace and name:

```shell
terraform import f5xc_infraprotect_deny_list_rule.example namespace/name
```
