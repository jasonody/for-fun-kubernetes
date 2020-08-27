# Secrets

[Source](https://medium.com/better-programming/how-to-use-kubernetes-secrets-for-storing-sensitive-config-data-f3c5e7d11c15)

## Using data section

Create secret and pod: `kubectl apply -f using-data-section.yaml`

Confirm secret was created: `kubectl get secret using-data-section -o yaml`

Access secret: `kubectl exec using-data-section-pod env | grep API_KEY`

Delete secret and pod: `kubectl delete -f using-data-section.yaml`

## Using stringData section

Create secrets and pod: `kubectl apply -f using-string-data-section.yaml`

Confirm secret was created: `kubectl get secret using-string-data-section -o yaml`

Access secrets: `kubectl exec using-string-data-section-pod env | grep foo` and `kubectl exec using-string-data-section-pod env | grep mac`

Delete secrets and pod: `kubectl delete -f using-string-data-section.yaml`

## Using `kubectl --from-literal`

Create secret: `kubectl create secret generic redis-credentials --from-literal=user=poweruser --from-literal=password='f0ob@r'`

Confirm secret was created: `kubectl get secrets redis-credentials -o yaml`

Delete secret: `kubectl delete secrets redis-credentials`

## Mounting secrets as volumes

Create secrets and pod: `kubectl apply -f mounting-secrets-as-volumes.yaml`

Confirm secrets were mounted as files: `kubectl exec mounting-secrets-as-volumes-pod ls /api-secrets`

Access secrets: `kubectl exec mounting-secrets-as-volumes-pod cat /api-secrets/apikey` and `kubectl exec mounting-secrets-as-volumes-pod cat /api-secrets/apisecret`

Delete secrets and pod: `kubectl delete -f mounting-secrets-as-volumes.yaml`