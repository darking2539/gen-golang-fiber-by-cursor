### Database ER Diagram

```mermaid
erDiagram
    USERS {
        INTEGER id PK
        TEXT username UK
        TEXT password
        TEXT name
        TEXT email
        TEXT avatar
        TEXT bio
        TEXT title
        TEXT location
        TEXT website
        INTEGER followers
        INTEGER following
        INTEGER posts
    }
```

Notes:
- USERS is the single table used (JWT is stateless, no refresh tokens stored).
- Column types follow SQLite affinity.
- `username` is unique; `id` is the primary key.
