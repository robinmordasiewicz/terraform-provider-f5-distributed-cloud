# Example: F5 Distributed Cloud Provider Configuration (AFTER migration)
# This shows the migrated configuration using the f5xc provider

terraform {
  required_providers {
    f5xc = {
      source  = "robinmordasiewicz/f5-distributed-cloud"
      version = "~> 1.0"
    }
  }
}

provider "f5xc" {
  api_url  = var.api_url
  p12_file = var.api_p12_file
  # Alternative: Use token authentication
  # api_token = var.api_token
}

variable "api_p12_file" {
  description = "Path to the P12 certificate file"
  type        = string
}

variable "api_url" {
  description = "F5 XC API URL (e.g., https://tenant.console.ves.volterra.io/api)"
  type        = string
}

# Namespace resource (renamed from volterra_namespace)
resource "f5xc_namespace" "example" {
  name = "example-namespace"
}

# Origin pool resource (renamed from volterra_origin_pool)
resource "f5xc_origin_pool" "example" {
  name                   = "example-origin-pool"
  namespace              = f5xc_namespace.example.name
  endpoint_selection     = "LOCAL_PREFERRED"
  loadbalancer_algorithm = "ROUND_ROBIN"

  origin_servers {
    public_ip {
      ip = "192.168.1.1"
    }
  }

  port = 80
}

# HTTP Load Balancer resource (renamed from volterra_http_loadbalancer)
resource "f5xc_http_loadbalancer" "example" {
  name                            = "example-lb"
  namespace                       = f5xc_namespace.example.name
  domains                         = ["example.com"]
  http {
    dns_volterra_managed = false
  }

  default_route_pools {
    pool {
      name      = f5xc_origin_pool.example.name
      namespace = f5xc_namespace.example.name
    }
    weight = 1
  }

  advertise_on_public_default_vip = true
}

# App Firewall resource (renamed from volterra_app_firewall)
resource "f5xc_app_firewall" "example" {
  name      = "example-waf"
  namespace = f5xc_namespace.example.name
}
