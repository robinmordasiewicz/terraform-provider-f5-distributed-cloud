# Example configuration for f5_distributed_cloud_terraform_parameters data source

data "f5_distributed_cloud_terraform_parameters" "example" {
  name      = "existing-terraform_parameters"
  namespace = "system"
}

output "terraform_parameters_id" {
  value = data.f5_distributed_cloud_terraform_parameters.example.id
}
