---
page_title: "f5xc_contact Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Customer or tenant contact details. These details allows sending emails or regular mails to a customer. Customer can have many contacts and they come in two flavors: * billing - used for invoicing ...
---

# f5xc_contact (Data Source)

Customer or tenant contact details. These details allows sending emails or regular mails to a customer. Customer can have many contacts and they come in two flavors: * billing - used for invoicing ...

## Example Usage

```hcl
data "f5xc_contact" "example" {
  name      = "example-contact"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
