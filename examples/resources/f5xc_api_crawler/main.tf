# Example configuration for f5xc_api_crawler

resource "f5xc_api_crawler" "example" {
  name        = "example-api_crawler"
  namespace   = "system"
  description = "Example APICrawler resource managed by Terraform"

  # Add additional configuration as needed
}
