# Kustomize repo per app example

Simple example of a [repo per app](https://fluxcd.io/flux/guides/repository-structure/#repo-per-app) model with [Kustomize bases and overlays](https://github.com/kubernetes-sigs/kustomize/tree/master/examples/helloWorld) per envionment.

_Note if you want a full GitOps example with Flux, Kustomize and Helm instead, see [fluxcd/flux2-kustomize-helm-example](https://github.com/fluxcd/flux2-kustomize-helm-example/tree/main)._

## Repo structure

```sh
├── deploy
│   ├── base
│   │   ├── deployment.yaml
│   │   └── kustomization.yaml
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
└── src
    └── main.go
```

See the difference between building for staging vs production:

```diff
$ diff -u \
  <(kustomize build deploy/staging/) \
  <(kustomize build deploy/production)

--- /dev/fd/11	2024-02-19 21:34:47
+++ /dev/fd/13	2024-02-19 21:34:47
@@ -3,7 +3,7 @@
 metadata:
   name: my-nginx
 spec:
-  replicas: 3
+  replicas: 10
   selector:
     matchLabels:
       run: my-nginx
@@ -15,18 +15,18 @@
       containers:
       - env:
         - name: HELLO
-          value: world
+          value: galaxy
         - name: WHATS
-          value: this
+          value: up doc
         - name: FOO
           value: bar
-        image: nginx:mainline
+        image: nginx:stable
         limits:
-          cpu: 500m
-          memory: 128Mi
+          cpu: 1
+          memory: 256Mi
         name: my-nginx
         ports:
         - containerPort: 80
         requests:
-          cpu: 250m
-          memory: 64Mi
+          cpu: 500m
+          memory: 128Mi
```
