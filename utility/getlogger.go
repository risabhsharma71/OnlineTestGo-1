package utility

import (
	
	"log"
        "gopkg.in/natefinch/lumberjack.v2"
)

func GetLogger(){

log.SetOutput(&lumberjack.Logger{
    Filename:   "/home/arun/gopath/src/OnlineTestGo/log.txt",
    MaxSize:    500, // megabytes
    MaxBackups: 3,
    MaxAge:     1, //days
})
return
}