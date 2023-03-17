# Contracts

This repo contains all the contracts used in nurture.farm. 

This repo uses a [bazel](https://bazel.build/) (Version = 3.7.1) to generate proto 
definitions and service stub calls in different languages like
Go/Java/CC etc.

### Add New Contracts
Follow below steps for adding new contracts:
- Step 1: Add a new folder for the contract.
- Step 2: Add a new bazel build file named BUILD.bazel
- Step 3: Add the proto definition in above folder.
- Step 4: Add the required build steps in BUILD file like: 
    - [proto_library](https://docs.bazel.build/versions/master/be/protocol-buffer.html#proto_library)
    - [java_proto_library](https://docs.bazel.build/versions/master/be/java.html#java_proto_library)
    - [java_grpc_library](https://github.com/grpc/grpc-java/blob/master/java_grpc_library.bzl)
    - [go_proto_library](https://github.com/bazelbuild/rules_go/blob/master/proto/core.rst#go_proto_library)
    - [cc_proto_library](https://docs.bazel.build/versions/master/be/c-cpp.html#cc_proto_library) 
    - [cc_grpc_library](https://github.com/grpc/grpc/blob/master/bazel/cc_grpc_library.bzl#L7) 
- Step 5: To just generate for the currently added definition use below command
 ```shell script
bazel build //Shield:all  --sandbox_debug --verbose_failures
```
- Step 6: Above steps generates the required code inside bazel-bin folder, 
to copy the same to our folder run one or all of the copy scripts eg:
```shell script
/bin/bash copy_gen_libs.sh
```
- Step 7: Above steps copies the corresponding code libs to Gen 
folder inside the module folder we created on first step.

- Step 8: Now raise a pull request to create a new tag for the current contract.
