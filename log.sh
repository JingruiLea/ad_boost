#!/bin/bash

# 远程服务器别名
remote_alias="ali"

# 在远程服务器上获取包含pdfgpt-main的Pod信息
pod_info=$(ssh ${remote_alias} "kubectl get pods --all-namespaces | grep ad-boost")

# 从Pod信息中提取Pod名称和命名空间
namespace=$(echo $pod_info | awk '{print $1}')
pod_name=$(echo $pod_info | awk '{print $2}')

# 获取Pod日志并输出到控制台
ssh ${remote_alias} "kubectl logs -f -n $namespace $pod_name"