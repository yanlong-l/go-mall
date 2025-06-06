package errcode

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
)

type AppError struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Cause    error  `json:"cause"`
	Occurred string `json:"occurred"` // 保存由底层错误导致AppErr发生时的位置

}

func (e *AppError) Error() string {
	if e == nil {
		return ""
	}
	formattedErr := struct {
		Code     int    `json:"code"`
		Msg      string `json:"msg"`
		Cause    string `json:"cause"`
		Occurred string `json:"occurred"`
	}{
		Code:     e.Code,
		Msg:      e.Msg,
		Occurred: e.Occurred,
	}
	if e.Cause != nil {
		formattedErr.Cause = e.Cause.Error()
	}
	errByte, _ := json.Marshal(formattedErr)
	return string(errByte)
}

func (e *AppError) String() string {
	return e.Error()
}

// WithCause 在逻辑执行中出现错误, 比如dao层返回的数据库查询错误
// 可以在领域层返回预定义的错误前附加上导致错误的基础错误。
// 如果业务模块预定义的错误码比较详细, 可以使用这个方法, 反之错误码定义的比较笼统建议使用Wrap方法包装底层错误生成项目自定义Error
// 并将其记录到日志后再使用预定义错误码返回接口响应
func (e *AppError) WithCause(err error) *AppError {
	e.Cause = err
	e.Occurred = getAppErrOccurredInfo()
	return e
}

// Wrap 用于逻辑中包装底层函数返回的error 和 WithCause 一样都是为了记录错误链条
// 该方法生成的error 用于日志记录, 返回响应请使用预定义好的error
func Wrap(msg string, err error) *AppError {
	if err == nil {
		return nil
	}
	appErr := &AppError{Code: -1, Msg: msg, Cause: err}
	appErr.Occurred = getAppErrOccurredInfo()
	return appErr
}

func getAppErrOccurredInfo() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	file = path.Base(file)
	funcName := runtime.FuncForPC(pc).Name()
	triggerInfo := fmt.Sprintf("func: %s, file: %s, line: %d", funcName, file, line)
	return triggerInfo
}

func newError(code int, msg string) *AppError {
	if code > -1 {
		if _, duplicated := codes[code]; duplicated {
			panic(fmt.Sprintf("错误码 %d 不能重复, 请检查后更换", code))
		}
		codes[code] = ""
	}
	return &AppError{Code: code, Msg: msg}
}
