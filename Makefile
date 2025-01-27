all:
	sudo docker build -t wasa-text-backend:latest -f Dockerfile.backend .
	sudo docker build -t wasa-text-frontend:latest -f Dockerfile.frontend .
	sudo docker run -it --rm -p 3000:3000 wasa-text-backend:latest