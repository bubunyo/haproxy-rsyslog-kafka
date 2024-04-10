PHONY: all build app
build:
	docker-compose -f docker-compose.yml up --remove-orphans

app:
	docker build -t main ./app && docker run -p 8090:8090 -t main