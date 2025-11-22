# Example configuration for f5xc_cdn_loadbalancer data source

data "f5xc_cdn_loadbalancer" "example" {
  name      = "existing-cdn_loadbalancer"
  namespace = "system"
}

output "cdn_loadbalancer_id" {
  value = data.f5xc_cdn_loadbalancer.example.id
}
