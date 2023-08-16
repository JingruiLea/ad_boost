#!/bin/bash

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
  echo "Image tag is required."
  usage
fi

REMOTE_HOST=ali

KUBE_CONFIG=$(cat <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: ad-boost
spec:
  replicas: 1
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
  selector:
    matchLabels:
      app: ad-boost
  template:
    metadata:
      labels:
        app: ad-boost
    spec:
      containers:
        - name: ad-boost
          image: registry.ap-southeast-1.aliyuncs.com/taimer/ad_boost:$IMAGE_TAG
          ports:
            - containerPort: 9000
          env:
            - name: TZ
              value: Asia/Shanghai
            - name: ENV
              value: production
          volumeMounts:
            - mountPath: "/opt/output/log"
              subPath: "log"
              name: log-storage
      imagePullSecrets:
        - name: aliyun-secret
      volumes:
        - name: log-storage
          persistentVolumeClaim:
            claimName: log-pvc
EOF
)

echo "Applying Kubernetes configuration on remote machine..."
echo "$KUBE_CONFIG" | ssh $REMOTE_HOST "kubectl apply -f -"

echo "Deployment and service successfully applied with image tag $IMAGE_TAG."