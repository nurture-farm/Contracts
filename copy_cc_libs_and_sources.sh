#!/bin/bash

# For go files
base_dir=${1:-$PWD}
# check for empty directory
echo "So we identified the code directory as $base_dir. \
If this is not true then pass the directory to shell script (first command line argument)"
cd $base_dir || exit

IFS=$'\n' read -r -d '' -a cc_files < <(find -L ./bazel-bin \( -iname "*.pb.h" -o -iname "*.pb.cc" \) && printf '\0')

for cc_file in "${cc_files[@]}"
do


  #  echo "CC file: $cc_file"
  new_file=$cc_file
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
#      echo "CC Full module name: $full_module_name, Module name: $module_name, New file: $new_file"
  done
  if [[ -d "./$full_module_name" ]]
  then
    required_dir="./$full_module_name/Gen/CC$module_name"
    mkdir -p "$required_dir"
    cp -f -R "$cc_file" "$required_dir/"
  else
    echo "Ignoring the module: $full_module_name"
  fi
done

