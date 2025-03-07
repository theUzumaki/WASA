all:
	docker rmi -f wasa-text-backend:latest
	docker rmi -f wasa-text-frontend:latest
	docker pull node:18.12.0
	docker pull golang:1.19.1
	docker pull debian:bullseye-slim
	docker pull nginx:stable
	sudo docker build -t wasa-text-backend:latest -f Dockerfile.backend .
	sudo docker build -t wasa-text-frontend:latest -f Dockerfile.frontend .
	sudo docker run -it --rm -p 8080:80 wasa-text-frontend:latest
	sudo docker run -it --rm -p 3000:3000 --network host wasa-text-backend:latest
