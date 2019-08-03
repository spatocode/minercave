package utils

import (
	"github.com/fatih/color"
)


var (
	LOG_ERR = color.New(color.FgRed).PrintfFunc()
	LOG_INFO = color.New(color.FgBlue).PrintfFunc()
	LOG_WARN = color.New(color.FgYellow).PrintfFunc()
	LOG_SUCCESS = color.New(color.FgGreen).PrintfFunc()
)
