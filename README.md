# Music Library Management Interview

![Logo](./logo.jpeg)

# Structure

```
├── api // OpenAPI/Swagger specs.
├── cmd // Main applications
├── config // Contain configuration
│   └── config.go // Environment variables
├── domain // Domain layer, interface for database layer
│   ├── entity // Model interface
│   └── repository // Repository interface
├── error // Error constants
├── handler // Controller, handle request
├── infra // Database layer
│   ├── mongo
│   │    ├── connect.go // Connect to mongo database
│   │    └── repository // Database repository
│   └── s3
│   │    ├── connect.go // Connect to aws s3
│   │    └── repository // Storage repository
├── middlewares // Put the before & after logic of handle request
├── internal // Package internal project for service
├── routers // Router for services use REST API
├── usecase // Business logic
├── validations // custom validation
├── vendor // Management dependency package
├── .dockerignore
├── .env.debug // Env for debug vs code
├── .env.example //Env for docker container
├── .gitignore
├── docker-compose.yml // Container configuration
├── Dockerfile // Define container image golang
├── go.mod
├── go.sum
├── logo.jpeg
└── README.md
```

# Instruction
## For Running (cd root folder)
```
1/ cp .env.example .env
2/ Setup account IAM AWS
    2.1/ Create account IAM
    2.2/ Create file accessKeys for account (Access key ID + Secret access key)
    2.3/ Copy value of Access key ID and fill to .env variable AWS_ACCESS_KEY
    2.3/ Copy value of Secret access key and fill to .env variable AWS_SECRET_KEY
    2.4/ Create a new bucket s3 (any name) and fill to .env variable S3_BUCKET_ID
3/ docker-compose up -d
4/ Check swagger on link: http://localhost:8080/
5/ curl -X 'GET' \ 'http://localhost:8000/api/v1/ping' \ -H 'accept: */*'
it works 🎉
```

## For develop (cd root folder)
```
1/ Stop container app running.
2/ Install extension recommend.
3/ Install go debug package https://github.com/go-delve/delve .
4/ Enter Ctrl+ Shift + D -> Launch Package.
it works 🎉
```

## Question
1/ I don't have an account, how can I run the application?
Answer:

```
1/ You can use api auth register to register an account (check swagger).
2/ Log in with the account you have just registered.
3/ Use the program's functions after logging in (access_token has been set).
```

## Tech Stack

**Server:** Golang, Gin, Gorm, MongoDB, AWS S3

## Authors

- [@SeangThai](https://github.com/aminofox)

## Feedback
If you have any feedback, please contact to us at aminofox165@gmail.com