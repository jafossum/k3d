#!/bin/bash

VERSION_ISTO=1.12.2

curl -L https://istio.io/downloadIstio | ISTIO_VERSION=$VERSION_ISTO sh -
mv istio-$VERSION_ISTO istio
