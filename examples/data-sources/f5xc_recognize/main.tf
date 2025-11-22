# Example configuration for f5xc_recognize data source

data "f5xc_recognize" "example" {
  name      = "existing-recognize"
  namespace = "system"
}

output "recognize_id" {
  value = data.f5xc_recognize.example.id
}
