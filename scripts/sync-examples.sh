#!/bin/bash
# Sync examples directories with registered resources/datasources
# This ensures documentation generation doesn't fail on orphaned examples

set -e

PROVIDER_FILE="internal/provider/provider.go"

# Extract registered resource names from provider.go
echo "Extracting registered resources..."
grep -o 'NewResource().*{$' "$PROVIDER_FILE" 2>/dev/null || true
RESOURCES=$(grep '\.New.*Resource,' "$PROVIDER_FILE" | \
    grep -v '//' | \
    sed 's/.*\.//' | \
    sed 's/Resource,.*//' | \
    awk '{print tolower($0)}' | \
    sed 's/\([a-z]\)\([A-Z]\)/\1_\2/g' | \
    tr '[:upper:]' '[:lower:]' | \
    sort -u)

# Extract registered data source names from provider.go
echo "Extracting registered data sources..."
DATASOURCES=$(grep 'ds_.*\.New.*DataSource,' "$PROVIDER_FILE" | \
    grep -v '//' | \
    sed 's/.*ds_//' | \
    sed 's/\.New.*DataSource,.*//' | \
    sort -u)

echo ""
echo "Found $(echo "$RESOURCES" | wc -w | tr -d ' ') resources"
echo "Found $(echo "$DATASOURCES" | wc -w | tr -d ' ') data sources"

# Check for orphaned resource examples
echo ""
echo "Checking resource examples..."
for dir in examples/resources/f5_distributed_cloud_*/; do
    if [ -d "$dir" ]; then
        name=$(basename "$dir" | sed 's/f5_distributed_cloud_//')
        if ! echo "$RESOURCES" | grep -qw "$name" 2>/dev/null; then
            # Try alternate naming (handle camelCase to snake_case differences)
            found=false
            for res in $RESOURCES; do
                if [ "$res" = "$name" ]; then
                    found=true
                    break
                fi
            done
            if [ "$found" = false ]; then
                echo "  Orphaned: $dir"
            fi
        fi
    fi
done

# Check for orphaned data source examples
echo ""
echo "Checking data source examples..."
orphaned_ds=()
for dir in examples/data-sources/f5_distributed_cloud_*/; do
    if [ -d "$dir" ]; then
        name=$(basename "$dir" | sed 's/f5_distributed_cloud_//')
        if ! echo "$DATASOURCES" | grep -qw "$name" 2>/dev/null; then
            echo "  Orphaned: $dir"
            orphaned_ds+=("$dir")
        fi
    fi
done

echo ""
echo "Summary:"
echo "  Total orphaned data source examples: ${#orphaned_ds[@]}"
echo ""
echo "To remove orphaned examples, run:"
echo "  $0 --cleanup"

if [ "$1" = "--cleanup" ]; then
    echo ""
    echo "Removing orphaned examples..."
    for dir in "${orphaned_ds[@]}"; do
        echo "  Removing: $dir"
        rm -rf "$dir"
    done
    echo "Done."
fi
