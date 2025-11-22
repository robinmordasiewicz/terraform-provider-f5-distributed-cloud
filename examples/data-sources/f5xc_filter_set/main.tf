# Example configuration for f5xc_filter_set data source

data "f5xc_filter_set" "example" {
  name      = "existing-filter_set"
  namespace = "system"
}

output "filter_set_id" {
  value = data.f5xc_filter_set.example.id
}
