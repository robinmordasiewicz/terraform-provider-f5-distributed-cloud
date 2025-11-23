# Example: Virtual Site Data Source
# This example retrieves information about an existing Virtual Site

data "f5distributedcloud_virtual_site" "example" {
  name      = "my-virtual-site"
  namespace = "my-namespace"
}

# Output the virtual site details
output "vsite_id" {
  description = "ID of the Virtual Site"
  value       = data.f5distributedcloud_virtual_site.example.id
}

output "vsite_description" {
  description = "Description of the Virtual Site"
  value       = data.f5distributedcloud_virtual_site.example.description
}

output "vsite_type" {
  description = "Type of the Virtual Site"
  value       = data.f5distributedcloud_virtual_site.example.site_type
}

# Use the data source to reference an existing virtual site
# Virtual sites are commonly used to group physical sites for policy application
resource "f5distributedcloud_http_loadbalancer" "example" {
  name        = "my-lb"
  namespace   = data.f5distributedcloud_virtual_site.example.namespace
  description = "Load balancer deployed to ${data.f5distributedcloud_virtual_site.example.name}"

  domains = ["app.example.com"]

  http {
    port = 80
  }

  advertise_custom {
    advertise_where {
      virtual_site {
        name      = data.f5distributedcloud_virtual_site.example.name
        namespace = data.f5distributedcloud_virtual_site.example.namespace
      }
    }
  }

  default_route_pools {
    pool {
      name      = "backend-pool"
      namespace = data.f5distributedcloud_virtual_site.example.namespace
    }
    weight = 1
  }
}
