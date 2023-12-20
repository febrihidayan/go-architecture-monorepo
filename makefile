auth-build:
	docker build --pull -t febrihidayan/go-architecture-monorepo/auth:latest . -f docker/auth/Dockerfile

user-build:
	docker build --pull -t febrihidayan/go-architecture-monorepo/user:latest . -f docker/user/Dockerfile