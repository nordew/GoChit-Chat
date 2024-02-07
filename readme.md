# GoChit-Chat

GoChit-Chat is a simple chat application built in Go using microservice architecture and websockets.

## Overview

This project consists of three main services:

1. **UserService**:
    - Utilizes gRPC as transport.
    - PostgreSQL as the database.
    - JWT Tokens for authorization.

2. **Chat**:
    - Provides rooms and the ability to join them.
    - Uses gin + websockets for transport.
    - PostgreSQL as the database.

3. **Gateway**:
    - Offers an API Gateway to all services based on gRPC.
    - Utilizes Fiber as the router.

Each service has its own Dockerfile, configuration, and migrations, and they are built using a unified Docker Compose file.

## Endpoints

### Gateway

- Authentication Endpoints:
    - `POST /auth/sign-up`: Sign up a new user.
    - `GET /auth/sign-in`: Sign in an existing user.
    - `GET /auth/refresh`: Refresh JWT token.

### Chat

- Room Management:
    - `POST /create-room`: Create a new chat room.
    - `GET /ws/join-room/:roomID`: Join a specific chat room via websocket.

## Request Examples

### SignUpRequest

```json
{
    "name": "string",
    "email": "string",
    "password": "string"
}
```

### SignInRequest
```json
{
    "email": "string",
    "password": "string"
}
```

### CreateRoomRequest
```json
{
    "name": "string"
}
```

## Usage
1. Clone the repository.
2. Navigate to the project directory.
3. docker-compose --env-file .env -f docker-compose.yml up and run the services.

### Don't forget to define .env configs for each services

Enjoy chatting with GoChit-Chat! 🚀