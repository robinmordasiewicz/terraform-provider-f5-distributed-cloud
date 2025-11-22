resource "f5xc_dns_load_balancer" "example" {
  name              = "my-dns-lb"
  namespace         = "system"
  description       = "DNS load balancer for global distribution"
  domain            = "app.example.com"
  load_balance_type = "round_robin"
}
