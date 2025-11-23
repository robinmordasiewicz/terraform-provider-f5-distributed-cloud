# Example configuration for f5_distributed_cloud_certificate_chain data source

data "f5_distributed_cloud_certificate_chain" "example" {
  name      = "existing-certificate_chain"
  namespace = "system"
}

output "certificate_chain_id" {
  value = data.f5_distributed_cloud_certificate_chain.example.id
}
