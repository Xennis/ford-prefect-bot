# FordPerfectBot service

## Development

Run the service by `go run *.go`.

## Deployment @ GCP Kubernetes Engine

First create the Kubernets secrets by running `make kubernetesCreateSecrets`. Than create the docker image, push it to the Docker registry and apply the Kubernetes resources by executing `make kubernetesApply`.

## Notes

### Call webhook for updates

```sh
export SERVICE_IP=$(kubectl get services ford-perfect-service -o=jsonpath='{.status.loadBalancer.ingress[0].ip}')
export ENDPOINTS_KEY=<token>
curl -v --insecure --request POST --header "content-type:application/json" --data '{"update_id": 10}' https://$SERVICE_IP/updatesHook\?key=$ENDPOINTS_KEY
```