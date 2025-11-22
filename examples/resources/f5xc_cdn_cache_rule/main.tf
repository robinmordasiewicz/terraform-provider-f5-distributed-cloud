resource "f5xc_cdn_cache_rule" "static_assets" {
  name        = "static-assets-cache"
  namespace   = "system"
  description = "Cache static assets for 1 hour"
  path_match  = "/static/*"
  cache_ttl   = 3600
}

resource "f5xc_cdn_cache_rule" "images" {
  name        = "image-cache"
  namespace   = "system"
  description = "Cache images for 24 hours"
  path_match  = "/images/*"
  cache_ttl   = 86400
}
