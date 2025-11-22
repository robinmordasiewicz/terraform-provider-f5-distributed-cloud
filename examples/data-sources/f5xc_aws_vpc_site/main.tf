# Example: Look up an existing AWS VPC site
data "f5xc_aws_vpc_site" "example" {
  name      = "my-aws-site"
  namespace = "system"
}

# Output the site state and AWS region
output "site_state" {
  value = data.f5xc_aws_vpc_site.example.site_state
}

output "aws_region" {
  value = data.f5xc_aws_vpc_site.example.aws_region
}
