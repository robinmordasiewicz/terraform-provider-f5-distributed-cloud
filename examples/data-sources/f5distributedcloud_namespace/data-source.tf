# Copyright (c) Robin Mordasiewicz
# SPDX-License-Identifier: MPL-2.0

# Example: Look up an existing namespace
data "f5distributedcloud_namespace" "existing" {
  name = "my-namespace"
}

output "namespace_id" {
  value = data.f5distributedcloud_namespace.existing.id
}

output "namespace_description" {
  value = data.f5distributedcloud_namespace.existing.description
}

# Example: Use namespace data in a resource
resource "f5distributedcloud_origin_pool" "example" {
  name              = "example-pool"
  namespace         = data.f5distributedcloud_namespace.existing.name
  description       = "Pool in ${data.f5distributedcloud_namespace.existing.name}"
  port              = 8080
  endpoint_protocol = "http"

  origin_servers {
    public_name {
      dns_name = "backend.example.com"
    }
  }
}
