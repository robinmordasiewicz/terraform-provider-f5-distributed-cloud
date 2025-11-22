# Example configuration for f5xc_cdn_cache_rule data source

data "f5xc_cdn_cache_rule" "example" {
  name      = "existing-cdn_cache_rule"
  namespace = "system"
}

output "cdn_cache_rule_id" {
  value = data.f5xc_cdn_cache_rule.example.id
}
