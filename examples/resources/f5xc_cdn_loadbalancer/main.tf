resource "f5xc_cdn_loadbalancer" "example" {
  name        = "my-cdn-lb"
  namespace   = "system"
  description = "CDN load balancer for static content"
  domains     = ["cdn.example.com", "static.example.com"]
}
