# Example configuration for f5xc_shape_bot_defense_instance data source

data "f5xc_shape_bot_defense_instance" "example" {
  name      = "existing-shape_bot_defense_instance"
  namespace = "system"
}

output "shape_bot_defense_instance_id" {
  value = data.f5xc_shape_bot_defense_instance.example.id
}
