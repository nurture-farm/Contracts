#!/bin/bash

# For go files
base_dir=${1:-$PWD}
# check for empty directory
echo "So we identified the code directory as $base_dir. \
If this is not true then pass the directory to shell script (first command line argument)"
cd $base_dir || exit

IFS=$'\n' read -r -d '' -a py_files < <(find -L ./bazel-bin \( -iname "*_pb2.py" -o -iname "*_pb2_grpc.py" \) && printf '\0')

for py_file in "${py_files[@]}"
do


  #  echo "Python file: $py_file"
  new_file=$py_file
#  echo "Dirname: $(dirname $new_file)"
#  echo "Basename: $(basename $new_file)"

  flag=false
  full_module_name=""
  while ! [[ "$new_file" == *bazel-bin ]]
  do
      module_name=$(basename $new_file)
      if  $flag
      then
        full_module_name="$module_name/$full_module_name"
      fi
      flag=true
      new_file=$(dirname $new_file)
#      echo "Python Full module name: $full_module_name, Module name: $module_name, New file: $new_file"
  done
  if [[ -d "./$full_module_name" ]]
  then
    required_dir="./$full_module_name/Gen/Python$module_name"
    mkdir -p "$required_dir"
    cp -f -R "$py_file" "$required_dir/"
  else
    echo "Ignoring the module: $full_module_name"
  fi
done

