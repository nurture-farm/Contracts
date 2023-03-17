#!/bin/bash

echo "Copying all generated libs"

/bin/bash  ./copy_cc_libs_and_sources.sh || exit
/bin/bash  ./copy_go_libs_and_sources.sh || exit
/bin/bash  ./copy_java_libs_and_sources.sh || exit
/bin/bash  ./copy_python_libs_and_sources.sh || exit
