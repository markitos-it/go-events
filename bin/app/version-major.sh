#!/usr/bin/env bash

set -euo pipefail


LATEST_TAG=$(git describe --tags --abbrev=0 --match "v*" 2>/dev/null || echo "v0.0.0")
VERSION=${LATEST_TAG#v}
MAJOR=$(echo "$VERSION" | cut -d. -f1)
MAJOR=$((MAJOR + 1))
MINOR=0
PATCH=0
NEW_TAG="v${MAJOR}.${MINOR}.${PATCH}"

git-chglog -o CHANGELOG.md
git add CHANGELOG.md
git commit -m "chore: update changelog for $NEW_TAG"

echo "🚀 Bumping MAJOR version: $LATEST_TAG -> $NEW_TAG"
git tag "$NEW_TAG" && git push --tags
echo "✅ Tag $NEW_TAG creado exitosamente."