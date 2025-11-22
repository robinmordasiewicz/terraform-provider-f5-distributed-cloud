# Example configuration for f5xc_cminstance data source

data "f5xc_cminstance" "example" {
  name      = "existing-cminstance"
  namespace = "system"
}

output "cminstance_id" {
  value = data.f5xc_cminstance.example.id
}
