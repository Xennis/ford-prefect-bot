# FordPerfectBot service


## Notes

### Call webhook for updates

```sh
export SERVICE_IP=$(kubectl get services ford-perfect-service -o=jsonpath='{.status.loadBalancer.ingress[0].ip}')
export ENDPOINTS_KEY=<token>
curl -v --insecure --request POST --header "content-type:application/json" --data '{"update_id": 10}' https://$SERVICE_IP/updatesHook\?key=$ENDPOINTS_KEY
```