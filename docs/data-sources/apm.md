---
page_title: "f5_distributed_cloud_apm Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BIG-IP APM Service handles the life-cycle management of BIG-IP appliances. BIG-IP APM Service manages the lifecycle of the BIG-IP appliance, which includes the functionalities like health checks, r...
---

# f5_distributed_cloud_apm (Data Source)

BIG-IP APM Service handles the life-cycle management of BIG-IP appliances. BIG-IP APM Service manages the lifecycle of the BIG-IP appliance, which includes the functionalities like health checks, r...

## Example Usage

```hcl
data "f5_distributed_cloud_apm" "example" {
  name      = "example-apm"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
