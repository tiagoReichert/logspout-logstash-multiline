# Logspout with logstash and multiline adpater

gliderlabs/logspout container with logstash and multiline adpater

## docker-compose 
Multiline example for java error stack trace message pattern
```yaml
version: '2'
services:
  logspout:
    image: tiagoreichert/logspout-logstash-multiline
    environment:
      ROUTE_URIS: multiline+logstash://<logstash-address>:5000
      MULTILINE_PATTERN: '(^.+Exception: .+)|(^\s+at .+)|(^\s+... \d+ more)|(^\s*Caused by:.+)'
    volumes:
    - /var/run/docker.sock:/tmp/docker.sock
```
