# Copyright (c) Robin Mordasiewicz
# SPDX-License-Identifier: MPL-2.0

# Example: Look up an existing origin pool
data "f5xc_origin_pool" "existing" {
  name      = "backend-pool"
  namespace = "my-namespace"
}

output "origin_pool_port" {
  value = data.f5xc_origin_pool.existing.port
}

output "origin_pool_protocol" {
  value = data.f5xc_origin_pool.existing.endpoint_protocol
}

output "origin_pool_algorithm" {
  value = data.f5xc_origin_pool.existing.loadbalancer_algorithm
}

# Example: Use origin pool data in an HTTP load balancer
resource "f5xc_http_loadbalancer" "example" {
  name        = "example-lb"
  namespace   = data.f5xc_origin_pool.existing.namespace
  description = "Load balancer using existing pool"
  domains     = ["app.example.com"]
  http_port   = 80

  advertise_on_public = true

  default_pool {
    name      = data.f5xc_origin_pool.existing.name
    namespace = data.f5xc_origin_pool.existing.namespace
  }
}
