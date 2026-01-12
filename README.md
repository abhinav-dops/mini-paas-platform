# Mini PaaS – Internal Developer Platform

## Project Overview

**Mini PaaS** is an internal developer platform built using **Go, Terraform, Docker, and AWS** that enables **self-service application deployment** on a single EC2 host.

The platform abstracts infrastructure provisioning and container orchestration behind a simple **REST API**, allowing developers to deploy applications without directly interacting with AWS, Terraform, or Docker.

This project demonstrates **Platform Engineering principles**, not just tool usage.

---

## Problem Statement

Developers often spend significant time:

* Provisioning infrastructure
* Setting up Docker environments
* Managing deployments manually

This leads to:

* Inconsistent environments
* Deployment errors
* Reduced developer productivity

**Goal:** Build a minimal internal PaaS that provides **one-command application deployment** with infrastructure automation and observability.

---

## Solution

Mini PaaS provides:

* Automated infrastructure provisioning using Terraform
* Remote Docker-based application deployment via SSH
* Repository-based application builds
* Per-application lifecycle management
* Health checks and basic resilience

All operations are exposed through a **Go-based control plane API**.

---

## Architecture

```
Developer (API Client)
        |
        v
Go Platform Service (Control Plane)
        |
        |-- Terraform (Provision EC2)
        |-- SSH Remote Execution
        |-- Git Clone
        |-- Docker Build & Run
        |
        v
AWS EC2 (Single Host)
        |
        v
Running Containers + Health Checks
```

---

## Technology Stack

| Layer          | Technology                 |
| -------------- | -------------------------- |
| Language       | Go                         |
| Infrastructure | Terraform                  |
| Containers     | Docker                     |
| Cloud          | AWS EC2 (Free Tier)        |
| API            | net/http                   |
| Provisioning   | SSH-based remote execution |

---

## Repository Structure

```
platform/
├── cmd/
│   └── platform/        # Main entrypoint
├── internal/
│   ├── api/             # HTTP handlers
│   ├── docker/          # Remote Docker operations
│   ├── terraform/       # Terraform executor
│   └── config/          # Environment configuration
├── terraform/           # Terraform modules
├── scripts/
├── examples/            # Sample deployable apps
├── .github/
├── go.mod
├── go.sum
└── README.md
```

---

## Application Contract

To deploy an application, the repository must:

* Contain a valid `Dockerfile`
* Expose a network port inside the container
* Be stateless

The platform requires the container port to be explicitly provided during deployment.

---

## 🔌 API Endpoints

### Provision Infrastructure

```http
POST /infra/provision
```

### Destroy Infrastructure

```http
POST /infra/destroy
```

### Deploy Application

```http
POST /apps/deploy
```

Request body:

```json
{
  "name": "sample-app",
  "repo": "https://github.com/user/repo",
  "port": 8080,
  "container_port": 5000
}
```

### List Applications

```http
GET /apps
```

### Application Status

```http
GET /apps/{name}/status
```

### Destroy Application

```http
DELETE /apps/{name}
```

---

## Observability & Health Checks

* Each deployed application is health-checked using HTTP probing
* Health checks include retry logic to handle startup delays
* Application state is tracked as:

  * `pending`
  * `running`
  * `unhealthy`
  * `failed`

---

## Security Considerations

* Infrastructure access via least-privilege IAM
* SSH key-based authentication
* No secrets committed to version control
* Environment-based configuration for sensitive paths

---

## Testing & Validation

* Manual and automated validation of:

  * Infrastructure provisioning
  * Remote Docker execution
  * Container lifecycle management
* Verified self-healing behavior using Docker restart policies

---

## Main 

> Designed and built an internal developer platform using Go, Terraform, Docker, and AWS that enables self-service application deployments with automated infrastructure provisioning, remote container orchestration, and health monitoring.

---

## Key Learnings

* Platform abstraction over infrastructure
* Remote execution and orchestration patterns
* Designing clean application contracts
* Balancing simplicity with extensibility

---

This project represents a **foundational internal PaaS implementation** aligned with real-world platform engineering practices.

