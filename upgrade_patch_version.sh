#!/bin/bash

set -x
curr_branch=`git branch --show-current`
git checkout master
git pull
git checkout curr_branch
git pull origin $curr_branch || exit
git push origin $curr_branch || exit

# Making a git tag for Go modules
last_tag_revision=$(git rev-list --tags --max-count=1)
tag=$(git describe --tags "$last_tag_revision")
next_tag=$tag

if [[ "$tag" =~ v([0-9]+).([0-9]+).([0-9]+)$ ]]
then
  major=${BASH_REMATCH[1]}
  minor=${BASH_REMATCH[2]}
  patch=${BASH_REMATCH[3]}
  next_patch=$((patch + 1))
  echo "minor: minor, next_minor: $next_minor"
  next_tag="v$major.$minor.$next_patch"
fi

echo "Last tag: $tag, new tag: $next_tag"
git tag "$next_tag"
git push origin "$next_tag"