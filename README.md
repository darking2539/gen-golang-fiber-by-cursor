# go-fiber clean architecture sample

## Run

```bash
go run ./cmd/server
```

Then open `http://localhost:8080/hello-world` -> `{ "message": "hello world" }`



## Graceful shutdown
- SIGINT/SIGTERM triggers `ShutdownWithContext` with 5s timeout.

## Prompt
```
Please generate project go-fiber with belows detail
- /hello-world for response "hello world" in json
- Please implement graceful shutdown.
- Use clean-arct structure 
```