# Example: Look up an existing Client Connection Limit
data "f5distributedcloud_client_connection_limit" "example" {
  name      = "my-connection-limit"
  namespace = "my-namespace"
}

output "enabled" {
  value = data.f5distributedcloud_client_connection_limit.example.enabled
}
