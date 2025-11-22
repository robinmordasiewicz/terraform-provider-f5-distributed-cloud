---
page_title: "f5xc_k8s_pod_security_policy Data Source - terraform-provider-f5-distributed-cloud"
subcategory: ""
description: |-
  Pod Security Policies enable fine-grained authorization of pod creation and updates.
---

# f5xc_k8s_pod_security_policy (Data Source)

Pod Security Policies enable fine-grained authorization of pod creation and updates.

## Example Usage

```hcl
data "f5xc_k8s_pod_security_policy" "example" {
  name      = "example-k8s_pod_security_policy"
  namespace = "system"
}
```

## Argument Reference

- `name` - (Required) Name of the resource.
- `namespace` - (Required) Namespace of the resource.

## Attribute Reference

- `id` - The unique identifier for this resource.
- `description` - Description of the resource.
