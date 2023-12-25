# Go Architecture Monorepo
Started rest API with clean architecture monorepo golang

**Table of Contents**
- [Go Architecture Monorepo](#go-architecture-monorepo)
  - [Required](#required)
  - [Installation](#installation)
  - [With Docker](#with-docker)
  - [Port HTTP and RPC Services](#port-http-and-rpc-services)

## Required
1. [Golang](https://go.dev/) v1.18.1 or above
2. [MongoDB](https://www.mongodb.com/)
3. [Buf](https://docs.buf.build/) to generate protobuf grpc

## Installation
1. copy environment by services `cp .env.example .env`
2. sync go modules `go mod tidy`
3. Run `sh run-service.sh service-name`

## With Docker
1. Build for all services `docker-compose build`
2. Run docker `docker-compose up -d`
3. See docker running `docker ps`

## Port HTTP and RPC Services
| Service Name | Port HTTP | Port RPC |
| :----------- | :-------- | :------- |
| auth         | 8083      | 9083     |
| user         | 8084      | 9084     |
| storage      | 8085      | 9085     |
