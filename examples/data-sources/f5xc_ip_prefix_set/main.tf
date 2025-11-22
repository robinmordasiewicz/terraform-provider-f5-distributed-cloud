# Example: Look up an existing IP prefix set
data "f5xc_ip_prefix_set" "example" {
  name      = "my-prefix-set"
  namespace = "my-namespace"
}

# Output the IP prefixes
output "prefixes" {
  value = data.f5xc_ip_prefix_set.example.prefixes
}

output "prefix_set_id" {
  value = data.f5xc_ip_prefix_set.example.id
}
