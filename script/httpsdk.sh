#!/usr/bin/env bash

set -e

WorkPath=$(pwd)

# export tools bin to PATH
PATH="${WorkPath}"/tools:"${PATH}"

function generateFunc() {
    mkdir -p httpsdk/"$1"/pb
    # shellcheck disable=SC2043
    for dir in proto/v1; do
      scopeVersion="$1"v1
      if [ -d "$dir" ]; then
          protoc --proto_path="$dir" --go_out="$(pwd)"/httpsdk/"$1"/pb --go-httpsdk_out=logtostderr=true,v=1,scopeVersion="$scopeVersion",sdkDir=httpsdk/"$1":"$(pwd)"/httpsdk/"$1" "$dir"/*.proto
      fi
    done
}

generateFunc myservice