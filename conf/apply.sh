#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: $0 <yaml-file>"
    exit 1
fi

yaml_file="$1"

if [ ! -f "$yaml_file" ]; then
    echo "File not found: $yaml_file"
    exit 1
fi

remote=alihz1

ssh "$remote" "kubectl apply -f -" < "$yaml_file"