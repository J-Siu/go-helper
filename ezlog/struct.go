/*
Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package ezlog

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/J-Siu/go-helper/v2/str"
	"github.com/J-Siu/go-helper/v2/strany"
)

type FuncOut func(msg *string)
type FuncDateTime func() string

type EzLog struct {
	StrAny         *strany.StrAny `json:"StrAny"`
	funcDateTime   FuncDateTime
	funcOut        FuncOut
	logLevel       EzLogLevel
	logLevelPrefix bool
	skipEmpty      bool
	strBuf         []string
	time           bool
	trim           bool // persistent trim
	// msg level
	msgLogLevel       EzLogLevel
	msgLogLevelPrefix bool
	msgNotEmpty       bool
	msgSkipEmpty      bool
	msgTrim           bool
}

func (t *EzLog) New() *EzLog {
	t.StrAny = new(strany.StrAny).New()
	return t.
		Clear().
		EnableTime(false).
		EnableTrim(true).
		SetDefaultDateTimeFunc().
		SetDefaultOutFun().
		SetLogLevel(ERR).
		SetLogLevelPrefix(true)
}

// Clear message
func (t *EzLog) Clear() *EzLog {
	t.msgLogLevelPrefix = false
	t.msgNotEmpty = false
	t.msgSkipEmpty = false
	t.msgTrim = false
	t.strBuf = nil
	return t
}

// Dump all control fields/variables. Use before Out().
func (t *EzLog) Dump(singleLine bool) *EzLog {
	indent := t.StrAny.GetIndentEnable()
	if singleLine {
		t.StrAny.IndentEnable(false)
		t.
			N("EzLog.Dump").Lm(t).
			N("logLevel").M(t.logLevel).
			N("logLevelPrefix").M(t.logLevelPrefix).
			N("skipEmpty").M(t.skipEmpty).
			N("time").M(t.time).
			N("trim").M(t.trim).
			N("msgLogLevel").M(t.msgLogLevel).
			N("msgNotEmpty").M(t.msgNotEmpty).
			N("msgSkipEmpty").M(t.msgSkipEmpty).
			N("msgTrim").M(t.msgTrim).
			N("t.msgLogLevel > DISABLED").M(t.msgLogLevel > DISABLED)
	} else {
		t.
			N("EzLog.Dump").Lm(t).
			Ln("logLevel").M(t.logLevel).
			Ln("logLevelPrefix").M(t.logLevelPrefix).
			Ln("skipEmpty").M(t.skipEmpty).
			Ln("time").M(t.time).
			Ln("trim").M(t.trim).
			Ln("msgLogLevel").M(t.msgLogLevel).
			Ln("msgNotEmpty").M(t.msgNotEmpty).
			Ln("msgSkipEmpty").M(t.msgSkipEmpty).
			Ln("msgTrim").M(t.msgTrim).
			Ln("t.msgLogLevel > DISABLED").M(t.msgLogLevel > DISABLED)
	}
	t.StrAny.IndentEnable(indent)
	return t
}

// Enable/Disable trim on `data`
func (t *EzLog) EnableSkipEmpty(enable bool) *EzLog {
	t.skipEmpty = enable
	return t
}

// Enable/Disable trim on `data`
func (t *EzLog) EnableTime(enable bool) *EzLog {
	t.time = enable
	return t
}

// Enable/Disable trim on `data`
func (t *EzLog) EnableTrim(enable bool) *EzLog {
	t.trim = enable
	return t
}

// Get log level
func (t *EzLog) GetLogLevel() EzLogLevel { return t.logLevel }

// Get log level prefix enable or not
func (t *EzLog) GetLogLevelPrefix() bool { return t.logLevelPrefix }

// Set log level
func (t *EzLog) SetLogLevel(level EzLogLevel) *EzLog {
	t.logLevel = level
	return t
}

// Set log level prefix true/false
func (t *EzLog) SetLogLevelPrefix(enable bool) *EzLog {
	t.logLevelPrefix = enable
	return t
}

// Set out function
func (t *EzLog) SetOutFunc(f FuncOut) *EzLog {
	t.funcOut = f
	return t
}

// Set out function to fmt.Println()
func (t *EzLog) SetDefaultOutFun() *EzLog {
	t.SetOutFunc(func(str *string) { fmt.Println(*str) })
	return t
}

// Set DateTime function
func (t *EzLog) SetDateTimeFunc(f FuncDateTime) *EzLog {
	t.funcDateTime = f
	return t
}

// Set default DateTime function
func (t *EzLog) SetDefaultDateTimeFunc() *EzLog {
	t.SetDateTimeFunc(func() string { return time.Now().Format("2006-01-02 15:04:05") })
	return t
}

// -- Set log message level

// Log message as `level`
func (t *EzLog) LogL(level EzLogLevel) *EzLog {
	t.msgLogLevel = level
	return t
}

// Log message without level (no level prefix)
func (t *EzLog) Log() *EzLog {
	t.msgLogLevel = LOG
	return t.Clear().Lp(false)
}

// Log message as `EMERG`
func (t *EzLog) Emerg() *EzLog {
	t.msgLogLevel = EMERG
	return t.Clear().Lp(t.logLevelPrefix)
}

// Log message as `ALERT`
func (t *EzLog) Alert() *EzLog {
	t.msgLogLevel = ALERT
	return t.Clear().Lp(t.logLevelPrefix)
}

// Log message as `CRIT`
func (t *EzLog) Crit() *EzLog {
	t.msgLogLevel = CRIT
	return t.Clear().Lp(t.logLevelPrefix)
}

// Log message as `ERR`
func (t *EzLog) Err() *EzLog {
	t.msgLogLevel = ERR
	return t.Clear().Lp(t.logLevelPrefix)
}

// Log message as `WARN`
func (t *EzLog) Warning() *EzLog {
	t.msgLogLevel = WARNING
	return t.Clear().Lp(t.logLevelPrefix)
}

// Log message as `NOTICE`
func (t *EzLog) Notice() *EzLog {
	t.msgLogLevel = NOTICE
	return t.Clear().Lp(t.logLevelPrefix)
}

// Log message as `INFO`
func (t *EzLog) Info() *EzLog {
	t.msgLogLevel = INFO
	return t.Clear().Lp(t.logLevelPrefix)
}

// Log message as `DEBUG`
func (t *EzLog) Debug() *EzLog {
	t.msgLogLevel = DEBUG
	return t.Clear().Lp(t.logLevelPrefix)
}

// Log message as `TRACE`
func (t *EzLog) Trace() *EzLog {
	t.msgLogLevel = TRACE
	return t.Clear().Lp(t.logLevelPrefix)
}

// --- Output

func (t *EzLog) Out() *EzLog {
	if t.msgLogLevel <= t.logLevel {
		// Skip empty?
		if !((t.skipEmpty || t.msgSkipEmpty) && !t.msgNotEmpty) {
			// DateTime
			if t.time {
				t.strBuf = append([]string{t.funcDateTime() + ":"}, t.strBuf...)
			}
			// Log level prefix
			if t.msgLogLevelPrefix && (t.msgLogLevel != DISABLED) {
				t.strBuf = append([]string{t.msgLogLevel.String() + ":"}, t.strBuf...)
			}
			t.funcOut(t.StringP())
		}
	}
	t.msgTrim = false
	return t
}

func (t *EzLog) String() string { return *t.StringP() }

func (t *EzLog) StringP() *string {
	var strOut string
	if t.msgLogLevel <= t.logLevel {
		if t.strBuf != nil {
			// str = strings.Join(l.strBuf, " ")
			for _, s := range t.strBuf {
				_, size := utf8.DecodeLastRuneInString(strOut)
				if size > 0 && strOut[len(strOut)-size] != '\n' {
					strOut += " "
				}
				strOut += s
			}
		}
	}
	return &strOut
}

// --- Build log message

// Add msg to log
func (t *EzLog) build(data any, isMsg bool) *EzLog {
	if t.msgLogLevel <= t.logLevel {
		tmp := *t.StrAny.Any(data)
		if t.trim || t.msgTrim {
			tmp = strings.Trim(tmp, "\n")
			tmp = strings.TrimSpace(tmp)
		}
		if isMsg {
			t.msgNotEmpty = t.msgNotEmpty || len(tmp) > 0
		}
		t.strBuf = append(t.strBuf, tmp)
	}
	return t
}

// --- Msg control

// enable/disable message log level prefix.
func (t *EzLog) Lp(enable bool) *EzLog {
	t.msgLogLevelPrefix = enable
	return t
}

// skip current message if `Msg` is empty. Current msg only.
func (t *EzLog) Se() *EzLog {
	t.msgSkipEmpty = true
	return t
}

// enable/disable trim data. Default to false. Reset each time Out() is called
func (t *EzLog) Tr(enable bool) *EzLog {
	t.msgTrim = enable
	return t
}

// enable/disable message log level prefix. (Alias for Lp())
func (t *EzLog) LogPrefix(enable bool) *EzLog { return t.Lp(enable) }

// skip current message if `Msg` is empty. Current msg only. (alias of Se())
func (t *EzLog) SkipEmpty() *EzLog { return t.Se() }

// enable/disable trim data. Default to false. Reset each time Out() is called. (alias of Trim())
func (t *EzLog) TrimData(enable bool) *EzLog { return t.Tr(enable) }

// --- base logging functions

// Append character/rune to message (shorthand for Sp())
func (t *EzLog) C(ch rune) *EzLog {
	if t.msgLogLevel <= t.logLevel {
		count := len(t.strBuf)
		if count == 0 {
			t.strBuf = append(t.strBuf, string(ch))
		} else {
			t.strBuf[count-1] += string(ch)
		}
	}
	return t
}

// Add newline to message.
func (t *EzLog) L() *EzLog { return t.C('\n') }

// Add msg to log
func (t *EzLog) M(data any) *EzLog { return t.build(data, true) }

// Add : after data
func (t *EzLog) N(data any) *EzLog { return t.build(data, false).C(':') }

// --- Shorthand

// Add new line before data (shorthand for L().M())
func (t *EzLog) Lm(data any) *EzLog { return t.L().M(data) }

// Add new line after data (shorthand for M().L())
func (t *EzLog) Ml(data any) *EzLog { return t.M(data).L() }

// Add new line before data and : after (shorthand for L().N())
func (t *EzLog) Ln(data any) *EzLog { return t.L().N(data) }

// Add : and new line after data (shorthand for N().L())
func (t *EzLog) Nl(data any) *EzLog { return t.N(data).L() }

// --- Expressive func name

// Add new line to log (alias of L())
func (t *EzLog) NewLine() *EzLog { return t.L() }

// Add msg to log (alias of M())
func (t *EzLog) Msg(data any) *EzLog { return t.M(data) }

// Add : after data (alias of N())
func (t *EzLog) Name(data any) *EzLog { return t.N(data) }

// --- Expressive alias

// Add new line before data (alias of L().M())
func (t *EzLog) NewLineMsg(data any) *EzLog { return t.Lm(data) }

// Add new line after data (alias of M().L())
func (t *EzLog) MsgNewLine(data any) *EzLog { return t.Ml(data) }

// Add new line before data and : after (alias of L().N())
func (t *EzLog) NewLineName(data any) *EzLog { return t.Ln(data) }

// Add : and new line after data (alias of N().L())
func (t *EzLog) NameNewLine(data any) *EzLog { return t.Nl(data) }

// --- Other shorthand

// Add tab to message.
func (t *EzLog) T() *EzLog { return t.C('\t') }

// Add tab to message. (alias for T())
func (t *EzLog) Tab() *EzLog { return t.T() }

// Add "End" to message. (shorthand for M("End"))
func (t *EzLog) TxtEnd() *EzLog { return t.M("End") }

// Add "Start" to message. (shorthand for M("Start"))
func (t *EzLog) TxtStart() *EzLog { return t.M("Start") }

// Add "OK"/"Fail" to message.
func (t *EzLog) Ok(data bool) *EzLog {
	return t.M(str.Ok(data))
}

// Add "Success"/"Fail" to message.
func (t *EzLog) Success(data bool) *EzLog {
	return t.M(str.Success(data))
}

// Add "Yes"/"No" to message.
func (t *EzLog) YesNo(data bool) *EzLog {
	return t.M(str.YesNo(data))
}
