# Configuration

[Source](https://itnext.io/learn-how-to-configure-your-kubernetes-apps-using-the-configmap-object-d8f30f99abeb)

## Pod env vars

Create pod: `kubectl apply -f pod-env-var.yaml`

Verify pod's env vars: `kubectl exec pod-env-var -ti env | grep ENV_VAR`

Delete Pod `kubectl delete -f pod-env-var.yaml`

## ConfigMap envFrom (including storing json)

Create pod: ``

Verify configmap: `kubectl describe configmaps my-configmap`

Verify pod's env vars: `kubectl exec configmap-envfrom -ti env`

Delete pod: `kubectl delete -f configmap-envfrom.yaml`