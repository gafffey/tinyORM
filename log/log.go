package log

import (
	"log"
	"os"
	"sync"
)

/*
*
* [info ] 颜色为蓝色，[error] 为红色
* 使用 log.Lshortfile 支持显示文件名和代码行号
* 暴露 Error，Errorf，Info，Infof 4个方法
*
 */
var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[31m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)
