# TFLint configuration for terraform-provider-f5distributedcloud
# See https://github.com/terraform-linters/tflint/blob/master/docs/user-guide/config.md

config {
  # Enable module inspection
  call_module_type = "local"

  # Force to return error code
  force = false

  # Disable specific rules
  disabled_by_default = false
}

# Terraform plugin for additional rules
plugin "terraform" {
  enabled = true
  preset  = "recommended"
}

# =============================================================================
# Naming convention rules
# =============================================================================

rule "terraform_naming_convention" {
  enabled = true
  format  = "snake_case"
}

# =============================================================================
# Documentation rules
# =============================================================================

rule "terraform_documented_variables" {
  enabled = true
}

rule "terraform_documented_outputs" {
  enabled = true
}

# =============================================================================
# Code quality rules
# =============================================================================

rule "terraform_deprecated_interpolation" {
  enabled = true
}

rule "terraform_deprecated_index" {
  enabled = true
}

rule "terraform_unused_declarations" {
  enabled = true
}

rule "terraform_required_version" {
  enabled = true
}

rule "terraform_required_providers" {
  enabled = true
}

rule "terraform_typed_variables" {
  enabled = true
}

rule "terraform_standard_module_structure" {
  enabled = true
}

# =============================================================================
# Style rules
# =============================================================================

rule "terraform_comment_syntax" {
  enabled = true
}

rule "terraform_empty_list_equality" {
  enabled = true
}

# =============================================================================
# Workspace rules (disabled for provider development)
# =============================================================================

rule "terraform_workspace_remote" {
  enabled = false
}
