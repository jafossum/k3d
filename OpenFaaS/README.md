# OpenFaaS on k3d

## Install OpenFaaS on Kubernetes

First Install `arkade`

    # For MacOS / Linux:
    curl -SLsf https://get.arkade.dev/ | sudo sh

    # For Windows (using Git Bash)
    curl -SLsf https://get.arkade.dev/ | sh

Install `faas-cli` using `arkade`

    arkade get faas-cli

Install OpenFaas using `arkade`

    arkade install openfaas
    # For more install options
    arkade install openfaas --help

## Set up environment

To verify that openfaas has started, run:

    kubectl -n openfaas get deployments -l "release=openfaas, app=openfaas"

Forward the gateway to your machine

    kubectl rollout status -n openfaas deploy/gateway
    kubectl port-forward -n openfaas svc/gateway 8080:8080 &

If basic auth is enabled, you can now log into your gateway

Get the password

    PASSWORD=$(kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo)

Log in the `faas-cli`

    echo -n $PASSWORD | faas-cli login --username admin --password-stdin

Deploy first funtion from the store

    faas-cli store deploy figlet
    faas-cli list

Get the password for web login

    echo $PASSWORD

Open your browser at http://127.0.0.1:8080 (username: admin)

Make sure you have loged in to docker to be able to push images to the registry
