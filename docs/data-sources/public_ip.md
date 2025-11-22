---
page_title: "f5xc_public_ip Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  public_ip object represents a public IP address that is available on a set of virtual sites
---

# f5xc_public_ip (Data Source)

public_ip object represents a public IP address that is available on a set of virtual sites

## Example Usage

```hcl
data "f5xc_public_ip" "example" {
  name      = "example-public_ip"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
