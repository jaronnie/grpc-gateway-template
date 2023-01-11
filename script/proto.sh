#!/bin/bash

set -e

WorkPath=$(pwd)
APP=$1

# export grpc plugin bin to PATH
PATH="${WorkPath}"/tools:"${PATH}"

function protomakefunc() {
  protos=""
  protosdir=""
  project="$1"
  for dir in apps/"$project"/proto/*; do
    if [ -d "$dir" ]; then
      if [[ "$dir" =~ apps/"$project"/proto/v[0-9]+ ]]; then
        v="${dir##*/}"
        for file in "$dir"/*; do
          if [ -d "$file" ]; then
            echo "$file"不是一个proto文件
          else
            echo "$file"
            protos="$protos ./$file"
            protosdir="$protosdir -I./$dir"
            protoc -I./"${dir}" -I./apps/"$project"/proto/ --go_out="$(pwd)"/base --go-grpc_out=require_unimplemented_servers=false:"$(pwd)"/base "$file"
          fi
        done
        protoc -I./"${dir}" -I./apps/"$project"/proto/"$v" --grpc-gateway_out=logtostderr=true:"$(pwd)"/base "${dir}"/"$project""_$v".proto
      fi
    fi
  done
}

protomakefunc "$APP"