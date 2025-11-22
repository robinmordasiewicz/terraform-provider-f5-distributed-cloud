---
page_title: "f5xc_cloud_elastic_ip Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Cloud Elastic IP object represents a cloud elastic IP address that are created for a cloud site
---

# f5xc_cloud_elastic_ip (Data Source)

Cloud Elastic IP object represents a cloud elastic IP address that are created for a cloud site

## Example Usage

```hcl
data "f5xc_cloud_elastic_ip" "example" {
  name      = "example-cloud_elastic_ip"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
