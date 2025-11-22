# Copyright (c) Robin Mordasiewicz
# SPDX-License-Identifier: MPL-2.0

# Example: Basic namespace creation
resource "f5xc_namespace" "example" {
  name        = "my-app-namespace"
  description = "Namespace for my application resources"
}

# Example: Namespace for a specific environment
resource "f5xc_namespace" "production" {
  name        = "production"
  description = "Production environment namespace"
}

resource "f5xc_namespace" "staging" {
  name        = "staging"
  description = "Staging environment namespace"
}

# Example: Namespace for a multi-tenant setup
resource "f5xc_namespace" "tenant_a" {
  name        = "tenant-a"
  description = "Isolated namespace for Tenant A"
}

resource "f5xc_namespace" "tenant_b" {
  name        = "tenant-b"
  description = "Isolated namespace for Tenant B"
}
