#!/bin/bash

# Apply Mongo related resources
kubectl delete -f mongo-secret.yaml
kubectl delete -f mongo-configmap.yaml
kubectl delete -f mongo.yaml
kubectl delete -f mongo-express.yaml

# Apply ingress for MongoExpress
kubectl delete -f ingress.yaml