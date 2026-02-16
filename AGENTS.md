# Kosmetes - web interface for Taskwarrior

In this project we build a web interface for Taskwarrior - a CLI tool for task management.

Our goals:
- support for main Taskwarrior features (tasks, projects, tags, due dates, filters)
- minimalistic interface
- simple tech stack
- responsive design for desktop and mobile devices

## Tech Stack

Backend:
- Go standard library only (no external dependencies)
- Communication with Taskwarrior via CLI commands (`task` command with JSON export)

Frontend:
- Plain HTML with Go templates
- Lightweight CSS frameworks (pico.css)
- Server-driven UI with minimal client-side JavaScript (htmx)

## Architecture

- Backend executes Taskwarrior CLI commands and parses JSON output
- No authentication - designed for local/single-user usage
- RESTful HTTP server on port 8000
- Keep all web stuff inside `internal/web`
