#!/bin/bash

# Install Kiali and the other addons and wait for them to be deployed.

kubectl apply -f istio/samples/addons
kubectl rollout status deployment/kiali -n istio-system