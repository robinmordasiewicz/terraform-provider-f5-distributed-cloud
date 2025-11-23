# Example configuration for f5_distributed_cloud_api_crawler

resource "f5_distributed_cloud_api_crawler" "example" {
  name        = "example-api_crawler"
  namespace   = "system"
  description = "Example APICrawler resource managed by Terraform"

  # Add additional configuration as needed
}
