# Modified

1. 通过 option 设置 log 的行为（flag 设置方式保留）

```go
    gglog.WithOptions(gglog.OptLogDir(".//log_dir/"), gglog.OptToStderr())
    gglog.Debug("some thing")
    gglog.Flush()
```

2. log 写入文件时，写入单独的文件
3. 添加 log threshold 小于等于 threshold 的日志不输出，可通过 http 设置
4. flushInterval 改为 5 秒
5. MaxSize 默认修改为 128M

# 例子

设置 logthreshold

```go
gglog.WithOptions(gglog.OptLogThreshold("DEBUG")) // INFO WARNING ERROR FATAL
gglog.WithOptions(gglog.OptLogThreshold("0"))     // 0		1     2     3
gglog.Debug("debug debug debug")
gglog.Info("info info info")

// output
20181128 16:39:29.310600 main.go:11 I] info info info
```

# glog

Leveled execution logs for Go.

This is an efficient pure Go implementation of leveled logs in the
manner of the open source C++ package
https://github.com/google/glog

By binding methods to booleans it is possible to use the log package
without paying the expense of evaluating the arguments to the log.
Through the -vmodule flag, the package also provides fine-grained
control over logging at the file level.

The comment from glog.go introduces the ideas:

    Package glog implements logging analogous to the Google-internal
    C++ INFO/ERROR/V setup.  It provides functions Info, Warning,
    Error, Fatal, plus formatting variants such as Infof. It
    also provides V-style logging controlled by the -v and
    -vmodule=file=2 flags.

    Basic examples:

    	glog.Info("Prepare to repel boarders")

    	glog.Fatalf("Initialization failed: %s", err)

    See the documentation for the V function for an explanation
    of these examples:

    	if glog.V(2) {
    		glog.Info("Starting transaction...")
    	}

    	glog.V(2).Infoln("Processed", nItems, "elements")

The repository contains an open source version of the log package
used inside Google. The master copy of the source lives inside
Google, not here. The code in this repo is for export only and is not itself
under development. Feature requests will be ignored.

Send bug reports to golang-nuts@googlegroups.com.
