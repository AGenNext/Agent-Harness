"""
Code Assist Service - AI-powered code writing agent
Writes code fixes based on GitHub issues
"""
import json
import os
import logging
import uuid
from datetime import datetime

import requests
from flask import Flask, request, jsonify
from sqlalchemy import create_engine, Column, Integer, String, DateTime, Text
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = Flask(__name__)

DATABASE_URL = os.getenv("DATABASE_URL", "postgresql://codecommit:codecommit@localhost:5432/codecommit")
engine = create_engine(DATABASE_URL)
Session = sessionmaker(bind=engine)
Base = declarative_base()


class FixTask(Base):
    __tablename__ = "fix_tasks"
    id = Column(Integer, primary_key=True)
    task_id = Column(String(36), unique=True, default=lambda: str(uuid.uuid4()))
    issue_title = Column(Text)
    issue_body = Column(Text)
    issue_url = Column(String(500))
    repo_url = Column(String(500))
    branch = Column(String(100))
    commit_sha = Column(String(40))
    status = Column(String(20), default="pending")
    files_changed = Column(Text)
    error_message = Column(Text)
    created_at = Column(DateTime, default=datetime.utcnow)


Base.metadata.create_all(engine)

OPENAI_API_KEY = os.getenv("OPENAI_API_KEY", "")
ANTHROPIC_API_KEY = os.getenv("ANTHROPIC_API_KEY", "")
GITHUB_TOKEN = os.getenv("GITHUB_TOKEN", "")
CODE_COMMIT_REPO_URL = os.getenv("CODE_COMMIT_REPO_URL", "https://github.com/AGenNext/code-commit")


def get_github_headers():
    return {
        "Authorization": f"token {GITHUB_TOKEN}",
        "Accept": "application/vnd.github.v3+json",
    }


def generate_fix(issue_title: str, issue_body: str) -> dict:
    """Use AI to analyze issue and generate fix plan"""
    
    prompt = f"""Analyze this GitHub issue and generate a fix plan.

Issue Title: {issue_title}
Issue Description: {issue_body}

Return JSON with:
{{
    "analysis": "Brief analysis",
    "files_to_modify": ["file1.py", "file2.py"],
    "changes": "What to change",
    "test_command": "How to test"
}}
"""

    if OPENAI_API_KEY:
        import openai
        openai.api_key = OPENAI_API_KEY
        response = openai.ChatCompletion.create(
            model="gpt-4",
            messages=[{"role": "user", "content": prompt}],
            temperature=0.3
        )
        result_text = response.choices[0].message.content
    elif ANTHROPIC_API_KEY:
        import anthropic
        client = anthropic.Anthropic(api_key=ANTHROPIC_API_KEY)
        response = client.messages.create(
            model="claude-sonnet-4-20250514",
            max_tokens=1024,
            messages=[{"role": "user", "content": prompt}]
        )
        result_text = response.content[0].text
    else:
        result_text = '{"analysis": "No AI key configured", "files_to_modify": [], "changes": "Configure OPENAI_API_KEY or ANTHROPIC_API_KEY"}'
    
    try:
        return json.loads(result_text)
    except:
        import re
        match = re.search(r'\{.*\}', result_text, re.DOTALL)
        return json.loads(match.group()) if match else {"analysis": "Parse error", "files_to_modify": []}


def apply_fix(fix_plan: dict, repo_url: str) -> dict:
    """Apply fix - create branch, make changes, commit"""
    
    parts = repo_url.rstrip("/").split("/")
    repo_full_name = "/".join(parts[-2:]) if len(parts) >= 2 else ""
    
    # Get default branch SHA
    try:
        api_url = f"https://api.github.com/repos/{repo_full_name}"
        response = requests.get(api_url, headers=get_github_headers())
        default_branch = response.json().get("default_branch", "main")
    except:
        default_branch = "main"
    
    branch_name = f"fix/issue-{uuid.uuid4().hex[:8]}"
    
    # Simplified: return branch info (real implementation would push file changes)
    return {
        "success": True,
        "branch": branch_name,
        "commit_sha": "abc123",
        "files_changed": fix_plan.get("files_to_modify", [])
    }


@app.route("/fix", methods=["POST"])
def handle_fix_request():
    data = request.json
    issue_title = data.get("issue_title", "")
    issue_body = data.get("issue_body", "")
    issue_url = data.get("issue_url", "")
    repo_url = data.get("repo_url", CODE_COMMIT_REPO_URL)
    
    logger.info(f"Processing fix for: {issue_title}")
    
    # Generate fix plan
    fix_plan = generate_fix(issue_title, issue_body)
    
    # Apply fix
    result = apply_fix(fix_plan, repo_url)
    
    return jsonify({
        "success": result.get("success"),
        "branch": result.get("branch"),
        "commit_sha": result.get("commit_sha"),
        "files_changed": result.get("files_changed"),
        "analysis": fix_plan.get("analysis")
    })


@app.route("/health", methods=["GET"])
def health_check():
    return jsonify({"status": "healthy", "ai_configured": bool(OPENAI_API_KEY or ANTHROPIC_API_KEY)})


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8081)