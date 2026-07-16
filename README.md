# Chirpy

Chirpy is a REST API project built with Go and PostgreSQL. It is a fictional app with no real-world functionality, but it helped me learn about endpoints, auth, & webhooks.

## Features

- User registration and account updates
- Password hashing
- JWT access tokens and refresh tokens
- Chirp creation, retrieval, filtering, sorting, and deletion
- Author-only chirp deletion
- 140-character chirp limit and profanity filtering
- Chirpy Red upgrades through webhook
- Admin traffic metrics

## Requirements

- Go
- PostgreSQL
- Goose

## Setup

Clone the repository:

```bash
git clone https://github.com/nk-reddy/chirpy.git
cd chirpy
```

Create the database:

```bash
psql postgres -c "CREATE DATABASE chirpy;"
```

Create a `.env` file and configure the values used by your server:

```env
DB_URL=postgres://username:password@localhost:5432/chirpy?sslmode=disable
JWT_SECRET=replace-with-a-secure-secret
POLKA_KEY=replace-with-your-polka-api-key
PLATFORM=dev
```

Run the migrations:

```bash
goose -dir sql/schema postgres "$DB_URL" up
```

Start the server:

```bash
go run .
```

The examples below use `http://localhost:8080`.

## Authentication

Protected endpoints require a bearer token:

```http
Authorization: Bearer <token>
```

Login returns a one-hour JWT access token and a refresh token.

## API Endpoints

| Method | Endpoint | Authentication | Function |
|---|---|---|---|
| `GET` | `/api/healthz` | None | Check server health. |
| `POST` | `/api/users` | None | Create a user with an email and password. |
| `PUT` | `/api/users` | Access token | Update the authenticated user's email and password. |
| `POST` | `/api/login` | None | Log in and receive access and refresh tokens. |
| `POST` | `/api/refresh` | Refresh token | Create a new access token. |
| `POST` | `/api/revoke` | Refresh token | Revoke a refresh token. |
| `POST` | `/api/chirps` | Access token | Create a chirp for the authenticated user. |
| `GET` | `/api/chirps` | None | Return all chirps, oldest first by default. |
| `GET` | `/api/chirps?sort=desc` | None | Return chirps newest first. |
| `GET` | `/api/chirps?author_id={userID}` | None | Return chirps created by one user. |
| `GET` | `/api/chirps/{chirpID}` | None | Return one chirp by ID. |
| `DELETE` | `/api/chirps/{chirpID}` | Access token | Delete a chirp owned by the authenticated user. |
| `POST` | `/api/polka/webhooks` | Polka API key | Upgrade a user after a `user.upgraded` event. |
| `GET` | `/admin/metrics` | None | Display file-server visit metrics. |
| `POST` | `/admin/reset` | Development only | Reset development data, when enabled. |

Filtering and sorting can be combined:

```text
GET /api/chirps?author_id={userID}&sort=desc
```

## Example Requests

Create a user:

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password"}'
```

Log in:

```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password"}'
```

Create a chirp:

```bash
curl -X POST http://localhost:8080/api/chirps \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"body":"Hello from Chirpy!"}'
```

Delete a chirp:

```bash
curl -X DELETE http://localhost:8080/api/chirps/{chirpID} \
  -H "Authorization: Bearer $ACCESS_TOKEN"
```

## Chirp Rules

- Chirps may contain at most 140 characters.
- `kerfuffle`, `sharbert`, and `fornax` are replaced with `****`.
- Only a chirp's author may delete it.
