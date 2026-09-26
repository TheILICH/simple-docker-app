docker build -t simple .


docker run -p 8080:8080 -e APP_ENV=in-docker --name in-docker simple:latest
