# 🌐 Web Page Analyzer (Go)

A concurrent, pluggable **web page analyzer** written in Go. It extracts and reports metadata from any public web page such as:

- ✅ HTML version
- ✅ Page title
- ✅ Heading structure (H1-H6)
- ✅ Number of internal/external/inaccessible links
- ✅ Login form detection

Built with Go routines, `gin`, `goquery`, `logrus`, and supports extensible analyzers.

---

## 🚀 Features

- 📄 **HTML Metadata Extraction**
- 🔍 **Heading Analysis** (H1 to H6 count)
- 🔐 **Login Form Detection**
- 🔗 **Internal/External Link Checks** (with concurrency and timeout handling)

---

## 📦 Tech Stack

| Tool        | Purpose                       |
|-------------|-------------------------------|
| Go          | Language                      |
| gin-gonic   | Web framework                 |
| goquery     | HTML parsing (like jQuery)    |
| logrus      | Structured logging            |
| httptest    | Testing support               |

---

## 🛠️ Requirements

    ✅ Docker must be installed

    ✅ Git (to clone the repo)

    ✅ Go (if running outside Docker)

💡 Make sure Docker is running on your system before proceeding.

## 🛠️ Installation


# 🐳 Running Web Page Analyzer with Docker

This project is Dockerized for easy deployment. You can run the analyzer using **Docker Compose**.

---

## 🚀 Quick Start with Docker Compose

Make sure Docker and Docker Compose are installed on your machine.

### Step 1: Clone the Repository

```bash
git clone https://github.com/sadeepaalw/webpage-analyzer.git
cd webpage-analyzer
```

### Step 2: Build and Run the App

```bash
docker-compose up --build

```

This command:

    Builds the Docker image using the provided Dockerfile

    Starts the web analyzer service

    Exposes it on http://localhost:8080

### Step 3: Use the Analyzer

Visit 

http://localhost:8080 in your browser to access the UI


### Stopping the Service

```bash
docker-compose down
```

## 🧪 Running Tests


Make sure go is installed on your machine.

### Run Unit Tests

```bash
go test ./... -v
```

### Run Tests with Code Coverage

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

Folder Structure

```bash
webpage-analyzer/
├── adapter/                 # HTTP client utilities
├── handler/                 # Request handler responsible for loading the UI and results or the errors
├── modals/                  # Data models and interfaces
├── routes/                  # Routes for the request handler
├── services/                # All analyzers (title, HTML version, links, etc.)
├── resources/               # Sample HTML files for testing
├── utils/                   # Utilities for logger and URL 
├── validators/              # Validators for validating URL
├── web/                     # Html content files 
├── Dockerfile               # Docker image definition
├── docker-compose.yml       # Docker Compose config
├── main.go                  # Entry point
└── README.md
```