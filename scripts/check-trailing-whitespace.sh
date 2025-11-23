#!/bin/bash
# Check for trailing whitespace (CHECK ONLY - does not fix)
# Excludes docs/ (generated) and .git/

set -euo pipefail

FILES=$(grep -rlP '[ \t]+$' \
    --include="*.go" \
    --include="*.tf" \
    --include="*.sh" \
    --include="*.yaml" \
    --include="*.yml" \
    --include="*.json" \
    . 2>/dev/null | grep -v "^./docs/" | grep -v "^./.git/" || true)

if [ -n "$FILES" ]; then
    echo "Files with trailing whitespace:"
    echo "$FILES"
    exit 1
fi

exit 0
