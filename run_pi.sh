export GOPATH="$PWD"

go install github.com/nightdeveloper/smartpiadapter/main
cd bin
./main
cd ..