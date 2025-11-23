# Example: Look up an existing Client Connection Limit
data "f5_distributed_cloud_client_connection_limit" "example" {
  name      = "my-connection-limit"
  namespace = "my-namespace"
}

output "enabled" {
  value = data.f5_distributed_cloud_client_connection_limit.example.enabled
}
