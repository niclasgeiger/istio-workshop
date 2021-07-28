# Istio Introduction

## Prerequisites 

### Install minikube
```bash
brew install minikube
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

### Install istioctl
```bash
brew install istioctl
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
