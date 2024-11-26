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
1. [Golang](https://go.dev/) v1.22.7 or above
2. [MongoDB](https://www.mongodb.com/)
3. [Buf](https://docs.buf.build/) to generate protobuf grpc (buf v1.47 or above)

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

You can see all endpoints at the following link: [Endpoints Go Architecture Monorepo](https://www.postman.com/bold-trinity-430312/workspace/go-architecture-monorepo)

### Notification Service
In the notification service you must save templates for email and FCM needs, here I will provide an example payload to save to the database. You can modify as you wish, you need to know that the use of email and FCM templates must be different. If they are the same then the concept of both should be the same as the welcome template.

**Welcome**:
```json
{
	"name": "welcome",
	"data": {
		"title": {
			"id": "Selamat Datang",
			"en": "Welcome"
		},
		"body": {
			"id": "Hai {{.name}}, kamu telah menjadi bagian dari kami.",
			"en": "Hi {{.name}}, you have become one of us."
		}
	}
}
```

**Email Verified**:
```json
{
	"name": "email-verified",
	"data": {
		"title": {
			"id": "Pemberitahuan Reset Kata Sandi",
			"en": "Reset Password Notification"
		},
		"link": {
			"id": "{{.link}}",
			"en": "{{.link}}"
		}
	}
}
```

**Password Reset**:
```json
{
	"name": "password-reset",
	"data": {
		"title": {
			"id": "Pemberitahuan Reset Kata Sandi",
			"en": "Reset Password Notification"
		},
		"link": {
			"id": "{{.link}}",
			"en": "{{.link}}"
		},
		"expire": {
			"id": "{{.expire}}",
			"en": "{{.expire}}"
		}
	}
}
```

> You can add templates via endpoint `/v1/notification/template`.