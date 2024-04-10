# [PoC] haproxy-rsyslogs-kafka

This is a PoC project to try sending incoming requests from haproxy to kafka through rsyslogs

## Questions
- what happens when haproxy cannot resolve a host
- does ha proxy write to file before rsyslog collects and ships to kafka, 