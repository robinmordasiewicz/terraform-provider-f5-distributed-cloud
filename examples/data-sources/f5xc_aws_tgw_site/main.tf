# Example configuration for f5xc_aws_tgw_site data source

data "f5xc_aws_tgw_site" "example" {
  name      = "existing-aws_tgw_site"
  namespace = "system"
}

output "aws_tgw_site_id" {
  value = data.f5xc_aws_tgw_site.example.id
}
