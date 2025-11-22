resource "f5xc_network_connector" "example" {
  name        = "my-network-connector"
  namespace   = "system"
  description = "Connector for site-to-site connectivity"
}
