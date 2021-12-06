# golang-modules/syslog
Golang log 모듈화

### func SetLogLevel
로그 레벨 세팅
```
const (
	DBG_LEVEL = 1
	STD_LEVEL = 2
	WAR_LEVEL = 3
	ERR_LEVEL = 4
)

func SetLogLevel(level int)
```

### func DBG | STD | WAR | ERR
로그 레벨에 의한 출력 관리
```
func DBG(format string, v ...interface{})
func STD(format string, v ...interface{})
func WAR(format string, v ...interface{})
func ERR(format string, v ...interface{})

func DBGLN(v ...interface{})
func STDLN(v ...interface{})
func WARLN(v ...interface{})
func ERRLN(v ...interface{})
```
[DBG] 2021/12/05 05:01:34 log 12 <br/>
[STD] 2021/12/05 05:01:34 log 34 <br/>
[WAR] 2021/12/05 05:01:34 log 56 <br/>
[ERR] 2021/12/05 05:01:34 log 78 <br/>
```
