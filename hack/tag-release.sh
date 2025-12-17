#!/bin/bash
set -u

# 1. Check Branch
BRANCH=$(git branch --show-current)
if [[ "$BRANCH" != "main" ]]; then
  echo "Current branch is $BRANCH. Release logic only runs on main."
  exit 0
fi

# 2. Get Latest Tag
# We fetch tags to ensure the runner sees them
git fetch --tags
TAG=$(git tag | sort -uV | tail -n 1)

# Default to v0.0.0 if no tags exist
if [[ -z "$TAG" ]]; then
  TAG="v0.0.0"
fi

echo "Latest tag: $TAG"

# 3. Analyze Commits
if [[ "$TAG" == "v0.0.0" ]]; then
  LOG=$(git log --pretty=format:%s)
else
  LOG=$(git log "$TAG..HEAD" --pretty=format:%s)
fi

# Use '|| true' so script doesn't crash if grep finds nothing
BREAK=$(echo "$LOG" | grep -e '!' || true)
FEAT=$(echo "$LOG" | grep -e '^feat' || true)
FIX=$(echo "$LOG" | grep -e '^fix' || true)

# 4. Calculate Next Version (Pure Bash, MacOS compatible)
VERSION=${TAG#v}
IFS='.' read -r major minor patch <<< "$VERSION"

if [[ -n "$BREAK" ]]; then
  echo "Breaking change detected."
  major=$((major+1)); minor=0; patch=0
elif [[ -n "$FEAT" ]]; then
  echo "New feature detected."
  minor=$((minor+1)); patch=0
elif [[ -n "$FIX" ]]; then
  echo "Fix detected."
  patch=$((patch+1))
else
  echo "No relevant changes (fix/feat/break). Skipping release."
  exit 0
fi

NEWTAG="v$major.$minor.$patch"
echo "New version: $NEWTAG"

if [[ "$NEWTAG" == "$TAG" ]]; then
  echo "Nothing to tag."
  exit 0
fi

# 5. Tag and Push
# Configure git if running in CI
if [[ -n "${CI:-}" ]]; then
  git config user.name "GitHub Actions"
  git config user.email "actions@github.com"
fi

echo "Tagging $NEWTAG..."
git tag "$NEWTAG"
git push origin "$NEWTAG"

# 6. Build and Package
echo "Installing Fyne..."
go install fyne.io/tools/cmd/fyne@latest

echo "Packaging App..."
# Set the internal app version to match the tag
fyne package -os darwin -icon assets/icon.png

echo "Creating DMG..."
# Clean up previous runs
rm -rf source_folder stackitStatus.dmg

mkdir source_folder
cp -r stackitStatus.app source_folder/
ln -s /Applications source_folder/Applications
hdiutil create -volname "STACKIT Status Installer" -srcfolder "./source_folder" -ov -format UDZO "stackitStatus.dmg"
rm -rf source_folder

# 7. Release to GitHub
echo "Releasing $NEWTAG to GitHub..."
# Install GH CLI if missing (GitHub Actions usually has it pre-installed, but safety first)
if ! command -v gh &> /dev/null; then
    brew install gh
fi

# Create release and upload the correct file
gh release create "$NEWTAG" "stackitStatus.dmg" \
  --title "Release $NEWTAG" \
  --notes "Automated release of $NEWTAG. Changes: $LOG"