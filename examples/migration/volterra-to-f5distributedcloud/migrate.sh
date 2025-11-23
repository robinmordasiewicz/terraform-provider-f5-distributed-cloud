#!/bin/bash
# Migration script: Volterra to F5 Distributed Cloud Provider
#
# This script helps migrate Terraform state from volterra_ resources
# to f5distributedcloud_ resources.
#
# Usage:
#   1. Update your .tf files to use f5distributedcloud_ prefix
#   2. Run this script to migrate state
#   3. Run terraform init && terraform plan to verify

set -e

# Resource types supported for migration (reference list)
# shellcheck disable=SC2034
RESOURCE_TYPES=(
    "namespace"
    "origin_pool"
    "http_loadbalancer"
    "tcp_loadbalancer"
    "app_firewall"
    "healthcheck"
    "virtual_site"
    "cloud_credentials"
    "aws_vpc_site"
    "azure_vnet_site"
    "gcp_vpc_site"
    "certificate"
    "dns_zone"
    "service_policy"
    "rate_limiter"
    "api_definition"
)

echo "F5 Distributed Cloud Provider Migration Script"
echo "=============================================="
echo ""

# Backup current state
echo "1. Backing up current state..."
terraform state pull > "terraform.tfstate.backup.$(date +%Y%m%d_%H%M%S)"
echo "   State backed up successfully."
echo ""

# List all volterra resources
echo "2. Finding volterra_ resources..."
VOLTERRA_RESOURCES=$(terraform state list | grep "^volterra_" || true)

if [ -z "$VOLTERRA_RESOURCES" ]; then
    echo "   No volterra_ resources found in state."
    echo "   Migration complete!"
    exit 0
fi

echo "   Found the following volterra resources:"
echo "$VOLTERRA_RESOURCES" | sed 's/^/   - /'
echo ""

# Migrate each resource
echo "3. Migrating resources..."
while IFS= read -r resource; do
    if [ -n "$resource" ]; then
        # Convert volterra_ to f5distributedcloud_
        new_resource="${resource/volterra_/f5distributedcloud_}"
        echo "   Migrating: $resource -> $new_resource"
        terraform state mv "$resource" "$new_resource" 2>/dev/null || {
            echo "   WARNING: Failed to migrate $resource (may already be migrated)"
        }
    fi
done <<< "$VOLTERRA_RESOURCES"

echo ""
echo "4. Migration complete!"
echo ""
echo "Next steps:"
echo "   1. Run 'terraform init -upgrade' to initialize with new provider"
echo "   2. Run 'terraform plan' to verify no changes needed"
echo "   3. If plan shows changes, review and adjust your configuration"
echo ""
echo "Backup state saved to: terraform.tfstate.backup.*"
