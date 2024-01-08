#!/bin/sh

set -e

usage() {
  echo "Usage: $0 -t <IMAGE_TAG>"
  exit 1
}

while getopts ":t:" opt; do
  case $opt in
    t)
      IMAGE_TAG="$OPTARG"
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      usage
      ;;
    :)
      echo "Option -$OPTARG requires an argument." >&2
      usage
      ;;
  esac
done

if [ -z "$IMAGE_TAG" ]; then
 # 获取当前时间，格式为 yyyyMMdd_HH:mm:ss
 IMAGE_TAG=$(date "+%Y%m%d%H%M%S")
fi

echo "The version is $IMAGE_TAG"

# 构建镜像
docker build . -t registry.cn-hangzhou.aliyuncs.com/pdfgpthz/ad_boost:"$IMAGE_TAG"
docker push registry.cn-hangzhou.aliyuncs.com/pdfgpthz/ad_boost:"$IMAGE_TAG"

