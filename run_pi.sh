export GOPATH="$PWD"

go install github.com/nightdeveloper/smartpiadapter/main
cd bin
sudo ./main
cd ..