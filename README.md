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