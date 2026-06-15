#!/usr/bin/env bash
set -euo pipefail

LATEST_TAG=$(git describe --tags --abbrev=0 --match "v*" 2>/dev/null || echo "v0.0.0")
VERSION=${LATEST_TAG#v}
MAJOR=$(echo "$VERSION" | cut -d. -f1)
NEW_TAG="v$((MAJOR + 1)).0.0"

echo "🚀 Bumping MAJOR version: $LATEST_TAG -> $NEW_TAG"

git-chglog --next-tag "$NEW_TAG" -o CHANGELOG.md

if git diff --quiet CHANGELOG.md; then
    echo "⚠️  No changes found in CHANGELOG.md. Aborting."
    exit 0
fi

git add CHANGELOG.md
git commit -m "chore: release $NEW_TAG"
git tag "$NEW_TAG"
git push origin main --tags
echo "✅ Tag $NEW_TAG created and pushed."