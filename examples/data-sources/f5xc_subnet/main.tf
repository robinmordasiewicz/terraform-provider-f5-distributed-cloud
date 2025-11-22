# Example configuration for f5xc_subnet data source

data "f5xc_subnet" "example" {
  name      = "existing-subnet"
  namespace = "system"
}

output "subnet_id" {
  value = data.f5xc_subnet.example.id
}
