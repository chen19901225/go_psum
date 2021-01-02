package log

import (
	"log"
	"os"
)

// DefaultLogger 默认的日志实例
var DefaultLogger = log.New(os.Stdout, "", log.LstdFlags)
