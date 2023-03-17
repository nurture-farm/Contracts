#!/bin/sh
set -x
display_usage() {
    echo "Upload the specified artifact to a S3 bucket."
    echo "\nUsage:\n$0 artifactId version filePath\n"
}
if [ $# -lt 3 ]; then
    display_usage
    exit 1
fi

command -v mvn >/dev/null 2>&1 || { echo >&2 "mvn is required but it's not installed.  Aborting."; exit 1; }

if [[ ( $# == "--help") || $# == "-h" ]];then
    display_usage
    exit 0
fi

REPOSITORY_ID=nurture-farm-s3-maven-release-repo
REPOSITORY_URL=s3://nurture.farm.repo.mvn.com/release
SNAPSHOT_URL=s3://nurture.farm.repo.mvn.com/snapshot

url=${REPOSITORY_URL}

if [[ $4 == "SNAPSHOT" ]]
then
  url=${SNAPSHOT_URL}
fi

echo "$url"

mvn deploy:deploy-file \
-Dpackaging=jar \
-DgroupId=farm.nurture.core.contracts \
-DartifactId=$1 \
-Dversion=$2 \
-Dfile=$3 \
-DrepositoryId=${REPOSITORY_ID} \
-Durl=${url}
