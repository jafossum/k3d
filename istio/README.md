# Simple Web Server

This demo runs multiple instances of a simple web-server written in [GO](https://go.dev/).

The server receives the incomming request, and responds with the unique service pod name.

The demo runs on top of a [k3d](https://k3d.io) cluster, with [Istio](https://istio.io) Service Mesh controlling the Ingress trafic.

The [Istio Ingress Gateway](https://istio.io/latest/docs/tasks/traffic-management/ingress/ingress-control/) can define traffic routing between the outside and internal k8s service routes, and this will be demoed here.

**Make sure your k3d cluster is up and running**

## Installing Istio

Run the provided util scripts for installing Istio and Kiali

    $ ./install-istio.sh
    $ ./install-kiali.sh

To launch the kiali dashboard, run the followingn in a separate terminal

    $ ./launch-kiali.sh

## Deploy the Servcie

Deploy the simple web-application using the following command

    $ kubectl apply -f web-app.yaml

Wait for the pods to run

    $ kubectl get pods
    $ kubectl get pods -w

## Apply Istio Ingress Rules

Deploy the Istio Gateway and VirtualService to get access to the service endpoints
[_(K3D needs to define ingress port mapping at creation)_](https://k3d.io/v5.2.2/usage/exposing_services/)

    $ kubectl apply -f istio-ingress.yaml

## Test Setup

Open your browser, and go to either `http://localhost:8080/status`, or `http://localhost:8080/ping`.
The browser should now reply a success message.

Looking at the VirtualService defintion, both the `/status` and `/ping` from outside the Ingress Gateway will be routed to the web-service, with the URL being just `/` when hitting the internal service endpoint.

## Delete Demo

Clean up the resources created

    $ kubectl delete all --all

To clean up this demo only

    $ kubectl delete -f istio-ingress.yaml
    $ kubectl delete -f web-app.yaml
