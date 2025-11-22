---
page_title: "f5xc_client_side_defense Resource - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Custom handler in Client-Side Defense microservice will forward request(s) to backend API(s)
---

# f5xc_client_side_defense (Resource)

Custom handler in Client-Side Defense microservice will forward request(s) to backend API(s)

## Example Usage

```hcl
resource "f5xc_client_side_defense" "example" {
  name        = "example-client_side_defense"
  namespace   = "system"
  description = "Example ClientSideDefense resource"
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

ClientSideDefense can be imported using the namespace and name:

```shell
terraform import f5xc_client_side_defense.example namespace/name
```
