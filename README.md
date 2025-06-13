# 🧠 MCP Server — Modular Code Platform Server (v1)

A powerful backend service built in Go that provides real-time access, browsing, and future GitHub integration for seamless code project management and AI-assisted workflows.

---

## 🚀 Features (v1)

- ✅ Modular Go Project Structure
- 🌲 File Tree API: Get a real-time directory structure as JSON
- 🌐 HTTP Server with CORS and Routing (using `gorilla/mux`)
- 🧩 Middleware-ready setup
- 💡 Clean code architecture (`cmd`, `internal`, `pkg`)
- 🛠️ Ready for future GitHub integration (v2)

---

## 🛠️ Project Structure

mcp-server/
├── cmd/ # Main application entry (main.go)
│ └── main.go
├── internal/ # Internal application logic
│ ├── api/ # HTTP handlers
│ │ └── filetree.go
│ ├── middleware/ # Middleware components
│ └── config/ # Config & constants (future)
├── pkg/ # Reusable utilities
│ └── utils/ # File tree logic
├── go.mod # Go module definition
└── README.md # You're here!

yaml
Copy
Edit

---

## 📦 File Tree API (v1)

> Returns the recursive directory structure from root (`.`)

**URL**: `GET /filetree`

**Sample Response**:
```json
[
  {
    "name": "cmd",
    "path": "cmd",
    "isDir": true,
    "children": [
      {
        "name": "main.go",
        "path": "cmd/main.go",
        "isDir": false
      }
    ]
  }
]
```
📁 Future Roadmap (v2+)
Feature	Status
GitHub Repo Integration	🕐 Planned
AI Copilot-like context support	🕐 Planned
Token-based Authentication	🕐 Planned
OAuth GitHub Login	🕐 Planned
MCP Desktop Agent / Client App	🕐 Planned
Project Metadata & Caching Layer	🕐 Planned

⚙️ Getting Started
1️⃣ Clone the Repo
```
git clone https://github.com/manishkandari9/mcp-server.git
cd mcp-server
```
2️⃣ Install Dependencie
```
go mod tidy
```
3️⃣ Run Server
```
cd cmd
go run main.g
```
```
Server will start at: http://localhost:8080
```
## 🧱 Tech Stack
Language: Go (v1.20+)

Router: Gorilla Mux

CORS: github.com/rs/cors

## 🔮 Roadmap (v2 Preview)
🔐 GitHub OAuth Integration

📂 Browse Repositories & Files via GitHub API

📥 Clone and parse GitHub repositories

🤖 AI/Agent-ready endpoints for Copilot-style assistants

## 🤝 Contributing
Pull requests are welcome. For major changes, open an issue first to discuss what you would like to change.








