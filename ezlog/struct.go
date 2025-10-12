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
	"unicode/utf8"

	"github.com/J-Siu/go-helper/v2/str"
	"github.com/J-Siu/go-helper/v2/strany"
)

type OutFunc func(msg *string)

type EzLog struct {
	StrAny         *strany.StrAny `json:"StrAny"`
	logLevel       EzLogLevel
	logLevelPrefix bool
	outFunc        OutFunc
	skipEmpty      bool
	strBuf         []string
	trim           bool
	// msg level
	msgLogLevel       EzLogLevel
	msgNotEmpty       bool
	msgSkipEmpty      bool
	msgLogLevelPrefix bool
}

func (t *EzLog) New() *EzLog {
	t.SetLogLevel(ERR)
	t.SetLogLevelPrefix(true)
	t.SetOutPrintLn()
	t.SetTrim(true)
	t.StrAny = new(strany.StrAny).New()
	t.msgNotEmpty = false
	t.msgSkipEmpty = false
	t.strBuf = nil
	return t
}

// Clear message
func (t *EzLog) Clear() *EzLog {
	t.msgNotEmpty = false
	t.strBuf = nil
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
func (t *EzLog) SetOutFunc(f OutFunc) *EzLog {
	t.outFunc = f
	return t
}

// Set out function to fmt.Print()
func (t *EzLog) SetOutPrint() *EzLog {
	t.SetOutFunc(func(str *string) { fmt.Print(*str) })
	return t
}

// Set out function to fmt.Println()
func (t *EzLog) SetOutPrintLn() *EzLog {
	t.SetOutFunc(func(str *string) { fmt.Println(*str) })
	return t
}

// Enable/Disable trim on `data`
func (t *EzLog) SetTrim(enable bool) *EzLog {
	t.trim = enable
	return t
}

// Enable/Disable trim on `data`
func (t *EzLog) SetSkipEmpty(enable bool) *EzLog {
	t.skipEmpty = enable
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

func (t *EzLog) Dump() *EzLog {
	new(EzLog).New().Log().
		Nn("EzLog.Dump").M(t).
		N("logLevel").Mn(t.logLevel).
		N("logLevelPrefix").Mn(t.logLevelPrefix).
		N("skipEmpty").Mn(t.skipEmpty).
		N("trim").Mn(t.trim).
		N("msgLogLevel").Mn(t.msgLogLevel).
		N("msgNotEmpty").Mn(t.msgNotEmpty).
		N("msgSkipEmpty").Mn(t.msgSkipEmpty).
		N("t.msgLogLevel > DISABLED").Mn(t.msgLogLevel > DISABLED).
		Out()
	return t
}

func (t *EzLog) Out() *EzLog {
	if t.msgLogLevel <= t.logLevel {
		// Skip empty?
		if !((t.skipEmpty || t.msgSkipEmpty) && !t.msgNotEmpty) {
			// Log level prefix
			if t.msgLogLevelPrefix && (t.msgLogLevel != DISABLED) {
				t.strBuf = append([]string{t.msgLogLevel.String() + ":"}, t.strBuf...)
			}
			t.outFunc(t.StringP())
		}
	}
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
		if t.trim {
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

// enable/disable message log level prefix.
func (t *EzLog) Lp(enable bool) *EzLog {
	t.msgLogLevelPrefix = enable
	return t
}

// enable/disable message log level prefix. (Alias for Lp())
func (t *EzLog) LogPrefix(enable bool) *EzLog { return t.Lp(enable) }

// Skip current message if `Msg` is empty. Current msg only.
func (t *EzLog) Se() *EzLog {
	t.msgSkipEmpty = true
	return t
}

// Skip current message if `Msg` is empty. Current msg only. (alias of Se())
func (t *EzLog) SkipEmpty() *EzLog { return t.Se() }

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

// Add new line to message (shorthand for M().L())
func (t *EzLog) Mn(data any) *EzLog { return t.M(data).L() }

// Add msg to log (alias for M())
func (t *EzLog) Msg(data any) *EzLog { return t.M(data) }

// Add new line to message (alias for Mn())
func (t *EzLog) MsgLn(data any) *EzLog { return t.Mn(data) }

// Add : after data
func (t *EzLog) N(data any) *EzLog { return t.build(data, false).C(':') }

// Add : and newline after data (shorthand for N().L())
func (t *EzLog) Nn(data any) *EzLog { return t.N(data).L() }

// Add : after data (alias of N())
func (t *EzLog) Name(data any) *EzLog { return t.N(data) }

// Add : and newline after data ()
func (t *EzLog) NameLn(data any) *EzLog { return t.Nn(data) }

// -- Other shorthand

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
