# Rifqah: Project Instructions & Conventions

This document outlines the architectural mandates and engineering standards for the Rifqah platform.

## 🏗️ Core Architecture
- **Privacy-First Design:** All user interactions must prioritize anonymity. Ephemeral aliases are mandatory for all meeting participants.
- **SFU Topology:** Centralized Selective Forwarding Unit (SFU) using LiveKit. The server acts as a packet reflector only; no audio decoding occurs server-side.
- **Edge Compute:** Heavy media processing (encryption, slicing, compression) is performed on client devices (Flutter/C++ FFI via Oboe).

## 🛠️ Backend Standards (Go)
- **Database Access:** Use `jmoiron/sqlx` for database interactions.
    - Prefer `NamedExec` for complex inserts to keep code clean and map structs directly.
    - Use `Get` and `Select` for mapping rows to structs.
- **Dependency Injection:** Services (Auth, Room) must receive the database connection (`*sqlx.DB`) via their constructor functions (e.g., `NewAuthService`).
- **Models:** Shared data structures must live in `internal/models`.
- **Identity:** All persistent entities use UUIDs. Meeting short IDs follow a 9-character alphanumeric format (e.g., `ABCD-1234`).
- **WebSocket Identification:** Clients must provide a `userID` query parameter when connecting to the WebSocket (e.g., `/ws/{roomID}?userID={uuid}`). This is used for session management and "Ghost Ball" protection.
- **Media Signaling:** Use `MediaService` to generate LiveKit access tokens. Clients must request a token before connecting to the SFU.
- **Turn Limits:** The default speaking limit is set to 3 minutes. The backend is authoritative and will automatically pass the ball upon expiration.
- **State Expiration:** All Redis keys managing temporary state (e.g., Ball queue, active speaker) must include a TTL (default 1 hour) to prevent state accumulation.

## 🔒 Privacy & Security
- **Aliases:** Mandatory ephemeral masks (Adjective + Noun) are generated for every participant joining a meeting. This is currently non-optional to reinforce the privacy mandate.
- **Password Security:** Use `bcrypt` for hashing. Raw passwords must never be stored or logged.

## 🎨 Design System (Serene Sanctuary)
- **Visual Identity:** All UI must adhere to the "Warm Organicism" aesthetic defined in `docs/DESIGN.md`.
- **Colors:**
    - Background: `#FDFBF7` (Cream)
    - Primary: `#1E3A3A` (Deep Sage Pine)
    - Secondary: `#D9AB8F` (Terracotta Sand)
    - Tertiary: `#F4EBD0` (Oatmeal)
- **Typography:** Use **Plus Jakarta Sans** with generous line-heights (1.6 for body).
- **Shapes:** Use hyper-organic, deeply rounded corners (minimum 24px radius). Pill shapes are preferred for buttons and inputs.
- **Elevation:** Use tonal layering (Oatmeal containers) instead of harsh shadows.

## 📜 Git Commit Conventions
We follow the **Conventional Commits** specification. Every commit must follow this format:
`<type>[optional scope]: <description>`

### Types:
- `feat`: A new feature (e.g., `feat(auth): add login endpoint`)
- `fix`: A bug fix (e.g., `fix(ws): resolve race condition in hub`)
- `docs`: Documentation changes only
- `style`: Changes that do not affect the meaning of the code (formatting, missing semi-colons, etc.)
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `build`: Changes that affect the build system or external dependencies
- `ci`: Changes to CI configuration files and scripts
- `chore`: Maintenance tasks (e.g., updating .gitignore)
- `revert`: Reverts a previous commit

### Guidelines:
- Use the imperative, present tense: "change" not "changed" nor "changes".
- Do not capitalize the first letter.
- No dot (.) at the end.

## 🧪 Testing Conventions
- **Co-location:** Tests must be co-located with the code they test (e.g., `internal/auth/auth_test.go`).
- **Frameworks:** Use standard `testing`, `github.com/stretchr/testify` for assertions, and `github.com/alicebob/miniredis/v2` for mocking Redis.
- **Setup:** Utilize `internal/testutils` to share common test environment initialization logic.
- **Execution:** All tests must pass running `go test ./...` from the `backend/` directory.
