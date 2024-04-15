# [PoC] haproxy-rsyslogs-kafka

This is a PoC project to try sending incoming requests from haproxy to kafka through rsyslogs

## Setup

- Run the whole pipeline using - `make up`. NB: you should have a running container runtime such as docker or colima.
- Simulate network requests - `make requests`
- View the messages in kafka 
  - Visit `http://localhost:8080/ui/clusters/4L6g3nShT-eMCtK--X86sw/all-topics/haproxy-messages/messages` to see messages stored in the kafka topic

## Reuslts 

a request of the form 
```sh
curl -X POST http://localhost:8100/json-reqeust?id=abcd1234&body=full -H "Content-Type: application/json" -d '{"username":"jon_snow","house":"Targaryen,Stark"}'
```
generates a log line in kafka such as 
```json
{
  "host": "a7a674c4886d",
  "queue": {
    "backend": 0,
    "srv": 0
  },
  "network": {
    "client_ip": "172.21.0.1"
  },
  "request": {
    "method": "POST",
    "uri": "/json-reqeust?id=abcd1234&body=full",
    "protocol": "HTTP/1.1",
    "headers": "host: localhost:8100\r\nuser-agent: curl/8.4.0\r\naccept: */*\r\ncontent-type: application/json\r\ncontent-length: 49\r\n\r\n",
    "body": "{\"username\":\"jon_snow\",\"house\":\"Targaryen,Stark\"}"
  },
  "response": {
    "status_code": 200
  }
}
```

## Questions

- What happens when haproxy cannot resolve a log host
  - It continues to work as usual.
- What happens when rsyslog cannot resolve a kafka host
  - it continues to work as usual.
- Does haproxy write to file before rsyslog collects and ships to kafka, - Answered
  - logs are sent directly to a socket (udp in this case). Rsyslog collects these logs and writes them to kafka directly without storing on file
- how does logging performance
- Truncation of log lines
  - Haproxy has a maximum limit for log lines beyond which it will truncate
- Logging request body only possible post 1.6.0 (October 2015)

