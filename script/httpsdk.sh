#!/usr/bin/env bash

set -e

Mode=$1

function generateFunc() {
    mkdir -p pkg/myservicesdk/pb
    # shellcheck disable=SC2043
    for dir in proto/v1; do
      scopeVersion=myservicev1
      if [ -d "$dir" ]; then
        protoc --proto_path="$dir" --go_out="$(pwd)"/pkg/myservicesdk/pb --go-httpsdk_out=logtostderr=true,v=1,scopeVersion="$scopeVersion",sdkDir=pkg/myservicesdk:"$(pwd)"/pkg/myservicesdk "$dir"/*.proto
      fi
    done
    sleep 1
}

generateFunc "$Mode"