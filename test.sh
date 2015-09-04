SERV="ant-worker"
go build
./$SERV install
./$SERV start
go test ./... -v
./$SERV stop
./$SERV uninstall