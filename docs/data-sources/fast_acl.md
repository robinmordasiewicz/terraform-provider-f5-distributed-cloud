---
page_title: "f5xc_fast_acl Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Fast ACL provides destination and specifies rules to protect the site from denial of service attacks.   - It is specified in terms of five tuple of the packet {destination ip, destination port, sou...
---

# f5xc_fast_acl (Data Source)

Fast ACL provides destination and specifies rules to protect the site from denial of service attacks.   - It is specified in terms of five tuple of the packet {destination ip, destination port, sou...

## Example Usage

```hcl
data "f5xc_fast_acl" "example" {
  name      = "example-fast_acl"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
