---
title: 'Docs'
---

# Microservices Architecture — Docker Compose Setup!

Welcome to the documentation for the **microservices architecture** development environment! This repository includes a `docker-compose.yml` file designed to spin up all necessary services for local development.

---

## 🚀 Overview

This Docker Compose setup orchestrates multiple services that together form a microservices-based application ecosystem. It is **only intended for development purposes** and includes:

- **NATS & NATS UI**: Messaging system with JetStream enabled for event streaming and an admin UI for monitoring.
- **MongoDB**: Document-oriented NoSQL database for data storage.
- **MinIO**: High-performance, S3-compatible object storage.
- **Various Application Services**:
    - **Gateway**: API gateway routing requests.
    - **User Service**: Manages user data and authentication.
    - **Angebot Service**: Business-specific service for offers management.
    - **Media Service**: Handles media file uploads and processing.
- **Frontend & NGINX**: Web frontend served through NGINX with HTTPS support.

---

## 📦 Services Included

| Service         | Description                                  | Ports           |
|-----------------|----------------------------------------------|-----------------|
| **nats**        | Messaging server with JetStream enabled      | 4222, 8222      |
| **nats-ui**     | NATS monitoring UI                           | 31311           |
| **mongo**       | MongoDB database                             | 27017           |
| **minio**       | Object storage (S3 compatible)               | 9000 (API), 9001 (Console) |
| **gateway**     | API Gateway                                 | 8081            |
| **user-service**| User management                             | 8082            |
| **angebot-service** | Offers management                         | 8084            |
| **media-service**| Media handling                              | 8083            |
| **frontend**    | Web frontend                                | 8080            |
| **nginx**       | Reverse proxy & HTTPS termination           | 80, 443         |

---

## ⚙️ How to Run

1. Make sure you have [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/) installed.

2. Clone the repository:

   ```bash
   git clone https://github.com/Konzepte-moderner-Softwareentwicklung/Backend.git
   cd Backend
