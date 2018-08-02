#! /bin/bash
echo "build server..."
go build server.go && nohup ./server 1>server.out 2>server.err &
