resource "f5xc_udp_loadbalancer" "dns" {
  name        = "dns-lb"
  namespace   = "system"
  description = "UDP load balancer for DNS traffic"
  port        = 53
}

resource "f5xc_udp_loadbalancer" "syslog" {
  name        = "syslog-lb"
  namespace   = "system"
  description = "UDP load balancer for syslog"
  port        = 514
}
