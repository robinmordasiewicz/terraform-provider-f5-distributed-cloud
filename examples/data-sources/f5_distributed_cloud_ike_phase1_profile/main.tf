# Example configuration for f5_distributed_cloud_ike_phase1_profile data source

data "f5_distributed_cloud_ike_phase1_profile" "example" {
  name      = "existing-ike_phase1_profile"
  namespace = "system"
}

output "ike_phase1_profile_id" {
  value = data.f5_distributed_cloud_ike_phase1_profile.example.id
}
