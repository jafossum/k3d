#!/bin/bash

# exit when any command fails
set -e

# Apply Mongo related resources
kubectl apply -f mongo-secret.yaml
kubectl apply -f mongo-configmap.yaml
kubectl apply -f mongo.yaml
kubectl apply -f mongo-express.yaml

# Apply ingress for MongoExpress
kubectl apply -f ingress.yaml