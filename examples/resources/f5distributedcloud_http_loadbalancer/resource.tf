# Copyright (c) Robin Mordasiewicz
# SPDX-License-Identifier: MPL-2.0

# Example: Basic HTTP Load Balancer
resource "f5distributedcloud_http_loadbalancer" "basic" {
  name        = "my-http-lb"
  namespace   = "my-namespace"
  description = "Basic HTTP Load Balancer"

  domains = ["app.example.com"]

  http_type      = "http"
  advertise_port = 80
}

# Example: HTTP Load Balancer with Origin Pool
resource "f5distributedcloud_http_loadbalancer" "with_origin_pool" {
  name        = "my-http-lb-with-pool"
  namespace   = "my-namespace"
  description = "HTTP Load Balancer with Origin Pool"

  domains = ["api.example.com"]

  http_type      = "http"
  advertise_port = 80

  default_origin_pools {
    pool_name      = "my-origin-pool"
    pool_namespace = "my-namespace"
    weight         = 1
    priority       = 1
  }
}

# Example: HTTPS Load Balancer
resource "f5distributedcloud_http_loadbalancer" "https" {
  name        = "my-https-lb"
  namespace   = "my-namespace"
  description = "HTTPS Load Balancer"

  domains = ["secure.example.com"]

  http_type = "https"

  default_origin_pools {
    pool_name      = "secure-origin-pool"
    pool_namespace = "my-namespace"
    weight         = 1
    priority       = 1
  }
}

# Example: Multi-domain Load Balancer
resource "f5distributedcloud_http_loadbalancer" "multi_domain" {
  name        = "multi-domain-lb"
  namespace   = "my-namespace"
  description = "Load Balancer serving multiple domains"

  domains = [
    "www.example.com",
    "api.example.com",
    "app.example.com"
  ]

  http_type      = "http"
  advertise_port = 80
  disable_waf    = false
}

# Example: Load Balancer with WAF disabled
resource "f5distributedcloud_http_loadbalancer" "no_waf" {
  name        = "no-waf-lb"
  namespace   = "my-namespace"
  description = "Load Balancer without WAF"

  domains = ["internal.example.com"]

  http_type   = "http"
  disable_waf = true
}
