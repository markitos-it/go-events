#!/usr/bin/env bash
set -euo pipefail

# 1. Calcular el nuevo tag
LATEST_TAG=$(git describe --tags --abbrev=0 --match "v*" 2>/dev/null || echo "v0.0.0")
VERSION=${LATEST_TAG#v}
MAJOR=$(echo "$VERSION" | cut -d. -f1)
MINOR=$(echo "$VERSION" | cut -d. -f2)
PATCH=$(echo "$VERSION" | cut -d. -f3)
NEW_TAG="v${MAJOR}.${MINOR}.$((PATCH + 1))"

echo "🚀 Bumping version: $LATEST_TAG -> $NEW_TAG"

git-chglog --next-tag "$NEW_TAG" -o CHANGELOG.md

if git diff --quiet CHANGELOG.md; then
    echo "⚠️  NO CHANGES ON CHANGELOG. Aborting tag creation."
    exit 0
fi

git add CHANGELOG.md
git commit -m "chore: update changelog for $NEW_TAG"

echo "🚀 Tagging and pushing: $NEW_TAG"
git tag "$NEW_TAG"
git push origin main --tags

echo "✅ Success: Tag $NEW_TAG created and pushed."