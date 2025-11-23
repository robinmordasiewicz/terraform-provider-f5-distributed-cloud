# Example configuration for f5distributedcloud_cdn_cache_rule

resource "f5distributedcloud_cdn_cache_rule" "example" {
  name        = "example-cdn_cache_rule"
  namespace   = "system"
  description = "Example CDNCacheRule resource managed by Terraform"

  # Add additional configuration as needed
}
