# Secrets

[Source](https://medium.com/better-programming/how-to-use-kubernetes-secrets-for-storing-sensitive-config-data-f3c5e7d11c15)

## Using data section

Create secret: `kubectl apply -f using-data-section.yaml`

Confirm secret was created: `kubectl get secret using-data-section -o yaml`

Delete secret: `kubectl delete -f using-data-section.yaml`

## Using stringData section

Create secret: `kubectl apply -f using-string-data-section.yaml`

Confirm secret was created: `kubectl get secret using-string-data-section -o yaml`

Delete secret: `kubectl delete -f using-string-data-section.yaml`

## Using `kubectl --from-literal`

Create secret: `kubectl create secret generic redis-credentials --from-literal=user=poweruser --from-literal=password='f0ob@r'`

Confirm secret was created: `kubectl get secrets redis-credentials -o yaml`

Delete secret: `kubectl delete secrets redis-credentials`

