#!/bin/sh
rtn=0
cd action
go test ./...
if [ $? -eq 0 ]
then
    echo "The script ran ok"
else
    echo "action tests failed" >&2
    rtn=$(($rtn+1))
fi
cd ..
cd operations
for op  in $(find . -mindepth 2 -maxdepth 2 -type d)
do 
    cd $op
    go test
    if [ $? -eq 0 ]
    then
        echo "The script ran ok"
    else
        echo "operations/$op failed" >&2
        rtn=$(($rtn+1))
    fi
    go test 
    cd ../..
done
exit $rtn