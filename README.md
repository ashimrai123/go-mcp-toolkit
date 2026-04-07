# Go MCP Server (Custom Implementation)

A high-performance implementation of the **Model Context Protocol (MCP)** built entirely in Go. This server serves as a foundational prototype for integrating domain-specific engineering tools into LLM-powered workflows.

## Features
- **SSE Transport:** Full support for Server-Sent Events (SSE) for real-time communication with LLM clients (Claude Desktop, etc.).
- **JSON-RPC 2.0:** Strict adherence to the MCP message protocol for tool discovery and execution.
- **Dynamic Tool Dispatch:** A clean internal API for registering Go functions as LLM-callable tools.
- **Zero Dependencies:** Built primarily using the Go standard library to ensure a lightweight footprint.

## Architecture
The server follows a modular "Skill" pattern. Each skill (e.g., Log Analysis, Trace Querying) is registered with the server and exposed via the `list_tools` capability. 



## Why this exists
This project was developed to explore the technical surface of MCP before the official `jaegermcp` extension was finalized. It demonstrates the feasibility of using Go-based agents to perform complex observability tasks.
