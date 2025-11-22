---
page_title: "f5xc_dns_domain Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  DNS Domain object is used for delegating DNS sub domain to volterra. It can also be used to just let volterra know about a verified sub domain that can be used for different types of load balancers...
---

# f5xc_dns_domain (Data Source)

DNS Domain object is used for delegating DNS sub domain to volterra. It can also be used to just let volterra know about a verified sub domain that can be used for different types of load balancers...

## Example Usage

```hcl
data "f5xc_dns_domain" "example" {
  name      = "example-dns_domain"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
