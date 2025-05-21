# Project Setup and Usage

## Prerequisites

- **Docker** & **Docker Compose**
- **Go** (>= 1.18)
- **PostgreSQL Client** (optional, for manual DB interaction)

---

## Setup

### 1. Environment Variables

Create a `.env` file in the root of your project with the following content:

```env
DB_HOST=0.0.0.0
DB_PORT=5432
DB_USER=root
DB_PASSWORD=root
DB_NAME=drx
APP_PORT=8080
```

### 2. Usage with Makefile
The project includes a Makefile for convenience. Use the following commands to interact with the app:

### 🚀 Start Docker Services
```bash 
make up`
```

Runs ```docker compose up -d``` to start services in detached mode.

▶️ Run in Dev Mode
```bash
make run
```
Runs the app using:
```go
go run main.go
```

🛠 Build and Execute Binary

```bash
make build
```
Compiles and runs the Go binary:

```bash
go build main.go && chmod +x main && ./main
```
* Stop Docker Services

```bash
make down
```
Shuts down running containers with:
```bash
docker-compose down
```
📜 View Logs

```bash
make logs
```
Tails Docker logs using:
```bash
docker-compose logs -f

```
Notes

* Ensure .env is correctly configured before running the application.

* If you want to persist data between container runs, consider adding Docker volumes.

⚙️ You can use a migration tool (e.g., golang-migrate, goose) to manage DB schema updates.

