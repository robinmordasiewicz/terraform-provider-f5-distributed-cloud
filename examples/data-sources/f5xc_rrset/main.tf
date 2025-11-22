# Example configuration for f5xc_rrset data source

data "f5xc_rrset" "example" {
  name      = "existing-rrset"
  namespace = "system"
}

output "rrset_id" {
  value = data.f5xc_rrset.example.id
}
