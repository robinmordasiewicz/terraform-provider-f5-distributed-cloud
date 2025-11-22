resource "f5xc_dns_lb_health_check" "example" {
  name        = "my-health-check"
  namespace   = "system"
  description = "HTTP health check for DNS load balancer"
  protocol    = "HTTP"
  port        = 80
  interval    = 30
  timeout     = 10
}

resource "f5xc_dns_lb_health_check" "https_example" {
  name        = "https-health-check"
  namespace   = "system"
  description = "HTTPS health check"
  protocol    = "HTTPS"
  port        = 443
  interval    = 60
  timeout     = 15
}
