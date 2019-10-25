#!/bin/bash

$GOPATH/src/k8s.io/code-generator/hack/update-codegen.sh all github.com/a8uhnf/rabbitmq-controller/pkg/crd/client github.com/a8uhnf/rabbitmq-controller/pkg/crd/apis a8uhnf.com:v1
echo $?