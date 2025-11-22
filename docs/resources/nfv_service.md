---
page_title: "f5xc_nfv_service Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  NFV Service manages the lifecycle  of the NFV appliance, which includes the functionalities like health checks, restarts, dynamic addition and deletion of NFV instances for auto scaling, defining t...
---

# f5xc_nfv_service (Resource)

NFV Service manages the lifecycle  of the NFV appliance, which includes the functionalities like health checks, restarts, dynamic addition and deletion of NFV instances for auto scaling, defining t...

## Example Usage

```hcl
resource "f5xc_nfv_service" "example" {
  name        = "example-nfv_service"
  namespace   = "system"
  description = "Example NfvService resource"
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

NfvService can be imported using the namespace and name:

```shell
terraform import f5xc_nfv_service.example namespace/name
```
