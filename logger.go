// Copyright 2015 Nevio Vesic
// Please check out LICENSE file for more information about what you CAN and what you CANNOT do!
// Basically in short this is a free software for you to do whatever you want to do BUT copyright must be included!
// I didn't write all of this code so you could say it's yours.
// MIT License

package goesl

import (
	"github.com/sirupsen/logrus"
)

var (
	// Log is the pointer for a logger engine
	Log *logrus.Logger
)

// Debug to log
func Debug(message string, args ...interface{}) {
	Log.Debugf(message, args...)
}

// Error to log
func Error(message string, args ...interface{}) {
	Log.Errorf(message, args...)
}

// Notice is alias to info
func Notice(message string, args ...interface{}) {
	Log.Infof(message, args...)
}

// Info to log
func Info(message string, args ...interface{}) {
	Log.Infof(message, args...)
}

// Warning to log
func Warning(message string, args ...interface{}) {
	Log.Warningf(message, args...)
}
