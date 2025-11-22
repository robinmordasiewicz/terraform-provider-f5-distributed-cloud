---
page_title: "f5xc_dns_compliance_checks Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Compliance Checks view defines the required parameters that can be used in CRUD, to create and manage DNS Compliance Checks. It can be used to create DNS Compliance Checks.
---

# f5xc_dns_compliance_checks (Data Source)

DNS Compliance Checks view defines the required parameters that can be used in CRUD, to create and manage DNS Compliance Checks. It can be used to create DNS Compliance Checks.

## Example Usage

```hcl
data "f5xc_dns_compliance_checks" "example" {
  name      = "example-dns_compliance_checks"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
