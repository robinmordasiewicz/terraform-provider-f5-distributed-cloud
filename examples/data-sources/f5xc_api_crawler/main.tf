# Example configuration for f5xc_api_crawler data source

data "f5xc_api_crawler" "example" {
  name      = "existing-api_crawler"
  namespace = "system"
}

output "api_crawler_id" {
  value = data.f5xc_api_crawler.example.id
}
