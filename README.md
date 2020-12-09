# gin-curd-demo

Sample Demo based on GIN framework.
- Logging
- Config

# Running Outputs

---

REQUEST:
http://localhost:8080

RESPONSE:
404 page not found

REQUEST:
http://localhost:8080/

RESPONSE:
404 page not found

REQUEST:
http://localhost:8080/info

RESPONSE:
{"appName":"gin-curd-demo","lastCommitId":"","kubePodId":"LAP-PSINGH","UpSince":"Sun, 01 Nov 2020 18:30:41 IST"}

REQUEST:
http://localhost:8080/health

RESPONSE:
{"status":"ok"}

---

C:\Users\Home\Desktop\gin-curd-demo_1>go build .

C:\Users\Home\Desktop\gin-curd-demo_1>set runEnv=dev

C:\Users\Home\Desktop\gin-curd-demo_1>gin-curd-demo.exe
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /info                     --> github.com/pksingh/gin-curd-demo/handler.GetInfo (3 handlers)
[GIN-debug] GET    /health                   --> github.com/pksingh/gin-curd-demo/handler.GetHealth (3 handlers)        
[GIN-debug] GET    /debug/pprof/             --> github.com/gin-gonic/gin.WrapF.func1 (3 handlers)
[GIN-debug] GET    /debug/pprof/cmdline      --> github.com/gin-gonic/gin.WrapF.func1 (3 handlers)
[GIN-debug] GET    /debug/pprof/profile      --> github.com/gin-gonic/gin.WrapF.func1 (3 handlers)
[GIN-debug] GET    /debug/pprof/symbol       --> github.com/gin-gonic/gin.WrapF.func1 (3 handlers)
[GIN-debug] GET    /debug/pprof/goroutine    --> github.com/gin-gonic/gin.WrapH.func1 (3 handlers)
[GIN-debug] GET    /debug/pprof/heap         --> github.com/gin-gonic/gin.WrapH.func1 (3 handlers)
[GIN-debug] GET    /debug/pprof/threadcreate --> github.com/gin-gonic/gin.WrapH.func1 (3 handlers)
[GIN-debug] GET    /debug/pprof/block        --> github.com/gin-gonic/gin.WrapH.func1 (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2020/11/01 - 18:31:16 | 404 |            0s |             ::1 | GET      "/"
[GIN] 2020/11/01 - 18:31:21 | 404 |            0s |             ::1 | GET      "/"
[GIN] 2020/11/01 - 18:31:28 | 200 |       154.8µs |             ::1 | GET      "/info"
[GIN] 2020/11/01 - 18:32:56 | 404 |            0s |             ::1 | GET      "/"
[GIN] 2020/11/01 - 18:33:41 | 200 |       604.7µs |             ::1 | GET      "/health"
---
## Test case Added

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^TestStartUpErr$ github.com/pksingh/gin-curd-demo

=== RUN   TestStartUpErr
2020-11-03T20:55:47.307+0530    PANIC   gin-curd-demo_1/main.go:43      failed to start application     {"appCrashed": "i want a startup err"}
github.com/pksingh/gin-curd-demo.handleStartUpErr
        c:/Users/Home/Desktop/gin-curd-demo_1/main.go:43
github.com/pksingh/gin-curd-demo.TestStartUpErr
        c:/Users/Home/Desktop/gin-curd-demo_1/main_test.go:47
testing.tRunner
        C:/DEV/Go/src/testing/testing.go:1446
--- PASS: TestStartUpErr (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo        (cached)


> Test run finished at 11/4/2020, 9:13:29 PM <

---

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^TestMiddlewaresInjections$ github.com/pksingh/gin-curd-demo/config/logperreq

=== RUN   TestMiddlewaresInjections
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

--- PASS: TestMiddlewaresInjections (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo/config/logperreq       0.909s


> Test run finished at 11/3/2020, 9:06:45 PM <

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^TestNewReqId$ github.com/pksingh/gin-curd-demo/config/logperreq

=== RUN   TestNewReqId
--- PASS: TestNewReqId (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo/config/logperreq       0.909s


> Test run finished at 11/3/2020, 9:07:05 PM <

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^TestGetRequestIdNoCtx$ github.com/pksingh/gin-curd-demo/config/logperreq

=== RUN   TestGetRequestIdNoCtx
--- PASS: TestGetRequestIdNoCtx (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo/config/logperreq       0.872s


> Test run finished at 11/3/2020, 9:07:34 PM <

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^TestBadLogConf$ github.com/pksingh/gin-curd-demo/config/jlog

=== RUN   TestBadLogConf
{"level":"error","ts":1604417897.1259308,"caller":"jlog/logConfig.go:49","msg":"failed to initialize zap logger with given configuration.defaulting to zap-prod configuration.","error":"missing EncodeTime in EncoderConfig","stacktrace":"github.com/pksingh/gin-curd-demo/config/jlog.ZapAppConf.CustomizeLogger\n\tc:/Users/Home/Desktop/gin-curd-demo_1/config/jlog/logConfig.go:49\ngithub.com/pksingh/gin-curd-demo/config/jlog.TestBadLogConf\n\tc:/Users/Home/Desktop/gin-curd-demo_1/config/jlog/logConfig_test.go:12\ntesting.tRunner\n\tC:/DEV/Go/src/testing/testing.go:1446"}
--- PASS: TestBadLogConf (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo/config/jlog    0.920s


> Test run finished at 11/3/2020, 9:08:17 PM <

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^TestFetchLogLevels$ github.com/pksingh/gin-curd-demo/config/jlog

=== RUN   TestFetchLogLevels
--- PASS: TestFetchLogLevels (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo/config/jlog    0.749s


> Test run finished at 11/3/2020, 9:09:09 PM <

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^TestBaseZapConf$ github.com/pksingh/gin-curd-demo/config/jlog

=== RUN   TestBaseZapConf
--- PASS: TestBaseZapConf (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo/config/jlog    0.742s


> Test run finished at 11/3/2020, 9:09:16 PM <

---

http.timeout.sample=1500

---

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^(TestGetInfo|TestGetHealth)$ github.com/pksingh/gin-curd-demo/handler

=== RUN   TestGetHealth
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

--- PASS: TestGetHealth (0.00s)
=== RUN   TestGetInfo
--- PASS: TestGetInfo (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo/handler        0.219s


> Test run finished at 11/10/2020, 10:53:10 PM <

---

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^(TestBasicAuthConfNotLoaded|TestBasicAuthApply)$ github.com/pksingh/gin-curd-demo/startup/middlewares/basicAuth

=== RUN   TestBasicAuthConfNotLoaded
=== RUN   TestBasicAuthConfNotLoaded/basic_auth_not_loaded
--- PASS: TestBasicAuthConfNotLoaded (0.00s)
    --- PASS: TestBasicAuthConfNotLoaded/basic_auth_not_loaded (0.00s)
=== RUN   TestBasicAuthApply
=== RUN   TestBasicAuthApply/basic_auth_env_var_not_set
warn    logger not found in context. defaulting to zap-prod configuration.
    c:\Users\Home\Desktop\gin-curd-demo_1\startup\middlewares\basicAuth\basic_test.go:31:
                Error Trace:    basic_test.go:31
                Error:          An error is expected but got nil.
                Test:           TestBasicAuthApply/basic_auth_env_var_not_set
=== RUN   TestBasicAuthApply/basic_auth_env_var_invalid
warn    logger not found in context. defaulting to zap-prod configuration.
    c:\Users\Home\Desktop\gin-curd-demo_1\startup\middlewares\basicAuth\basic_test.go:37:
                Error Trace:    basic_test.go:37
                Error:          An error is expected but got nil.
                Test:           TestBasicAuthApply/basic_auth_env_var_invalid
=== RUN   TestBasicAuthApply/basic_auth_env_no_username_or_password
warn    logger not found in context. defaulting to zap-prod configuration.
    c:\Users\Home\Desktop\gin-curd-demo_1\startup\middlewares\basicAuth\basic_test.go:43:
                Error Trace:    basic_test.go:43
                Error:          An error is expected but got nil.
                Test:           TestBasicAuthApply/basic_auth_env_no_username_or_password
=== RUN   TestBasicAuthApply/happy_path
warn    logger not found in context. defaulting to zap-prod configuration.
--- FAIL: TestBasicAuthApply (0.00s)
    --- FAIL: TestBasicAuthApply/basic_auth_env_var_not_set (0.00s)
    --- FAIL: TestBasicAuthApply/basic_auth_env_var_invalid (0.00s)
    --- FAIL: TestBasicAuthApply/basic_auth_env_no_username_or_password (0.00s)
    --- PASS: TestBasicAuthApply/happy_path (0.00s)
FAIL
FAIL    github.com/pksingh/gin-curd-demo/startup/middlewares/basicAuth  0.267s


> Test run finished at 11/28/2020, 11:35:34 PM <

---

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^TestGetAll$ github.com/pksingh/gin-curd-demo/startup/appProps

=== RUN   TestGetAll
=== RUN   TestGetAll/GetAllWithoutLoad
=== RUN   TestGetAll/GetAllSuccess
configuration loading err - open ./resources/app.properties: The system cannot find the path specified..
 Empty Properties returned
--- PASS: TestGetAll (0.00s)
    --- PASS: TestGetAll/GetAllWithoutLoad (0.00s)
    --- PASS: TestGetAll/GetAllSuccess (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo/startup/appProps       0.343s


> Test run finished at 11/28/2020, 11:38:14 PM <

---

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^(TestCreateServer|TestCreateServerInit|Test_randSeq|Test_generateRandomBytes)$ github.com/pksingh/gin-curd-demo/startup/server

=== RUN   TestCreateServerInit
--- PASS: TestCreateServerInit (0.00s)
=== RUN   Test_randSeq
--- PASS: Test_randSeq (0.00s)
=== RUN   Test_generateRandomBytes
--- PASS: Test_generateRandomBytes (0.00s)
=== RUN   TestCreateServer
--- PASS: TestCreateServer (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo/startup/server 0.346s


> Test run finished at 11/30/2020, 11:47:44 PM <

---

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^(TestDB|TestDBError)$ github.com/pksingh/gin-curd-demo/startup/db

=== RUN   TestDB
configuration loading err - open ../../../resources/app.properties: The system cannot find the path specified..
 Empty Properties returned
warn    logger not found in context. defaulting to zap-prod configuration.
info    connection string:      {"host": "", "port": "0", "user": "", "dbname": "", "sslmode": ""}
cleartextPwd PWD=
error   sql.Open failed, error:         {"error": "failed to connect to `host=port=0 user=password= database=sslmode=`: hostname resolving error (lookup port=0: no such host)"}
--- PASS: TestDB (0.01s)
=== RUN   TestDBError
warn    logger not found in context. defaulting to zap-prod configuration.
info    connection string:      {"host": "", "port": "0", "user": "", "dbname": "", "sslmode": ""}
cleartextPwd PWD=
error   sql.Open failed, error:         {"error": "failed to connect to `host=port=0 user=password= database=sslmode=`: hostname resolving error (lookup port=0: no such host)"}
--- PASS: TestDBError (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo/startup/db     0.306s


> Test run finished at 12/6/2020, 12:15:42 AM <

---

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^(TestVerifyLogConfHappy|TestBadLogConf|TestGetLoggerWithNilCtxt|TestGetLoggerWithCustomTypeCtxt|TestGetLoggerWithStringCtxt|TestGetLoggerWithNotFoundInCtxt|TestGetRequestIdNotFoundInCtxt)$ github.com/pksingh/gin-curd-demo/log

=== RUN   TestVerifyLogConfHappy
--- PASS: TestVerifyLogConfHappy (0.00s)
=== RUN   TestBadLogConf
{"level":"error","ts":1607194008.0438504,"caller":"jlog/logConfig.go:49","msg":"failed to initialize zap logger with given configuration.defaulting to zap-prod configuration.","error":"missing EncodeTime in EncoderConfig","stacktrace":"github.com/pksingh/gin-curd-demo/config/jlog.ZapAppConf.CustomizeLogger\n\tc:/Users/Home/Desktop/gin-curd-demo_1/config/jlog/logConfig.go:49\ngithub.com/pksingh/gin-curd-demo/log.TestBadLogConf\n\tc:/Users/Home/Desktop/gin-curd-demo_1/log/logConf_test.go:24\ntesting.tRunner\n\tC:/DEV/Go/src/testing/testing.go:1446"}
--- PASS: TestBadLogConf (0.00s)
=== RUN   TestGetLoggerWithNilCtxt
error   input context is nil. defaulting to zap-prod configuration.
--- PASS: TestGetLoggerWithNilCtxt (0.00s)
=== RUN   TestGetLoggerWithCustomTypeCtxt
--- PASS: TestGetLoggerWithCustomTypeCtxt (0.00s)
=== RUN   TestGetLoggerWithStringCtxt
--- PASS: TestGetLoggerWithStringCtxt (0.00s)
=== RUN   TestGetLoggerWithNotFoundInCtxt
warn    logger not found in context. defaulting to zap-prod configuration.
--- PASS: TestGetLoggerWithNotFoundInCtxt (0.00s)
=== RUN   TestGetRequestIdNotFoundInCtxt
--- PASS: TestGetRequestIdNotFoundInCtxt (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo/log    0.229s


> Test run finished at 12/6/2020, 12:16:48 AM <

---

Running tool: C:\DEV\Go\bin\go.exe test -timeout 30s -run ^TestStartUpErr$ github.com/pksingh/gin-curd-demo

warn    logger not found in context. defaulting to zap-prod configuration.
info    connection string:      {"host": "", "port": "0", "user": "", "dbname": "", "sslmode": ""}
cleartextPwd PWD=
error   sql.Open failed, error:         {"error": "failed to connect to `host=port=0 user=password= database=sslmode=`: hostname resolving error (lookup port=0: no such host)"}
warn    logger not found in context. defaulting to zap-prod configuration.
=== RUN   TestStartUpErr
2020-12-08T00:23:33.070+0530    PANIC   gin-curd-demo_1/main.go:51      failed to start application      {"appCrashed": "i want a startup err"}
github.com/pksingh/gin-curd-demo.handleStartUpErr
        c:/Users/Home/Desktop/gin-curd-demo_1/main.go:51
github.com/pksingh/gin-curd-demo.TestStartUpErr
        c:/Users/Home/Desktop/gin-curd-demo_1/main_test.go:51
testing.tRunner
        C:/DEV/Go/src/testing/testing.go:1446
--- PASS: TestStartUpErr (0.00s)
PASS
ok      github.com/pksingh/gin-curd-demo        0.352s


> Test run finished at 12/8/2020, 12:23:33 AM <

---
