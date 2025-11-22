# Example: AWS TGW Site
resource "f5xc_aws_tgw_site" "example" {
  name        = "my-aws-tgw-site"
  namespace   = "system"
  description = "Example AWS TGW site for F5 XC connectivity"
}
