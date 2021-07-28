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
### Access Prometheus (Metric Dashboards)
```bash
istioctl dashboard prometheus
```
### Access Grafana (Metric Dashboards)
```bash
istioctl dashboard grafana
```
### Access Kiali (Network)
```bash
istioctl dashboard kiali
```
The Login Credentials are: admin/admin
### Access Jaeger (Tracing)
```bash
istioctl dashboard jaeger
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
