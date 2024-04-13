# [PoC] haproxy-rsyslogs-kafka

This is a PoC project to try sending incoming requests from haproxy to kafka through rsyslogs

## Setup

- Run the whole pipeline using - `make up`. NB: you should have a running container runtime such as docker or colima.
- Simulate network requests - `make requests`
- View the messages in kafka 
  - Visit `http://localhost:8080/ui/clusters/4L6g3nShT-eMCtK--X86sw/all-topics/haproxy-messages/messages` to see messages stored in the kafka topic

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

