# Example configuration for f5xc_navigation_tile data source

data "f5xc_navigation_tile" "example" {
  name      = "existing-navigation_tile"
  namespace = "system"
}

output "navigation_tile_id" {
  value = data.f5xc_navigation_tile.example.id
}
