# Example configuration for f5xc_dc_cluster_group data source

data "f5xc_dc_cluster_group" "example" {
  name      = "existing-dc_cluster_group"
  namespace = "system"
}

output "dc_cluster_group_id" {
  value = data.f5xc_dc_cluster_group.example.id
}
