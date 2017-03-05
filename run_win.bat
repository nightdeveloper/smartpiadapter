set GOPATH=%cd%

call go install github.com/nightdeveloper/smartpiadapter/main
cd bin
call main.exe
cd ..