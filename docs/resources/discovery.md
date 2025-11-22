---
page_title: "f5xc_discovery Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Service discovery in F5XC performs following  Dynamic service discovery: Resolving the load balancer endpoints for a ADC cluster Dynamic VIP discovery: Publishing virtual IP to attract traffic from...
---

# f5xc_discovery (Resource)

Service discovery in F5XC performs following  Dynamic service discovery: Resolving the load balancer endpoints for a ADC cluster Dynamic VIP discovery: Publishing virtual IP to attract traffic from...

## Example Usage

```hcl
resource "f5xc_discovery" "example" {
  name        = "example-discovery"
  namespace   = "system"
  description = "Example Discovery resource"
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

Discovery can be imported using the namespace and name:

```shell
terraform import f5xc_discovery.example namespace/name
```
