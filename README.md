# 🚀 Distributed Event-Driven Task Suite

![Architecture](https://img.shields.io/badge/Architecture-Event--Driven-orange)
![RabbitMQ](https://img.shields.io/badge/Broker-RabbitMQ_3--Management-ff6600?logo=rabbitmq)
![Python](https://img.shields.io/badge/Publisher-Python_3.9-3776ab?logo=python)
![Go](https://img.shields.io/badge/Worker-Golang_1.21-00add8?logo=go)

A high-performance **Distributed Task Processing System** built on an **Event-Driven Architecture (EDA)**. This project demonstrates how to decouple high-traffic APIs from heavy background processing using a message broker.

---

## 🏛️ Architecture Overview

The system utilizes a **Publisher/Subscriber** pattern to ensure asynchronous task execution and system reliability:

1.  **Publisher (Python Flask):** A REST API that receives tasks and pushes them into a RabbitMQ queue.
2.  **Message Broker (RabbitMQ):** Manages the "task_queue," ensuring messages are stored safely until a worker is ready.
3.  **Worker (Golang):** A high-speed consumer that pulls tasks from the queue, processes them, and logs the results.
4.  **Frontend (Vue.js + Vuetify):** A dashboard to trigger tasks and monitor the real-time flow of events.

[Image of RabbitMQ Publisher Consumer architecture with Message Broker and Worker]

---

## 🛠️ Tech Stack

| Component | Technology | Role |
| :--- | :--- | :--- |
| **Orchestration** | Docker & Docker Compose | Container management |
| **Message Broker** | RabbitMQ | Reliable message delivery (AMQP) |
| **API / Publisher** | Python (Flask + Pika) | Fast request handling |
| **Worker / Consumer** | Golang (streadway/amqp) | Efficient parallel processing |
| **UI** | Vue.js + Vuetify | User interaction and monitoring |

---

## 🚀 Getting Started

### 1. Prerequisites
- Docker and Docker Compose installed.

### 2. Launch the Infrastructure
Clone the repository and run:
```bash
docker-compose up -d --build
