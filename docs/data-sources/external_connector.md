---
page_title: "f5xc_external_connector Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  External Connector configuration mainly includes the following: 1. Tunnel connfiguration between F5 CE and external (3rd party) device 2. Tunnel specific parameters (IPSec, GRE)
---

# f5xc_external_connector (Data Source)

External Connector configuration mainly includes the following: 1. Tunnel connfiguration between F5 CE and external (3rd party) device 2. Tunnel specific parameters (IPSec, GRE)

## Example Usage

```hcl
data "f5xc_external_connector" "example" {
  name      = "example-external_connector"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
