---
page_title: "f5xc_nfv_service Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  NFV Service manages the lifecycle  of the NFV appliance, which includes the functionalities like health checks, restarts, dynamic addition and deletion of NFV instances for auto scaling, defining t...
---

# f5xc_nfv_service (Data Source)

NFV Service manages the lifecycle  of the NFV appliance, which includes the functionalities like health checks, restarts, dynamic addition and deletion of NFV instances for auto scaling, defining t...

## Example Usage

```hcl
data "f5xc_nfv_service" "example" {
  name      = "example-nfv_service"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
