#!/bin/sh
​
pushd action
​
go test 
​
popd
​
pushd operations
​
for op  in $(ls -d ./* | grep -v sample_README.MD)
do 
 pushd $op
		go test 
 popd
done
​
