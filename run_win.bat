set GOPATH=%cd%

cd src/rpio
copy rpio.go.win rpio.go
cd ../../

call go install smartadapter
cd bin
call smartadapter.exe
cd ..