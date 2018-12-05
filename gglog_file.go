// Go support for leveled logs, analogous to https://code.google.com/p/google-glog/
//
// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// File I/O for logs.

package gglog

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// MaxSize is the maximum size of a log file in bytes.
var MaxSize uint64 = 1024 * 1024 * 512

// logDirs lists the candidate directories for new log files.
var logDirs []string

// If non-empty, overrides the choice of directory in which to write logs.
// See createLogDirs for the full list of possible destinations.
var logDir = flag.String("log_dir", "", "If non-empty, write log files in this directory")

func createLogDirs() {
	if *logDir != "" {
		logDirs = append(logDirs, *logDir)
	}
	logDirs = append(logDirs, os.TempDir())
}

var (
	// 	pid      = os.Getpid()
	// 	program  = filepath.Base(os.Args[0])
	host = "unknownhost"

// 	userName = "unknownuser"
)

//
func init() {
	h, err := os.Hostname()
	if err == nil {
		host = shortHostname(h)
	}
	//
	// 	current, err := user.Current()
	// 	if err == nil {
	// 		userName = current.Username
	// 	}
	//
	// 	// Sanitize userName since it may contain filepath separators on Windows.
	// 	userName = strings.Replace(userName, `\`, "_", -1)
}

// shortHostname returns its argument, truncating at the first period.
// For instance, given "www.google.com" it returns "www".
func shortHostname(hostname string) string {
	if i := strings.Index(hostname, "."); i >= 0 {
		return hostname[:i]
	}
	return hostname
}

// logName returns a new log file name containing tag, with start time t, and
// the name for the symlink for tag.
func logName(t time.Time) (name string) {
	// name = fmt.Sprintf("%s.%s.%s.log.%s.%04d%02d%02d-%02d%02d%02d.%d",
	// 	program,
	// 	host,
	// 	userName,
	// 	tag,
	// 	t.Year(),
	// 	t.Month(),
	// 	t.Day(),
	// 	t.Hour(),
	// 	t.Minute(),
	// 	t.Second(),
	// 	pid)
	// 按大小分隔 文件名格式 20180102-183824.log
	// 按日期分隔 文件名格式 20180102.log

	format := "%02d%02d%02d"
	if logging.fileNum > 0 {
		format += fmt.Sprintf(".%d", logging.fileNum)
	}

	name = fmt.Sprintf(format+".log",
		t.Year(),
		t.Month(),
		t.Day(),
	)

	return
}

var onceLogDirs sync.Once

// create creates a new log file if none exists and returns the file and its state
func create(t time.Time) (f *os.File, isFileExist bool, err error) {
	// onceLogDirs.Do(createLogDirs)
	// if len(logDirs) == 0 {
	// 	return nil, "", errors.New("log: no log dirs")
	// }
	if logging.logDir == "" {
		return nil, isFileExist, errors.New("log: no log dirs")
	}

	name := logName(t)
	fname := filepath.Join(logging.logDir, name)
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		if _, err := os.Stat(logging.logDir); os.IsNotExist(err) {
			if err := os.MkdirAll(logging.logDir, 0755); err != nil {
				return nil, isFileExist, err
			}
		}
	} else {
		isFileExist = true
	}

	f, err = os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, isFileExist, fmt.Errorf("log: cannot create log: %v", err)
	}

	return f, isFileExist, nil
}
