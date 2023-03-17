#!/bin/bash

start=`date +%s`
. job_pool.sh

set -x -e -u

job_pool_init 8 0

bazel build //Common:all  --sandbox_debug --verbose_failures || exit
bazel build //LocationService:all  --sandbox_debug --verbose_failures || exit
bazel build //CommunicationEngine:all  --sandbox_debug --verbose_failures || exit
bazel build //AuthService:all  --sandbox_debug --verbose_failures || exit
bazel build //MessageBus:all --sandbox_debug --verbose_failures || exit
bazel build //CampaignService:all  --sandbox_debug --verbose_failures || exit
bazel build //EventPortal:all  --sandbox_debug --verbose_failures || exit
job_pool_wait

# don't forget to shut down the job pool
job_pool_shutdown
end=`date +%s`

runtime=$((end-start))
echo $runtime
# check the $job_pool_nerrors for the number of jobs that exited non-zero
echo "job_pool_nerrors: ${job_pool_nerrors}"

/bin/bash  ./copy_cc_libs_and_sources.sh || exit
/bin/bash  ./copy_go_libs_and_sources.sh || exit
/bin/bash  ./copy_java_libs_and_sources_new.sh || exit
/bin/bash  ./copy_python_libs_and_sources.sh || exit
