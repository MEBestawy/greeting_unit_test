# Demo Of Environment Variables In Golang Unit Tests

This repo is a demo of how to read environment variables in Golang unit tests using `godotenv`.

## How To Run The Tests

```bash
godotenv -f ./.env go test    
```

## How To Run The Tests Without Environment Variables

Notice how this test will fail as it does not have access to the necessary environment variables! 
```bash
go test
```