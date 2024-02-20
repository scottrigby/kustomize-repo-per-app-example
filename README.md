# Kustomize repo per app example

Simple non-gitops example of a [repo per app](https://fluxcd.io/flux/guides/repository-structure/#repo-per-app) model with [Kustomize bases and overlays](https://github.com/kubernetes-sigs/kustomize/tree/master/examples/helloWorld) per envionment.

## Repo structure

```sh
├── Dockerfile
├── README.md
├── deploy
│   ├── base
│   │   ├── deployment.yaml
│   │   └── kustomization.yaml
│   ├── local
│   │   ├── kustomization.yaml
│   │   ├── patch-env.yaml
│   │   ├── patch-replicas.yaml
│   │   └── patch-tag.yaml
│   ├── production
│   │   ├── kustomization.yaml
│   │   ├── patch-cpu-memory.yaml
│   │   ├── patch-env.yaml
│   │   ├── patch-replicas.yaml
│   │   └── patch-tag.yaml
│   └── staging
│       ├── kustomization.yaml
│       ├── patch-cpu-memory.yaml
│       ├── patch-env.yaml
│       ├── patch-replicas.yaml
│       └── patch-tag.yaml
├── kind-local-registry.sh
└── src
    ├── go.mod
    └── main.go

```

## Preview difference

See the difference between building for local vs staging vs production:

<details>
  <summary>local vs staging</summary>

```diff
$ diff -u \
  <(kustomize build deploy/local/) \
  <(kustomize build deploy/staging)

--- /dev/fd/11	2024-02-20 00:17:58
+++ /dev/fd/13	2024-02-20 00:17:58
@@ -3,7 +3,7 @@
 metadata:
   name: hello
 spec:
-  replicas: 1
+  replicas: 3
   selector:
     matchLabels:
       run: hello
@@ -15,12 +15,18 @@
       containers:
       - env:
         - name: HELLO
-          value: world
+          value: galaxy
         - name: WHATS
           value: this
         - name: FOO
           value: bar
-        image: localhost:5001/hello:v3
+        image: localhost:5001/hello:v2
+        limits:
+          cpu: 500m
+          memory: 128Mi
         name: hello
         ports:
         - containerPort: 8080
+        requests:
+          cpu: 250m
+          memory: 64Mi

```

</details>

<details>
  <summary>staging vs production</summary>

```diff
$ diff -u \
  <(kustomize build deploy/staging/) \
  <(kustomize build deploy/production)

--- /dev/fd/11	2024-02-20 00:20:20
+++ /dev/fd/13	2024-02-20 00:20:20
@@ -3,7 +3,7 @@
 metadata:
   name: hello
 spec:
-  replicas: 3
+  replicas: 10
   selector:
     matchLabels:
       run: hello
@@ -15,18 +15,18 @@
       containers:
       - env:
         - name: HELLO
-          value: galaxy
+          value: universe
         - name: WHATS
-          value: this
+          value: up doc
         - name: FOO
           value: bar
-        image: localhost:5001/hello:v2
+        image: localhost:5001/hello:v1
         limits:
-          cpu: 500m
-          memory: 128Mi
+          cpu: 1
+          memory: 256Mi
         name: hello
         ports:
         - containerPort: 8080
         requests:
-          cpu: 250m
-          memory: 64Mi
+          cpu: 500m
+          memory: 128Mi
```

</details>

## Local K8s demo

Create local cluster with docker registry:

```sh
./kind-local-registry.sh
```

_To use local registry, see <https://kind.sigs.k8s.io/docs/user/local-registry/>._

Build the image:

```sh
docker build . -f deploy/docker/Dockerfile -t hello
```

Create and push some demo tags to local registry:

```sh
docker build . -t hello
# Same image. Tag names here only illustrate kustomize
# patching different tags for local, staging and production.
for v in v1 v2 v3; do
  docker tag hello localhost:5001/hello:$v
  docker push localhost:5001/hello:$v
done
```

Deploy local environment:

```sh
kubectl apply -k deploy/local

# preview
$ kubectl port-forward deployments/hello 8080:8080

# seperate tab
$ curl localhost:8080
Hello world
```

Simulate staging environment locally:

```sh
kubectl create ns staging
kubectl -n staging apply -k deploy/staging

# preview
$ kubectl -n staging port-forward deployments/hello 8080:8080

# seperate tab
$ curl localhost:8080
Hello galaxy
```

## Cleanup

```sh
kind delete cluster
```
