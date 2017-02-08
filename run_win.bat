set GOPATH=F:\Andrey\workspase.sh.go\

cd src/rpio
copy rpio.go.win rpio.go
cd ../../

call go install smartadapter
cd bin
call smartadapter.exe
cd ..