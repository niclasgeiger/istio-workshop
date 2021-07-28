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
istioctl dashboard prometheus
```
### Grafana (Metric Dashboards)
```bash
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.10/samples/addons/prometheus.yaml
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
### Jaeger (Tracing)
```bash
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.10/samples/addons/extras/zipkin.yaml
istioctl dashboard zipkin
```

### Deploy Applications
```bash
skaffold run
```

### Enable istio injection for default namespace
```bash
kubectl label namespace default istio-injection=enabled 
```

### Deploy Application again and check the sidecars

```bash
skaffold delete
skaffold run
```

### Apply istio gateway
```bash
kubectl apply -f istio/gateway.yaml 
```
