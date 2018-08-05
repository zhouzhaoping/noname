#! /bin/bash
echo "build server..."
go build ~/pickme/backend/server/src/server/server.go
if [ $? -ne 0 ]
then
	echo "build fail"
	exit 1
fi

echo "build done"
nohup ./server 1>server.out 2>server.err &
if [ $? -ne 0 ]
then
        echo "run fail"
        exit 1
fi

echo "run server..."
