# Rifqah: Project Instructions & Conventions

This document outlines the architectural mandates and engineering standards for the Rifqah platform.

## 🏗️ Core Architecture
- **Privacy-First Design:** All user interactions must prioritize anonymity. Ephemeral aliases are mandatory for all meeting participants.
- **SFU Topology:** Centralized Selective Forwarding Unit (SFU) using LiveKit. The server acts as a packet reflector only; no audio decoding occurs server-side.
- **Edge Compute:** Heavy media processing (encryption, slicing, compression) is performed on client devices (Flutter/C++ FFI).

## 🛠️ Backend Standards (Go)
- **Database Access:** Use `jmoiron/sqlx` for database interactions.
    - Prefer `NamedExec` for complex inserts to keep code clean and map structs directly.
    - Use `Get` and `Select` for mapping rows to structs.
- **Dependency Injection:** Services (Auth, Room) must receive the database connection (`*sqlx.DB`) via their constructor functions (e.g., `NewAuthService`).
- **Models:** Shared data structures must live in `internal/models`.
- **Identity:** All persistent entities use UUIDs. Meeting short IDs follow a 9-character alphanumeric format (e.g., `ABCD-1234`).

## 🔒 Privacy & Security
- **Aliases:** Mandatory ephemeral masks (Adjective + Noun) are generated for every participant joining a meeting. This is currently non-optional to reinforce the privacy mandate.
- **Password Security:** Use `bcrypt` for hashing. Raw passwords must never be stored or logged.

## 📅 Roadmap Summary
Refer to `Plan.md` for the detailed 6-phase implementation lifecycle.
