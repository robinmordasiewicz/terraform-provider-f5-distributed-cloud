# Example configuration for f5distributedcloud_route data source

data "f5distributedcloud_route" "example" {
  name      = "existing-route"
  namespace = "system"
}

output "route_id" {
  value = data.f5distributedcloud_route.example.id
}
