# Example configuration for f5_distributed_cloud_rrset data source

data "f5_distributed_cloud_rrset" "example" {
  name      = "existing-rrset"
  namespace = "system"
}

output "rrset_id" {
  value = data.f5_distributed_cloud_rrset.example.id
}
