# Simple Web Server

This demo runs multiple instances of a simple web-server written in [GO](https://go.dev/).

The server receives the incomming request, and responds with the unique service pod name.

The demo runs on top of a [k3d](https://k3d.io) cluster, so an ingress exposes the web-server service through ingress rules outside the cluster.

**Make sure your k3d cluster is up and running**

## Deploy WebServer

First we need to deploy the `web-app.yaml`file. This file contains the deployment and service k8s resources for the web-server

    $ kubectl apply -f web-app.yaml

Verify that the deployemnt is running

    $ kubectl get pods

## Deploy Ingress

No services are exposed to the outside, so for this we need to deploy ingress rules.
[_(K3D needs to define ingress port mapping at creation)_](https://k3d.io/v5.2.2/usage/exposing_services/)

    kubectl apply -f ingress.yaml

## Test Setup

Now open your browser, and go to `http://localhost:8080`. The browser should now reply a success message

## Delete Demo

Clean up the resources created

    $ kubectl delete all --all

To clean up this demo only

    $ kubectl delete -f web-app.yaml
