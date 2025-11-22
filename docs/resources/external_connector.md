---
page_title: "f5xc_external_connector Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  External Connector configuration mainly includes the following: 1. Tunnel connfiguration between F5 CE and external (3rd party) device 2. Tunnel specific parameters (IPSec, GRE)
---

# f5xc_external_connector (Resource)

External Connector configuration mainly includes the following: 1. Tunnel connfiguration between F5 CE and external (3rd party) device 2. Tunnel specific parameters (IPSec, GRE)

## Example Usage

```hcl
resource "f5xc_external_connector" "example" {
  name        = "example-external_connector"
  namespace   = "system"
  description = "Example ExternalConnector resource"
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

ExternalConnector can be imported using the namespace and name:

```shell
terraform import f5xc_external_connector.example namespace/name
```
