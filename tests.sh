#!/bin/sh
cd action
go test 
cd ..
cd operations
for op  in $(ls -d ./* | grep -v sample_README.MD)
do 
 cd $op
		go test 
 cd ../..
done
​
