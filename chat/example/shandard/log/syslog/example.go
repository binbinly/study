package main

import (
	"log"
	"log/syslog"
)

// syslog 包提供一个简单的系统日志服务的接口
func main()  {
	sl, err := syslog.Dial("tcp", "localhost:1234",
		syslog.LOG_WARNING|syslog.LOG_DAEMON|syslog.LOG_INFO, "testTag")
	if err != nil {
		log.Fatal(err)
	}

	if _, err = sl.Write([]byte("hello world")); err != nil {
		log.Fatal(err)
	}

	if err = sl.Close(); err != nil {
		log.Fatal(err)
	}

	news, err := syslog.New(syslog.LOG_DEBUG, "testNew")
	if err != nil {
		log.Fatal(err)
	}
	defer news.Close()

	l, err := syslog.NewLogger(syslog.LOG_DEBUG, log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		log.Fatal(err)
	}
	l.Fatal("退出")
}