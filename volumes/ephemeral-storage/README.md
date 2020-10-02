# Volumes: Epemeral Storage

Ephemeral storage is tightly coupled to a pods lifetime and are deleted if the pod is removed for any reason

[Source](https://itnext.io/learn-about-the-basics-of-kubernetes-persistence-part-1-b1fa2847768f)

## Create Docker image (if needed)

Build image: `docker build -t kvstore .`

Tag image: `docker tag kvstore:latest jasonody/kvstore:1.0.0`

Login to Docker: `docker login`

Push image to Docker: `docker push jasonody/kvstore:1.0.0`

## Using emptyDir

Create deployment: `kubectl apply -f empty-dir.yaml`

Confirm deployment was created: `kubectl get deployments empty-dir`

Confim pod is ready: `kubectl get pod -l app=empty-dir`

Get the service's port: `kubectl get services empty-dir-service | egrep -o ':(.*?)\/' | grep '[0-9]'`

Save to the store: `curl -X POST http://localhost:<service port>/save -d 'hello=world'`

Read from the value: `curl 'http://localhost:<service port>/read?key=hello'`

Save to the store: `curl -X POST http://localhost:<service port>/save -d 'oh=yeah'`

Get the pod name: `kubectl get pod -l app=empty-dir`

Peek inside /data: `kubectl exec <pod name> -- ls data/`

Confirm value for "hello" key: `kubectl exec <pod name> -- cat data/hello`

Delete deployment: `kubectl delete -f empty-dir.yaml`