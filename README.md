# End-to-End DevOps Project using Golang, Kubernetes & GitOps

A production-style DevOps project demonstrating containerization, Kubernetes orchestration, CI/CD automation, and GitOps deployment workflows using Golang, Docker, Kubernetes, GitHub Actions, Helm, and Argo CD.

---

# Project Overview

This project showcases an end-to-end DevOps implementation for a Golang web application.

The application is containerized using Docker, deployed on Kubernetes, and automated through a CI/CD pipeline using GitHub Actions.

The project also follows GitOps practices using Argo CD and supports scalable Kubernetes deployments with Helm charts.

---

# Technologies Used

- Golang
- Docker
- Kubernetes
- Minikube
- GitHub Actions
- Helm
- Argo CD
- NGINX Ingress Controller
- Docker Hub

---

# Key Features

- Golang Web Application
- Static Content Hosting
- Multi-stage Docker Build
- Optimized Lightweight Docker Image
- Kubernetes Deployment & Service Configuration
- Health Monitoring using Liveness Probes
- CI/CD Automation using GitHub Actions
- GitOps Workflow using Argo CD
- Helm Chart Support for Multiple Environments
- Ingress Controller Configuration
- HTTPS/TLS Ready Architecture

---

# Project Structure

```text
devops-project/
│
├── app/
│   ├── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── static/
│       └── index.html
│
├── k8s/
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
│
├── helm/
│
├── .github/
│   └── workflows/
│       └── ci-cd.yaml
│
└── README.md
```

# Commands Used in This Project

# Golang Commands

## Initialize Go Module

```bash
go mod init devops-app
```

## Download Dependencies

```bash
go mod tidy
```

## Run Application

```bash
go run main.go
```

## Build Application

```bash
go build -o devops-app
```

---

# Docker Commands

## Build Docker Image

```bash
docker build -t devaasirvatham/devops-app:v1 .
```

## Check Docker Images

```bash
docker images
```

## Run Docker Container

```bash
docker run -d -p 8080:8080 devaasirvatham/devops-app:v1
```

## Docker Hub Login

```bash
docker login
```

## Push Docker Image

```bash
docker push devaasirvatham/devops-app:v1
```
