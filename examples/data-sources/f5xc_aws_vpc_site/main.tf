# Example configuration for f5xc_aws_vpc_site data source

data "f5xc_aws_vpc_site" "example" {
  name      = "existing-aws_vpc_site"
  namespace = "system"
}

output "aws_vpc_site_id" {
  value = data.f5xc_aws_vpc_site.example.id
}
