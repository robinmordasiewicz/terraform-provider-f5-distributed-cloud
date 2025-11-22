# Copyright (c) Robin Mordasiewicz
# SPDX-License-Identifier: MPL-2.0

# Example: Origin Pool with Public IP
resource "f5xc_origin_pool" "public_ip" {
  name        = "public-ip-pool"
  namespace   = "my-namespace"
  description = "Origin pool with public IP servers"

  port              = 8080
  endpoint_protocol = "http"

  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
  }

  origin_servers {
    public_ip {
      ip = "192.0.2.2"
    }
  }
}

# Example: Origin Pool with DNS Name
resource "f5xc_origin_pool" "dns_name" {
  name        = "dns-pool"
  namespace   = "my-namespace"
  description = "Origin pool with DNS-based servers"

  port              = 443
  endpoint_protocol = "https"

  origin_servers {
    public_name {
      dns_name = "api.backend.example.com"
    }
  }
}

# Example: Origin Pool with Private IP on Site
resource "f5xc_origin_pool" "private_ip" {
  name        = "private-ip-pool"
  namespace   = "my-namespace"
  description = "Origin pool with private servers on site"

  port              = 8080
  endpoint_protocol = "http"

  origin_servers {
    private_ip {
      ip             = "10.0.1.10"
      site_name      = "my-site"
      inside_network = true
    }
  }

  origin_servers {
    private_ip {
      ip             = "10.0.1.11"
      site_name      = "my-site"
      inside_network = true
    }
  }
}

# Example: Origin Pool with Custom Load Balancing
resource "f5xc_origin_pool" "custom_lb" {
  name        = "custom-lb-pool"
  namespace   = "my-namespace"
  description = "Origin pool with least request load balancing"

  port                     = 8080
  endpoint_protocol        = "http"
  loadbalancer_algorithm   = "LEAST_REQUEST"
  health_check_port        = 8081

  origin_servers {
    public_name {
      dns_name = "server1.backend.example.com"
    }
  }

  origin_servers {
    public_name {
      dns_name = "server2.backend.example.com"
    }
  }
}
