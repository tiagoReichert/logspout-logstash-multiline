package main

import (
  _ "github.com/looplab/logspout-logstash"
  _ "github.com/gliderlabs/logspout/adapters/multiline"
  _ "github.com/gliderlabs/logspout/transports/udp"
  _ "github.com/gliderlabs/logspout/transports/tcp"
  _ "github.com/gliderlabs/logspout/transports/tls"
)
