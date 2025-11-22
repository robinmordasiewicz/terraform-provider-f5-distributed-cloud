# Example configuration for f5xc_lma_region data source

data "f5xc_lma_region" "example" {
  name      = "existing-lma_region"
  namespace = "system"
}

output "lma_region_id" {
  value = data.f5xc_lma_region.example.id
}
