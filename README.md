# 📦 Real-Time Event Processing Platform on Kubernetes

## 🧭 Project Overview

This project is a real-time, high-throughput event processing pipeline built with Go, Kafka, Tarantool, TimescaleDB, and deployed on a lightweight Kubernetes cluster (k3s). The architecture is designed for scalability, observability, and extensibility.

It includes:

* Event generation and ingestion
* Stream processing
* Caching and hot data logic
* Time-series analytics
* Monitoring and GitOps deployment

---

## 🔄 Full Architecture

```
[event-generator] → [ingestion-service] → [Kafka] → [event-processor] → [Tarantool + TimescaleDB] → [Grafana/Prometheus]
```

---

## 🛣️ Development & Deployment Roadmap

### **🧱 Stage 1: Local Development and Containerization**

* [ ] Write `event-generator` (Go) — simulate JSON events and send via HTTP/gRPC
* [ ] Write `ingestion-service` (Go) — validate and publish to Kafka
* [ ] Write `event-processor` (Go) — consume from Kafka, process, store in DBs
* [ ] Create `Dockerfile` for each service
* [ ] Run full pipeline locally with Docker Compose

### **☁️ Stage 2: K3s Setup and Helm Deployment**

* [ ] Install k3s on a VM or local machine
* [ ] Connect kubectl to the cluster
* [ ] Install Helm
* [ ] Deploy Kafka via `bitnami/kafka`
* [ ] Deploy TimescaleDB via `bitnami/postgresql`
* [ ] Deploy Tarantool via custom chart or StatefulSet
* [ ] Deploy Prometheus and Grafana
* [ ] Configure Ingress and access Grafana

### **⚙️ Stage 3: Deploying Microservices in Kubernetes**

* [ ] Create Helm charts for `ingestion-service`, `event-processor`, `event-generator`
* [ ] Add `values.yaml`, ConfigMaps, Services
* [ ] Use PVCs for Kafka, Timescale, Tarantool
* [ ] Setup HPA for `event-processor`
* [ ] Test end-to-end data flow inside the cluster

### **📊 Stage 4: Observability & Visualization**

* [ ] Connect Grafana to TimescaleDB and Tarantool
* [ ] Build dashboards: events/sec, processing latency, aggregations
* [ ] Setup Prometheus metrics scraping
* [ ] Configure alert rules

### **🔁 Stage 5: CI/CD with GitHub Actions**

* [ ] Use DockerHub or GHCR as image registry
* [ ] Write GitHub Actions pipeline:

  * [ ] Build and test Go services
  * [ ] Build and push Docker images
  * [ ] Package and deploy Helm charts

### **🔄 Stage 6: GitOps with ArgoCD or Flux**

* [ ] Install ArgoCD or Flux in cluster
* [ ] Create `infrastructure-config` repo
* [ ] Store all HelmRelease and values.yaml here
* [ ] Connect repo to ArgoCD/Flux and auto-sync

### **🎲 Stage 7: Advanced Features and Experimentation**

* [ ] Add Istio or Linkerd for service mesh features
* [ ] Setup Kiali and Jaeger for observability
* [ ] Add real-time Web UI with React/WebSocket
* [ ] Create anomaly detection pipeline in event-processor

---

## 📁 Suggested Repo Structure

```
project-root/
├── charts/                 # Helm charts for services
├── cmd/                    # Go code for generator, ingestion, processor
├── deployments/            # Docker Compose for local
├── k8s-infra/              # GitOps manifests (Flux/ArgoCD)
├── .github/workflows/      # CI/CD YAMLs
└── README.md               # This file
```

---

## 💬 Next Steps

* [ ] Bootstrap Helm charts for ingestion-service
* [ ] Push base services into registry
* [ ] Initialize GitOps repo with ArgoCD structure

