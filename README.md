This repository contains a Student CRUD REST API implemented in Golang using the Gin framework. The project serves as a learning exercise to explore best practices in building RESTful APIs, containerization, CI/CD pipelines, deployment, and observability. It is aimed at

Features

Functional Requirements

Create a new student: Add a student record to the database.

Retrieve all students: Fetch all student records.

Retrieve a specific student: Fetch details of a student by ID.

Update student information: Modify details of an existing student.

Delete a student: Remove a student record from the database.

Healthcheck endpoint: Monitor API health.

Non-Functional Requirements

API versioning: Supports versioning (e.g., api/v1/<resource>).

Meaningful logging: Emits structured logs with appropriate log levels.

Environment variables: Configurations are externalized and injected.

Unit testing: Includes tests for all endpoints.

Database migrations: Automates schema creation and updates.

Postman collection: Pre-configured requests for API testing.

Technologies Used

Core Technologies

Programming Language: Golang

Framework: Gin

Database: PostgreSQL

Containerization and CI/CD

Docker: Multi-stage builds and containerization

Docker Compose: Simplified local setup

GitHub Actions: CI/CD pipeline with build, test, and deployment stages

Deployment

Vagrant: Bare-metal deployment

Nginx: Load balancing

Kubernetes: Orchestration using Minikube

Helm: Deployment templating

ArgoCD: GitOps-based continuous deployment

Observability

Prometheus: Metrics monitoring

Grafana: Visualizations

Loki: Log aggregation

Promtail: Log shipping

Getting Started

Prerequisites

Install Docker

Install Docker Compose

Install GNU Make

Install Vagrant

Install kubectl

Install Minikube

Local Setup

Clone the repository:

git clone https://github.com/<your-username>/student-crud-api.git
cd student-crud-api

Run database migrations:

make db-migrate

Build and run the API locally:

make run

Test the API using the provided Postman collection.

Containerization

Build Docker Image

make docker-build

Run Docker Container

make docker-run

Inject Environment Variables

Environment variables can be injected at runtime using a .env file or passed directly during container execution.

One-Click Local Setup

Start the database and API:

make compose-up

Verify the API healthcheck:

curl http://localhost:8080/healthcheck

Stop all services:

make compose-down

CI/CD Pipeline

The CI/CD pipeline is implemented using GitHub Actions with the following stages:

Build API

Run Tests

Perform Code Linting

Build and Push Docker Image

Manual Trigger Support

Deployment

Bare-Metal Deployment (Vagrant)

Spin up the Vagrant box:

vagrant up

Deploy services using Docker Compose:

make deploy

Kubernetes Deployment

Start Minikube:

minikube start

Apply Kubernetes manifests:

kubectl apply -f k8s/

Verify API:

kubectl get pods -n student-api

Helm Charts

Deploy stack using Helm:

helm install student-api ./helm/student-api

Observability Stack

Deploy Prometheus, Loki, and Grafana:

helm install observability ./helm/observability

Access Grafana dashboard:

kubectl port-forward svc/grafana 3000:3000 -n observability

Add Prometheus and Loki as data sources in Grafana.

GitOps with ArgoCD

Deploy ArgoCD:

kubectl apply -f argocd/

Access ArgoCD UI:

kubectl port-forward svc/argocd-server 8080:443 -n argocd

Configure ArgoCD to sync Helm charts with GitHub.

Roadmap

Complete API functionality with unit and integration tests.

Implement advanced observability for end-to-end monitoring.

Scale deployment using Helm charts and ArgoCD.

Optimize container images for size and performance.

Improve CI/CD pipeline with advanced GitHub Actions.
