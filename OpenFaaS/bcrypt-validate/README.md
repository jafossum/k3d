# GO Bcrypt Validate OpenFaaS Function

GO has a very nice library for encrypting passwords ands strings using bcrypt.
The resulting hashed string can be stored in clear text, and will be sendt together
with the clear text password for validation.

This Validate functionality is now wrapped in a OpenFaaS function

## JSON Body

The Validate function only accepts a POST request with the following JSON body

```json
{
  "text": "some-passwd",
  "hash": "$2a$10$jmIre8ZqAYenwc5y8gTNe.DMTrbuVjCwoc/g6LPvmjcX3tcqtzbvK"
}
```

# Create an OpenFaas Function

Create a folder for your function

    mkdir bcrypt-validate && cd bcrypt-validate

Get the golang-middleware template

    faas-cli template store pull golang-middleware

## Create the function template

Create a new OpenFaaS function using the golang-middleware template

    faas-cli new bcrypt-validate --lang golang-middleware --prefix=<YOUR-DOCKER-ID>

## GO Modules

To make the repository a compatible GO project, use go modules or vendoring

    go mod init <YOR GO RPEO NAME>
    go mod tidy

## Push your code to OpenFaaS

Push your code to OpenFaaS

    faas-cli up -f bcrypt-validate.yml

## Synchronously invoke your function

Invoke your function through the OpenFaaS web console, curl, or with faas-cli
(Remember to put a `\` before each \$ in the resulting hash in the JSON request)

    curl http://127.0.0.1:8080/function/bcrypt-validate -d "{\"text\":\"some-passwd\",\"hash\":\"\$2a\$10\$vftrptYx21rjd3v09Vhmn.Gw3dsDDxzY6RD.hA10KfgCuYltV9.wq\"}"

## Asynchronously invoce your function

For this to work you need a webserver listening on port `:8888` on your IP.
in the `OpenFaas/web` there is small GO function for testing this.

    cd ../web
    go run web.go

Invoke your function with curl by calling `/async-function/bcrypt-validate`

    curl http://127.0.0.1:8080/async-function/bcrypt-validate -d "{\"text\":\"some-passwd\",\"hash\":\"\$2a\$10\$vftrptYx21rjd3v09Vhmn.Gw3dsDDxzY6RD.hA10KfgCuYltV9.wq\"}" -H "X-Callback-Url: http://<YOUR-IP>:8888"

### Test with Docker

If OpenFaaS is not running, you can test it using the resulting docker container

    docker run --rm -p 8080:8080 jafossum/bcrypt-validate:latest

Invoke with curl (Remember to put a `\` before each \$ in the resulting hash in the JSON request)

    curl -d "{\"text\":\"some-passwd\",\"hash\":\"\$2a\$10\$vftrptYx21rjd3v09Vhmn.Gw3dsDDxzY6RD.hA10KfgCuYltV9.wq\"}" localhost:8080

## Read the Logs

All OpenFaaS functions run in the `openfaas-fn` namespace

    kubectl get pods -n openfaas-fn
    kubectl logs bcrypt-validate-d6c8cdf6d-rnchn -n openfaas-fn

## Remove the function from OpenFaaS

List the functions on the cluster

    faas-cli list

Remove the selected function

    faas-cli remove bcrypt-validate
