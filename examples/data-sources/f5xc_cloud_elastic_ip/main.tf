# Example configuration for f5xc_cloud_elastic_ip data source

data "f5xc_cloud_elastic_ip" "example" {
  name      = "existing-cloud_elastic_ip"
  namespace = "system"
}

output "cloud_elastic_ip_id" {
  value = data.f5xc_cloud_elastic_ip.example.id
}
