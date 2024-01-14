# Go Architecture Monorepo
Started rest API with clean architecture monorepo golang

**Table of Contents**
- [Go Architecture Monorepo](#go-architecture-monorepo)
  - [Required](#required)
  - [Installation](#installation)
  - [With Docker](#with-docker)
  - [Port HTTP and RPC Services](#port-http-and-rpc-services)
  - [Docs Services](#docs-services)
    - [Notification Service](#notification-service)

## Required
1. [Golang](https://go.dev/) v1.21.5 or above
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
| notification | 8086      | 9086     |

## Docs Services
Here I will provide the information that the service needs to set some default configurations.

### Notification Service
In the notification service you must store templates for email and FCM needs, here I will provide sample data from the database. You can modify as you wish, you need to know that the use of email and FCM templates must be different. If they are the same then the concept of both must be the same as the welcome template.

**Welcome**:
```json
{
  "_id": "458c523a-b5a0-4ca0-9915-b389bae31770",
  "name": "welcome",
  "type": "email",
  "data": "{"body":{"en":"Hi {{.name}}, you have become one of us.","id":"Hai {{.name}}, kamu telah menjadi bagian dari kami."},"title":{"en":"Welcome","id":"Selamat Datang"}}",
  "created_at": "2024-01-11T07:52:59.953+00:00",
  "updated_at": "2024-01-11T07:52:59.953+00:00"
}
```

**Email Verified**:
```json
{
  "_id": "80faa307-7f5d-4c2c-9d66-05f77641b161",
  "name": "email-verified",
  "data": "{"link":{"en":"{{.link}}","id":"{{.link}}"},"title":{"en":"Verify Email Address","id":"Verifikasi Alamat Email"}}",
  "created_at": "2024-01-11T07:52:59.953+00:00",
  "updated_at": "2024-01-11T07:52:59.953+00:00"
}
```

**Password Reset**:
```json
{
  "_id": "2adbcd01-ae67-491e-a95d-72faa78cd246",
  "name": "password-reset",
  "data": "{"expire":{"en":"{{.expire}}","id":"{{.expire}}"},"link":{"en":"{{.link}}","id":"{{.link}}"},"title":{"en":"Reset Password Notification","id":"Pemberitahuan Reset Kata Sandi"}}",
  "created_at": "2024-01-11T07:52:59.953+00:00",
  "updated_at": "2024-01-11T07:52:59.953+00:00"
}
```

> You can add templates via endpoint `/v1/notification/template`.