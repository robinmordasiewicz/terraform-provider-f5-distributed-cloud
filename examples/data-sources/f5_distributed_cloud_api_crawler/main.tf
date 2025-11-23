# Example configuration for f5_distributed_cloud_api_crawler data source

data "f5_distributed_cloud_api_crawler" "example" {
  name      = "existing-api_crawler"
  namespace = "system"
}

output "api_crawler_id" {
  value = data.f5_distributed_cloud_api_crawler.example.id
}
