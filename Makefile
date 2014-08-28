.PHONY: all proto

all:
	go install -v ./...

proto:
	cd hdfs && protoc --go_out=. *.proto -I. -I../hadoop_common
	cd hdfs/datanode && protoc --go_out=. *.proto -I. -I.. -I../../hadoop_common
	cd hdfs/namenode && protoc --go_out=. *.proto -I. -I.. -I../../hadoop_common
	cd hdfs/namenode && protoc --go_out=. *.proto -I. -I.. -I../../hadoop_common
	cd hdfs/qjournal && protoc --go_out=. *.proto -I. -I.. -I../../hadoop_common
