# Example: Look up an existing known label
data "f5_distributed_cloud_known_label" "example" {
  name      = "my-known-label"
  namespace = "shared"
}

# Output the label key
output "label_key" {
  value = data.f5_distributed_cloud_known_label.example.label_key
}

output "known_label_id" {
  value = data.f5_distributed_cloud_known_label.example.id
}
