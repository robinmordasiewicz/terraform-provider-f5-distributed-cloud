# Example configuration for f5_distributed_cloud_aws_tgw_site data source

data "f5_distributed_cloud_aws_tgw_site" "example" {
  name      = "existing-aws_tgw_site"
  namespace = "system"
}

output "aws_tgw_site_id" {
  value = data.f5_distributed_cloud_aws_tgw_site.example.id
}
