---
page_title: "f5_distributed_cloud_dns_compliance_checks Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Compliance Checks view defines the required parameters that can be used in CRUD, to create and manage DNS Compliance Checks. It can be used to create DNS Compliance Checks.
---

# f5_distributed_cloud_dns_compliance_checks (Resource)

DNS Compliance Checks view defines the required parameters that can be used in CRUD, to create and manage DNS Compliance Checks. It can be used to create DNS Compliance Checks.

## Example Usage

```hcl
resource "f5_distributed_cloud_dns_compliance_checks" "example" {
  name        = "example-dns_compliance_checks"
  namespace   = "system"
  description = "Example DNSComplianceChecks resource"
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

DNSComplianceChecks can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_dns_compliance_checks.example namespace/name
```
