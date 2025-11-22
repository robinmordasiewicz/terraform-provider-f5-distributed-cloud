# Example: Look up an existing route
data "f5xc_route" "example" {
  name      = "my-route"
  namespace = "my-namespace"
}

# Output the route type
output "route_type" {
  value = data.f5xc_route.example.route_type
}

output "route_id" {
  value = data.f5xc_route.example.id
}
