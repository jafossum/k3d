#!/bin/bash

# Install istio onto Kubernetes cluster

if [[ ! -d "istio" ]]
then
    echo "./istio does not exist on your filesystem - downloading"
    /bin/bash ./download-istio.sh
else
    echo "installing istio from ./istio folder"
fi

cd istio || exit
./bin/istioctl install --set profile=demo -y

kubectl label namespace default istio-injection=enabled
