PHONY: all build app
up:
	docker-compose -f docker-compose.yml up --build --remove-orphans

requests:
	for i in {1..10}; do \
		curl -X POST http://localhost:8100/form-reqeust?id=$$i\&body=full -H "Content-Type: application/x-www-form-urlencoded" -d "param1=value1&param2=value2"; \
		curl -X POST http://localhost:8100/json-reqeust?id=$$i\&body=full -H "Content-Type: application/json" -d '{"username":"jon_snow","house":"Targaryen,Stark"}'; \
	done;
	#&& for j in {500..1000}; curl -X POST http://localhost:8100/json-reqeust?id=$j\?body=full -H "Content-Type: application/json" -d '{"username":"jon_snow","house":"Targaryen,Stark"}'
