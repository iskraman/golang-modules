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

로그 레벨에 의한 출력 관리

example)
[DBG] 2021/12/05 05:01:34 log 12
<span color=cyan>[STD]</span> 2021/12/05 05:01:34 log 34
<span color=purple>[WAN]</span> 2021/12/05 05:01:34 log 56
<span color=red>[ERR]</span> 2021/12/05 05:01:34 log 78
```
