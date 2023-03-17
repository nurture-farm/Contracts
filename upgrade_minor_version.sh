#!/bin/bash

git pull origin master || exit
git push origin master || exit

# Making a git tag for Go modules
last_tag_revision=$(git rev-list --tags --max-count=1)
tag=$(git describe --tags "$last_tag_revision")
next_tag=$tag

if [[ "$tag" =~ v([0-9]+).([0-9]+).([0-9]+)$ ]]
then
  major=${BASH_REMATCH[1]}
  minor=${BASH_REMATCH[2]}
  next_minor=$((minor + 1))
  echo "minor: minor, next_minor: $next_minor"
  next_tag="v$major.$next_minor.0"
fi

echo "Last tag: $tag, new tag: $next_tag"
git tag "$next_tag"
git push origin "$next_tag"