# Agent Harness

Open-source AI agent platform with 4 specialized agents.

## Features

- 🤖 **4 AI Agents** - code-assist, code-review, code-tester, code-deploy
- 💬 **Notifications** - Slack, Discord, Mattermost, Teams, Email
- 🎨 **Workflows** - LangGraph orchestrator, LangFlow visual
- 🔐 **Security** - Policy, Guardrail, Vulnerability scanner
- 🔏 **Privacy** - SSO, PII removal, Audit logs
- 🔑 **Secrets** - Infisical integration
- 👂 **Listeners** - GitHub, Slack, Email, Jira, Linear
- 🖥️ **Microsoft** - Copilot/Teams, M365, SharePoint
- 📚 **Learning** - GoodReads recommendations
- 📋 **Project** - Linear, Jira integration

## Quick Start

```bash
# Clone
git clone https://github.com/AGenNext/agent-harness.git
cd agent-harness

# Configure
cp deploy/.env.example .env

# Docker
docker build -t agent-harness .
docker run -d -p 3000:3000 agent-harness

# Docker Compose
docker-compose up -d
```

## API

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/v1/agents` | GET | List all agents |
| `/api/v1/agents/:name/run` | POST | Run agent |
| `/api/v1/workflows` | GET | List workflows |
| `/api/v1/policies` | GET | List policies |

## Docker

```bash
# Build
docker build -t agent-harness .

# Push to GHCR
docker push ghcr.io/youruser/agent-harness:latest

# Push to Docker Hub
docker push youruser/agent-harness:latest
```

## Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `GITHUB_TOKEN` | Yes | GitHub PAT |
| `OPENAI_API_KEY` | No | OpenAI key |
| `ANTHROPIC_API_KEY` | No | Claude key |

## License

Apache 2.0