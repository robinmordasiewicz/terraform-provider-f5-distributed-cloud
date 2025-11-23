---
page_title: "f5_distributed_cloud_apm Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  BIG-IP APM Service handles the life-cycle management of BIG-IP appliances. BIG-IP APM Service manages the lifecycle of the BIG-IP appliance, which includes the functionalities like health checks, r...
---

# f5_distributed_cloud_apm (Resource)

BIG-IP APM Service handles the life-cycle management of BIG-IP appliances. BIG-IP APM Service manages the lifecycle of the BIG-IP appliance, which includes the functionalities like health checks, r...

## Example Usage

```hcl
resource "f5_distributed_cloud_apm" "example" {
  name        = "example-apm"
  namespace   = "system"
  description = "Example Apm resource"
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

Apm can be imported using the namespace and name:

```shell
terraform import f5_distributed_cloud_apm.example namespace/name
```
