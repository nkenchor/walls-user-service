# 🧱 Walls User Service

> A dangerously talented, high-performance **User Service** powering Walls — a boundaryless connection platform. This microservice is designed with obsession over clean architecture, fault tolerance, and future-proof identity logic, built for scale across distributed systems.

---

## 🎯 Purpose

The **Walls User Service** acts as the identity and profile backbone of the Walls ecosystem. It handles user registration, authentication, verification, profile updates, organization membership, and access policies — built to be resilient, testable, and extensible.

---

## 🏗️ Architecture

Designed using **Clean Architecture**, **DDD**, and **Hexagonal Principles**. Every layer knows its boundary, and every dependency points inward.

```text
┌──────────────────────────┐
│     Delivery Layer       │  ← HTTP, Controllers, DTOs
├──────────────────────────┤
│     Application Layer    │  ← UseCases, Services
├──────────────────────────┤
│       Domain Layer       │  ← Entities, Aggregates, Repositories (Interfaces)
├──────────────────────────┤
│    Infrastructure Layer  │  ← MongoDB, Redis, Security, Logging, Email
└──────────────────────────┘
```

> Domain logic is king — infrastructure is replaceable.

---

## ⚙️ Tech Stack

| Component      | Tech                                |
|----------------|-------------------------------------|
| Language       | Go (Golang)                         |
| Framework      | Custom + net/http                   |
| DB             | MongoDB                             |
| Cache/Session  | Redis                               |
| Email Service  | SMTP / Pluggable Interface          |
| Auth           | JWT + Secure Token Flow             |
| Validation     | Fluent-style Custom Validation Layer|
| Logging        | Structured JSON Logger              |
| Deployment     | Docker / Bare Metal                 |
| Testing        | Go Test + Mocks                     |

---

## 📦 Features

- 🔐 User Registration (with email verification)
- ✅ Login/Logout + Refresh Token Rotation
- 🧾 Profile Management
- 🧠 Organisation & Membership Roles
- 🗝️ Password Reset & Email Confirmation
- 🧰 Rate-limiting, brute-force protection
- 🔄 Token blacklist & session integrity
- 🧪 Comprehensive unit and integration tests
- 🧩 Decoupled DTO and internal domain logic
- 🚦 Centralized error handling using custom `ApplicationError`

---

## 🔐 Security Philosophy

- ✅ Passwords are hashed using bcrypt with cost 12+
- ✅ Tokens are rotated securely, tied to fingerprint/device
- ✅ All write operations are idempotent by design
- ✅ Zero trust assumed — even internally
- ✅ Configurable CORS, rate-limits, and signature verification

---

## 📚 Getting Started

```bash
git clone https://github.com/your-username/walls-user-service.git
cd walls-user-service
cp .env.example .env
go run main.go
```

---

## 🧪 Testing

```bash
go test ./...
```

Test coverage includes:

- Domain use cases
- Validation rules
- Controller inputs/outputs
- Redis session management
- MongoDB repositories (mocked)

---

## 🧱 Core Domain: `User`

```go
type User struct {
  ID           string
  Email        string
  PasswordHash string
  Verified     bool
  CreatedAt    time.Time
  Profile      UserProfile
}
```

> Immutable root aggregate with value-object-style `UserProfile`

---

## 🧰 DTOs

Strict separation between domain and delivery. Examples:

```go
type RegisterUserDto struct {
  Email    string `json:"email"`
  Password string `json:"password"`
}

type LoginResponseDto struct {
  AccessToken  string `json:"access_token"`
  RefreshToken string `json:"refresh_token"`
}
```

---

## 🧠 Highlights

- 🚀 Inspired by production-grade identity systems like Auth0, Okta
- 🧪 Resilient against malformed input and malicious clients
- 📦 Designed for service mesh & API gateway integration
- 🧬 Built with the future in mind — extend with OAuth2, MFA, etc.
- 🧠 Validation rules built from reusable `ValidationRule` structs

---

## 📈 Roadmap

- [x] Redis-backed token sessions
- [x] Email verification flow
- [x] Organisation invite flow
- [ ] WebAuthn/FIDO2 support
- [ ] Admin impersonation
- [ ] Activity audit trails
- [ ] OTP/MFA integrations
- [ ] Metrics + observability layer

---

## 🧠 Who Built This?

Crafted by **Nkenchor Osemeke** — a hands-on Software & Technical Manager obsessed with design clarity, performance, and real-world scalability.

📧 nkenchor@osemeke.com  
🌍 [github.com/nkenchor](https://github.com/nkenchor)

---

## 📝 License

MIT License — see `LICENSE` file.