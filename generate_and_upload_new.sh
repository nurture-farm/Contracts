#!/bin/bash
set -x -e -u
/bin/bash ./build_and_copy_new.sh

# Git upload
branch="master"

git add .
git commit -m generate_and_upload_commit || true

if [[ $# -ne 1 ]]
then
 branch="master"
else
branch="$1"
fi

if [ $branch == "master" ]
then
  suffix="RELEASE"
else
  suffix="SNAPSHOT"
fi

git pull origin $branch
git push origin $branch || exit

# Making a git tag for Go modules
git fetch --tags --force
#last_tag_revision=$(git rev-list --tags --max-count=1)
#tag=$(git describe --tags "$last_tag_revision")
tag=$(git tag -l --sort=-version:refname | head -n1)
next_tag=$tag

if [[ "$tag" =~ v([0-9]+).([0-9]+).([0-9]+)([-A-Z]*)$ ]]
then
  major=${BASH_REMATCH[1]}
  minor=${BASH_REMATCH[2]}
  patch=${BASH_REMATCH[3]}
  next_patch=$((patch + 1))
  echo "patch: $patch, next_patch: $next_patch"
  next_version="$major.$minor.$next_patch-$suffix"
  next_tag="v$major.$minor.$next_patch-$suffix"
fi

echo "Last tag: $tag, new tag: $next_tag"
git tag "$next_tag"
git push origin "$next_tag"
git fetch --tags --force

/bin/bash ./new_publish.sh "${next_tag}" "${suffix}"
