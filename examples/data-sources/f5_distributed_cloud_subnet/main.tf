# Example configuration for f5_distributed_cloud_subnet data source

data "f5_distributed_cloud_subnet" "example" {
  name      = "existing-subnet"
  namespace = "system"
}

output "subnet_id" {
  value = data.f5_distributed_cloud_subnet.example.id
}
