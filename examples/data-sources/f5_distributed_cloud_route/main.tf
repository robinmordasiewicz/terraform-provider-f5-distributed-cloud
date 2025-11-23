# Example configuration for f5_distributed_cloud_route data source

data "f5_distributed_cloud_route" "example" {
  name      = "existing-route"
  namespace = "system"
}

output "route_id" {
  value = data.f5_distributed_cloud_route.example.id
}
