#! /bin/bash
mypwd="/root/pickme/backend/server"
if [[ $GOPATH =~ .*$mypwd.* ]]
then
    echo "currnet project is already in GOPATH"
else
    export GOPATH=$GOPATH:$mypwd
    echo "fininsh setting $mypwd in GOPATH"
fi

#mypwd="/root/go"
#if [[ $GOPATH =~ .*$mypwd.* ]]
#then
#    echo "currnet project is already in GOPATH"
#else
#    export GOPATH=$GOPATH:$mypwd
#    echo "fininsh setting $mypwd in GOPATH"
#fi

echo "build server..."
go build ../backend/server/src/server/server.go
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
