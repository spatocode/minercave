package utils

import (
	"github.com/fatih/color"
)


var (
	LOG_ERR = color.New(color.FgRed, color.Bold).PrintfFunc()
	LOG_INFO = color.New(color.FgBlue, color.Bold).PrintfFunc()
	LOG_WARN = color.New(color.FgYellow, color.Bold).PrintfFunc()
	LOG_SUCCESS = color.New(color.FgGreen, color.Bold).PrintfFunc()
)
