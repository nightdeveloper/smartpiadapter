package logger

import (
	"fmt"
	"time"
)

func CurrentTime() string {
	return time.Now().Format("2006.01.02 15:04:05.00000");
}

func Log(str string) {
	fmt.Println(CurrentTime() + " " + str)
}

func Info(str string) {
	Log("INFO " + str);
}

func Error(str string, err error) {
	if err != nil {
		Log("ERROR " + str + ": " + err.Error());
	} else {
		Log("ERROR " + str);
	}
}