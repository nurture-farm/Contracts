#!/bin/bash

start=`date +%s`
. job_pool.sh

display_usage() {
    echo "Publish artifacts to s3 bucket"
    printf "Usage:\n/bin/bash %s version\n" "$0"
}

job_pool_init 8 0


if [ $# -lt 1  ]; then
    display_usage
    exit 1
fi

# Upload the artifacts to s3 repo
while IFS="" read -r p || [ -n "$p"  ]
do
    IFS=' ' read -r -a file_artifact <<< "$p"
    job_pool_run /bin/bash new_upload.sh "${file_artifact[1]}" "$1" "${file_artifact[0]}" "$2"
done < artifacts.txt

job_pool_wait

# don't forget to shut down the job pool
job_pool_shutdown
end=`date +%s`

runtime=$((end-start))
echo $runtime
# check the $job_pool_nerrors for the number of jobs that exited non-zero
echo "job_pool_nerrors: ${job_pool_nerrors}"
