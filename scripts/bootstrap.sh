#!/bin/bash

if [ "X" != "X$IS_PROD_RUNTIME" ] && [ "X" == "X$ENV" ]; then
    echo "[danger] on tce must set ENV, exit now."
    exit 2
fi

CURDIR=$(cd $(dirname $0); pwd)
if [ "X$1" != "X" ]; then
    RUNTIME_ROOT=$1
else
    RUNTIME_ROOT=${CURDIR}
fi

echo "CURDIR: $CURDIR"

if [ "X$RUNTIME_ROOT" == "X" ]; then
    echo "There is no RUNTIME_ROOT support."
    echo "Usage: ./bootstrap.sh $RUNTIME_ROOT"
    exit 2
fi

BinaryName="ad-boost"

export WORK_DIR=$CURDIR

echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}