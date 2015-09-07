export GIN_MODE=release
SERV="ant-worker"
go build
sudo ./$SERV install
sudo ./$SERV start
go test ./... -v
sudo ./$SERV stop
sudo ./$SERV uninstall