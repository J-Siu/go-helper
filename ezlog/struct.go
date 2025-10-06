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

	"github.com/J-Siu/go-helper/v2/strany"
)

type OutFunc func(msg *string)

type EzLog struct {
	StrAny         *strany.StrAny `json:"str_any,omitempty"`
	logLevel       Level
	logLevelPrefix bool
	msgLogLevel    Level
	outFunc        OutFunc
	strBuf         []string
	trim           bool
}

func (t *EzLog) New() *EzLog {
	t.SetLogLevel(ERR)
	t.SetLogLevelPrefix(true)
	t.SetOutPrintLn()
	t.SetTrim(true)
	t.StrAny = new(strany.StrAny).New()
	return t
}

// Clear message
func (t *EzLog) Clear() *EzLog {
	t.strBuf = nil
	return t
}

// Get log level
func (t *EzLog) GetLogLevel() Level { return t.logLevel }

// Get log level prefix enable or not
func (t *EzLog) GetLogLevelPrefix() bool { return t.logLevelPrefix }

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

// Set log level
func (t *EzLog) SetLogLevel(level Level) *EzLog {
	t.logLevel = level
	return t
}

// Set log level prefix true/false
func (t *EzLog) SetLogLevelPrefix(enable bool) *EzLog {
	t.logLevelPrefix = enable
	return t
}

// Enable/Disable trim on `data`
func (t *EzLog) SetTrim(enable bool) *EzLog {
	t.trim = enable
	return t
}

// --- Output

func (t *EzLog) Out() *EzLog {
	if t.msgLogLevel <= t.logLevel {
		t.outFunc(t.StringP())
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

// Add newline to message. (shorthand for Ln())
func (t *EzLog) L() *EzLog { return t.Ln() }

// Add newline to message
func (t *EzLog) Ln() *EzLog { return t.Sp('\n') }

// Add msg to log
func (t *EzLog) M(date any) *EzLog { return t.Msg(date) }

// Add new line to message (shorthand for MsgLn())
func (t *EzLog) Mn(date any) *EzLog { return t.MsgLn(date) }

// Add new line to message (shorthand for MsgLn())
func (t *EzLog) MLn(date any) *EzLog { return t.MsgLn(date) }

// Add msg to log
func (t *EzLog) Msg(data any) *EzLog {
	if t.msgLogLevel <= t.logLevel {
		tmp := *t.StrAny.Any(data)
		if t.trim {
			tmp = strings.Trim(tmp, "\n")
			tmp = strings.TrimSpace(tmp)
		} else {
		}
		t.strBuf = append(t.strBuf, tmp)
	}
	return t
}

// Add new line to message (shorthand for Msg().Ln())
func (t *EzLog) MsgLn(data any) *EzLog { return t.Msg(data).Ln() }

// Add : after data (shorthand for Name())
func (t *EzLog) N(data any) *EzLog { return t.Name(data) }

// Add : and newline after data (shorthand for NameLn))
func (t *EzLog) Nn(data any) *EzLog { return t.NameLn(data) }

// Add : and newline after data (shorthand for NameLn))
func (t *EzLog) NLn(data any) *EzLog { return t.NameLn(data) }

// Add : after data (shorthand for Msg().Sp(':'))
func (t *EzLog) Name(data any) *EzLog { return t.Msg(data).Sp(':') }

// Add : and newline after data (shorthand for Msg().Sp(':').Ln())
func (t *EzLog) NameLn(data any) *EzLog { return t.Name(data).Ln() }

// Append character/rune to message (shorthand for Sp())
func (t *EzLog) S(ch rune) *EzLog { return t.Sp(ch) }

// Append character/rune to message
func (t *EzLog) Sp(ch rune) *EzLog {
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

// Add tab to message. (shorthand for Tab())
func (t *EzLog) T() *EzLog { return t.Tab() }

// Add tab to message
func (t *EzLog) Tab() *EzLog { return t.Sp('\t') }

// Add "End" to message. (shorthand for Msg("End"))
func (t *EzLog) TxtEnd() *EzLog { return t.Msg("End") }

// Add "Start" to message. (shorthand for Msg("Start"))
func (t *EzLog) TxtStart() *EzLog { return t.Msg("Start") }

// -- Set log message level

// Log message as `level`
func (t *EzLog) LogL(level Level) *EzLog {
	t.Clear().msgLogLevel = level
	return t
}

// Log message without level (no level prefix)
func (t *EzLog) Log() *EzLog {
	t.Clear().msgLogLevel = LogLevel
	return t
}

// Log message as `EMERG`
func (t *EzLog) Emerg() *EzLog {
	t.Clear().msgLogLevel = EMERG
	if t.logLevelPrefix {
		t.Name(t.msgLogLevel.String())
	}
	return t
}

// Log message as `ALERT`
func (t *EzLog) Alert() *EzLog {
	t.Clear().msgLogLevel = ALERT
	if t.logLevelPrefix {
		t.Name(t.msgLogLevel.String())
	}
	return t
}

// Log message as `CRIT`
func (t *EzLog) Crit() *EzLog {
	t.Clear().msgLogLevel = CRIT
	if t.logLevelPrefix {
		t.Name(t.msgLogLevel.String())
	}
	return t
}

// Log message as `ERR`
func (t *EzLog) Err() *EzLog {
	t.Clear().msgLogLevel = ERR
	if t.logLevelPrefix {
		t.Name(t.msgLogLevel.String())
	}
	return t
}

// Log message as `WARN`
func (t *EzLog) Warning() *EzLog {
	t.Clear().msgLogLevel = WARNING
	if t.logLevelPrefix {
		t.Name(t.msgLogLevel.String())
	}
	return t
}

// Log message as `NOTICE`
func (t *EzLog) Notice() *EzLog {
	t.Clear().msgLogLevel = NOTICE
	if t.logLevelPrefix {
		t.Name(t.msgLogLevel.String())
	}
	return t
}

// Log message as `INFO`
func (t *EzLog) Info() *EzLog {
	t.Clear().msgLogLevel = INFO
	if t.logLevelPrefix {
		t.Name(t.msgLogLevel.String())
	}
	return t
}

// Log message as `DEBUG`
func (t *EzLog) Debug() *EzLog {
	t.Clear().msgLogLevel = DEBUG
	if t.logLevelPrefix {
		t.Name(t.msgLogLevel.String())
	}
	return t
}

// Log message as `TRACE`
func (t *EzLog) Trace() *EzLog {
	t.Clear().msgLogLevel = TRACE
	if t.logLevelPrefix {
		t.Name(t.msgLogLevel.String())
	}
	return t
}
