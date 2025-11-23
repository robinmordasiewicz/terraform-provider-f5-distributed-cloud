# Example: Look up an existing service mesh
data "f5distributedcloud_service_mesh" "example" {
  name      = "my-service-mesh"
  namespace = "my-namespace"
}

# Output the mesh type
output "mesh_type" {
  value = data.f5distributedcloud_service_mesh.example.mesh_type
}

output "service_mesh_id" {
  value = data.f5distributedcloud_service_mesh.example.id
}
