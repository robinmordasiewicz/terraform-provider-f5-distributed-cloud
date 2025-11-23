#!/bin/bash
# Check for CRLF line endings (CHECK ONLY - does not fix)
# Excludes .git/

set -euo pipefail

FILES=$(grep -rlP '\r\n' \
    --include="*.go" \
    --include="*.tf" \
    --include="*.sh" \
    --include="*.yaml" \
    --include="*.yml" \
    . 2>/dev/null | grep -v "^./.git/" || true)

if [ -n "$FILES" ]; then
    echo "Files with CRLF line endings:"
    echo "$FILES"
    exit 1
fi

exit 0
