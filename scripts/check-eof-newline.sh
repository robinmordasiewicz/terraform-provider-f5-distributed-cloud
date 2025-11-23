#!/bin/bash
# Check for missing newline at end of files (CHECK ONLY - does not fix)
# Excludes docs/ (generated) and .git/

set -euo pipefail

FAILED=0

while IFS= read -r -d '' f; do
    if [ -s "$f" ]; then
        if [ "$(tail -c1 "$f" | wc -l)" -eq 0 ]; then
            echo "No newline at end: $f"
            FAILED=1
        fi
    fi
done < <(find . -type f \( -name "*.go" -o -name "*.tf" -o -name "*.sh" -o -name "*.yaml" -o -name "*.yml" \) ! -path "./.git/*" ! -path "./docs/*" -print0 2>/dev/null)

exit $FAILED
