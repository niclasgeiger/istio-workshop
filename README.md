# Istio Introduction

## Prerequisites 

### Install minikube
```bash
brew install minikube
```
### Install istioctl
```bash
brew install istioctl
```
### Install skaffold
```bash
brew install skaffold
```

### Start cluster
```bash
minikube start --memory=16384 --cpus=4 --kubernetes-version=v1.20.2
```

### Update kubecontext to the minikube cluster
```bash
minikube update-context
```

### Start Loadbalancer for minikube
```bash
minikube tunnel
```

### Point Docker Daemon to Minikube
```bash
eval $(minikube -p minikube docker-env)
```

## Install Istio
### Install istio with the demo profile
```bash
istioctl install --set profile=demo
```

## Monitoring
### Prometheus (Metric Dashboards)
```bash
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.10/samples/addons/prometheus.yaml
istioctl dashboard prometheus
```
### Grafana (Metric Dashboards)
```bash
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.10/samples/addons/grafana.yaml
istioctl dashboard grafana
```
### Kiali (Network)
Might have to run this twice because of CRDs.
```bash
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.10/samples/addons/kiali.yaml
istioctl dashboard kiali
```
The Login Credentials are: admin/admin
### Jaeger (Tracing)
```bash
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.10/samples/addons/jaeger.yaml
istioctl dashboard jaeger
```
### Zipkin (Tracing)
```bash
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.10/samples/addons/extras/zipkin.yaml
istioctl dashboard zipkin
```

### Deploy Applications
```bash
make run0
# to stop again
make stop0
```

### Enable istio injection for default namespace
```bash
kubectl label namespace default istio-injection=enabled 
```

### Run first example
The first example is just a normal setup with 1 server and 1 client connecting via istio and exposed via its egress gateway.
```bash
make run0
# to stop
make stop0
```

### Run second example
The second example adds traffic shifting, shifting 50/50 the traffic from the client to two different versions of the server (see clock difference).
```bash
make run1
# to stop
make stop1
```

### Run third example
The third example adds support for JWT token verification. In order to get this running you have to supply a valid JWT token according to the configuration of the issuer.
This example uses a privately owned cognito pool for this right now.
The Header has to include `Authorization=Bearer <JWT-Token>`.
```bash
make run2
# to stop
make stop2
```

### Run fourth example
The fourth example shows how to artificially inject delay into the client networking.
This artificial delay can be enabled via the Header Value `Fault=true`
```bash
make run3
# to stop
make stop3
```
