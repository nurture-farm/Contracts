#!/bin/bash
display_usage() {
    echo "Publish artifacts to s3 bucket"
    printf "Usage:\n/bin/bash %s version\n" "$0"
}

if [ $# -lt 1  ]; then
    display_usage
    exit 1
fi

# Upload the artifacts to s3 repo
while IFS="" read -r p || [ -n "$p"  ]
do
    IFS=' ' read -r -a file_artifact <<< "$p"
    /bin/bash upload.sh "${file_artifact[1]}" "$1" "${file_artifact[0]}" "$2"
done < artifacts.txt

