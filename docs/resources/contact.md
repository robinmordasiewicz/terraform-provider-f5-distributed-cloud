---
page_title: "f5xc_contact Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Customer or tenant contact details. These details allows sending emails or regular mails to a customer. Customer can have many contacts and they come in two flavors: * billing - used for invoicing ...
---

# f5xc_contact (Resource)

Customer or tenant contact details. These details allows sending emails or regular mails to a customer. Customer can have many contacts and they come in two flavors: * billing - used for invoicing ...

## Example Usage

```hcl
resource "f5xc_contact" "example" {
  name        = "example-contact"
  namespace   = "system"
  description = "Example Contact resource"
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

Contact can be imported using the namespace and name:

```shell
terraform import f5xc_contact.example namespace/name
```
