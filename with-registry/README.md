# Nginx with K3D Registry

This demo shows how to use the [local registry provided with k3d](https://k3d.io/v5.3.0/usage/registries/#using-k3d-managed-registries), to host your own image registry fully integrated with the k3d cluster for testing

**Make sure your k3d cluster is up and running**

## Publish Nginx resource to the Registry

First we need to download and tag an image before pushing it to the local k3d cluster registry.

There are som addressing- and naming- tricks in this demo, in order to make things work. In the example here we use `*.localhost` in the registry name, because this gets translated automatically to `127.0.0.1` on most Linux systems. This is just to show how the registry addressing works, externaly and internaly in the cluster.

For more information on setting this up, look in the ["Using a local registry" documentation](https://k3d.io/v5.3.0/usage/registries/#using-a-local-registry)

### Resgistry Naming Trick

`*.localhost` gets replaced with `127.0.0.1` on most linux systems, so naming the image `k3d-registry.localhost:12345/nginx:latest` actually gets translated to `127.0.0.1:12345/nginx:latest`. This makes the docker push and pull commands work like magic. This is not the case on Windows and Mac systems. The reason for showing this, is that with this naming the docker commands and the internal kubectl image names are the same.

The k3d registry uses the registry name internaly, but the docker push and pull commands does not know about this naming resolution if you do not use `*.localhost` in the name, or add a lookup in the hosts file like `127.0.0.1   k3d-registry.localhost`.

**Docker commands on LINUX**

This works only because `*.localhost` gets replaced with `127.0.0.1` on most linux systems. (The docker commands in the next section will also work)

    $ docker pull nginx:latest
    $ docker tag nginx:latest k3d-registry.localhost:12345/nginx:latest
    $ docker push k3d-registry.localhost:12345/nginx:latest

**Docker commands on other systems**

    $ docker pull nginx:latest
    $ docker tag nginx:latest localhost:12345/nginx:latest
    $ docker push localhost:12345/nginx:latest

## Deploy Nginx

First we need to deploy the `nginx.yaml`file. This file contains the deployment and service k8s resources for the nginx container referencing the local k3d registry

    $ kubectl apply -f nginx.yaml

Verify that the deployemnt is running

    $ kubectl get pods

## Deploy Ingress

No services are exposed to the outside, so for this we need to deploy ingress rules.
[_(K3D needs to define ingress port mapping at creation)_](https://k3d.io/v5.3.0/usage/exposing_services/)

    kubectl apply -f ingress.yaml

## Test Setup

Now open your browser, and go to `http://localhost:8080`. The browser should now display the Nginx welcome page

## Delete Demo

Clean up the resources created

    $ kubectl delete all --all

To clean up this demo only

    $ kubectl delete -f nginx.yaml
