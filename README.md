# 🚀 Enterprise Event-Driven Task Suite

![Architecture](https://img.shields.io/badge/Architecture-Event--Driven-orange)
![RabbitMQ](https://img.shields.io/badge/Broker-RabbitMQ_3--Management-ff6600?logo=rabbitmq)
![Python](https://img.shields.io/badge/Publisher-Python_3.9-3776ab?logo=python)
![Go](https://img.shields.io/badge/Worker-Golang_1.21-00add8?logo=go)
![Vue](https://img.shields.io/badge/UI-Vue.js_2-4fc08d?logo=vuedotjs)

A professional-grade **Distributed Task Processing System**. This project demonstrates how to decouple high-traffic APIs from heavy background processing using **RabbitMQ** as a message broker, ensuring high availability and horizontal scalability.

---

## 🏛️ Architecture Overview

The system implements a reactive **Publisher/Subscriber** pattern:

1.  **Dashboard (Vue.js + Vuetify):** User interface to trigger and monitor tasks in real-time.
2.  **Publisher (Python Flask):** A REST API that validates incoming data and pushes tasks to the broker.
3.  **Message Broker (RabbitMQ):** Manages the `task_queue` with persistence, acting as a buffer during traffic spikes.
4.  **Worker (Golang):** A high-concurrency consumer that processes tasks asynchronously (e.g., simulations, report generation).



---

## 🚀 Getting Started

### 1. Prerequisites
- Docker and Docker Compose installed.

### 2. Launch the Suite
Clone the repository and run the build command:
```bash
docker-compose up -d --build

```

### 3. Service Access Points
| Service | URL / Port | Description |
| :--- | :--- | :--- |
| **Main Dashboard** | [http://localhost](http://localhost) | Trigger tasks via Web UI (Port 80) |
| **Publisher API** | [http://localhost:5000](http://localhost:5000) | REST API Entry point |
| **RabbitMQ Admin** | [http://localhost:15672](http://localhost:15672) | Monitor Queues (guest/guest) |

---

## ✅ System Verification

### Automatic Test (Web UI)
1. Open [http://localhost](http://localhost) in your browser.
2. Enter a task description (e.g., "Batch Processing #101").
3. Click **"Send Task"**.
4. Monitor the Go Worker logs to see the execution:
   ```bash
   docker-compose logs -f worker

### Manual API Test (CLI)
```bash
curl -X POST http://localhost:5000/send \
     -H "Content-Type: application/json" \
     -d '{"message": "Manual CLI Task"}'
```

## 🛠️ Key Technical Features

- **CORS-Ready:** Integrated `flask-cors` to allow seamless communication between the Vue.js frontend and the Python API.
- **Robust Go Build:** Containerized module initialization to resolve dependency management across different environments.
- **Fault Tolerance:** Implemented connection retry logic in both Python and Go to handle RabbitMQ's startup latency.
- **Horizontal Scaling:** Easily scale worker instances to handle larger queues:
  `docker-compose up -d --scale worker=5`
