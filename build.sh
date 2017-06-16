export GOPATH=$(pwd)/gopath
go build -buildmode=c-shared -o foolproc.so github.com/taowen/foolproc