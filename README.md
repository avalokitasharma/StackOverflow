  

A production-ready implementation of a **Stack Overflow-scale like Q&A platform** 

## Features:
- 5-minute vote lock + edit-reset rule
- Instant reputation reversal on vote flips
- Hard daily reputation cap of 200 (atomic enforcement)
- Sophisticated tag system (array + junction table + async counts)
- Event-driven reputation engine with exactly-once guarantees

### Non-Functional Aspects
- Supports millions of questions, 10M votes per day
- Strong consistency for votes & reputation
- Eventual consistency for search & notifications
- Horizontal scalability for all components
- 99.99% availability target

### Design Decisions
- Microservices (independent deploy + scale).
- API Gateway as single entry point ( AWS API GW for prod / Traefik locally).
- Q&A service gets the most traffic → Postgres is shared by question_id hash (or user_id for writes).
- All writes go through Kafka → async reputation/badge calc, notifications, search indexing.
- Reads 99% served from Redis cache (question + top answers + votes). Cache invalidation on write via Kafka.
- Search completely decoupled (Elasticsearch cluster).
- Multi-region active-active with global DB replication + conflict resolution on votes.


## Setup

### Prerequisites
- Docker & Docker Compose (v2.0+)
- Go 1.23+ 

### Clone & Prepare

```bash
git clone https://github.com/avalokitasharma/StackOverflow.git
cd StackOverflow
