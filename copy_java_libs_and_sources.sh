#!/bin/bash

# For java files
base_dir=${1:-$PWD}
# check for empty directory
echo "So we identified the code directory as $base_dir. \
If this is not true then pass the directory to shell script (first command line argument)"
cd $base_dir || exit

IFS=$'\n' read -r -d '' -a java_files < <(find -L ./bazel-bin -iname "lib*.jar" && printf '\0')

# Remove artifacts file if present
rm -f artifacts.txt

for java_file in "${java_files[@]}"
do
  if [[ $(basename $java_file) == lib*-hjar.jar ]] || [[ $(basename $java_file) == lib*-native-header.jar ]]
  then
    continue
  fi

#  echo "Java file: $java_file"
  new_file=$java_file
#  echo "Dirname: $(dirname $new_file)"
#  echo "Basename: $(basename $new_file)"

  full_module_name=""
  flag=false
  while ! [[ "$new_file" == *bazel-bin ]]
  do
      module_name=$(basename $new_file)
      if  $flag
      then
        full_module_name="$module_name/$full_module_name"
      fi
      flag=true
      new_file=$(dirname $new_file)
#      echo "Java Full module name: $full_module_name, Module name: $module_name, New file: $new_file"
  done
  if [[ -d "./$full_module_name" ]]
  then
    required_dir="./${full_module_name}Gen/Java$module_name"
    mkdir -p "$required_dir"
    cp -f -R "$java_file" "$required_dir/"

    #extract file name from java file path
    fileName=$(basename "$java_file")

    #if filename is not a src jar then add it to artifacts txt file
    if [[ ${fileName} != *"-src"* ]];then
      artifactId=$(echo "$fileName" | awk -v FS="(lib|(_java|-speed))" '{print $2}')
      if [[ ${artifactId} != *"_proto"* ]];then
        artifactId="${artifactId}_grpc"
      fi
      echo "$required_dir/$fileName $artifactId" >> artifacts.txt
    fi
  else
    echo "Ignoring the module: $module_name"
  fi
done

