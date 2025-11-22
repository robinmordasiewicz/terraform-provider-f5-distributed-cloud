# Example: Volterra Provider Configuration (BEFORE migration)
# This shows a typical Volterra provider setup that will be migrated

terraform {
  required_providers {
    volterra = {
      source  = "volterraedge/volterra"
      version = "~> 0.11"
    }
  }
}

provider "volterra" {
  api_p12_file = var.api_p12_file
  url          = var.api_url
}

variable "api_p12_file" {
  description = "Path to the P12 certificate file"
  type        = string
}

variable "api_url" {
  description = "F5 XC API URL (e.g., https://tenant.console.ves.volterra.io/api)"
  type        = string
}

# Namespace resource
resource "volterra_namespace" "example" {
  name = "example-namespace"
}

# Origin pool resource
resource "volterra_origin_pool" "example" {
  name                   = "example-origin-pool"
  namespace              = volterra_namespace.example.name
  endpoint_selection     = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"

  origin_servers {
    public_ip {
      ip = "192.168.1.1"
    }
  }

  port = 80
}

# HTTP Load Balancer resource
resource "volterra_http_loadbalancer" "example" {
  name                            = "example-lb"
  namespace                       = volterra_namespace.example.name
  domains                         = ["example.com"]
  http {
    dns_volterra_managed = false
  }

  default_route_pools {
    pool {
      name      = volterra_origin_pool.example.name
      namespace = volterra_namespace.example.name
    }
    weight = 1
  }

  advertise_on_public_default_vip = true
}

# App Firewall resource
resource "volterra_app_firewall" "example" {
  name      = "example-waf"
  namespace = volterra_namespace.example.name
}
