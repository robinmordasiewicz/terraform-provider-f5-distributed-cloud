# Example configuration for f5_distributed_cloud_ike_phase2_profile data source

data "f5_distributed_cloud_ike_phase2_profile" "example" {
  name      = "existing-ike_phase2_profile"
  namespace = "system"
}

output "ike_phase2_profile_id" {
  value = data.f5_distributed_cloud_ike_phase2_profile.example.id
}
