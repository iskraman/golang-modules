# golang-modules/syslog
Golang log 모듈화

### func DBG / ST / WARN / ERR
```
// Default log level
logLevel = DBG_LEVEL

func DBG(format string, v ...interface{})
func ST(format string, v ...interface{})
func WARN(format string, v ...interface{})
func ERR(format string, v ...interface{})
```

로그 레벨에 의한 출력 관리

[DBG] 2021/12/05 05:01:34 log 12 <br/>
[STD] 2021/12/05 05:01:34 log 34 <br/>
[WAN] 2021/12/05 05:01:34 log 56 <br/>
[ERR] 2021/12/05 05:01:34 log 78 <br/>
```