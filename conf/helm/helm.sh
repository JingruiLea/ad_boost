#!/bin/sh

helm upgrade nginx-ingress ingress-nginx/ingress-nginx \
--namespace ingress-nginx \
--create-namespace \
-f ./nginx-values.yaml --debug

proxychains helm install ginx-ingress ingress-nginx/ingress-nginx --namespace ingress-nginx --create-namespace -f ./nginx-values.yaml --debug