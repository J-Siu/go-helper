/*
Copyright © 2022 John, Sing Dao, Siu <john.sd.siu@gmail.com>

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

package helper

import (
	"fmt"
)

// ReportT is the base structure for the Report.StringP() function, which support printing a wide range of data types.
type ReportT struct {
	Data       any    `json:"Data"`       // Data to be printed
	Title      string `json:"Title"`      // Title of print out
	ModeStatus bool   `json:"ModeStatus"` // bool to "OK/Failed"
	SkipEmpty  bool   `json:"SkipEmpty"`  // Return empty string if Data is empty
	SingleLine bool   `json:"SingleLine"` // No need line after title
}

// Print any data(optional) with title(optional).
// If <skipEmpty> is true, will not print title if <data> is empty.
// If <singleLine> is true, <data> will not start on new line.
func Report(data any, title string, skipEmpty bool, singleLine bool) {
	fmt.Print(*ReportNew(data, title, skipEmpty, singleLine).StringP())
}

// Only print if helper.Debug is true.
// Print any data(optional) with title(optional).
// If <skipEmpty> is true, will not print title if <data> is empty.
// If <singleLine> is true, <data> will not start on new line.
func ReportDebug(data any, title string, skipEmpty bool, singleLine bool) {
	if Debug {
		fmt.Print(*ReportNew(data, title, skipEmpty, singleLine).StringP())
	}
}

// Print any data(optional) with title(optional) into a string.
// If <skipEmpty> is true, will not print title if <data> is empty.
// If <singleLine> is true, <data> will not start on new line.
// Return a string pointer.
func ReportSp(data any, title string, skipEmpty bool, singleLine bool) *string {
	return ReportNew(data, title, skipEmpty, singleLine).StringP()
}

// Only print if helper.Debug is true.
// Print any data(optional) with title(optional) into a string.
// If <skipEmpty> is true, will not print title if <data> is empty.
// If <singleLine> is true, <data> will not start on new line.
// Return a string pointer.
func ReportSpDebug(data any, title string, skipEmpty bool, singleLine bool) *string {
	if Debug {
		return ReportNew(data, title, skipEmpty, singleLine).StringP()
	} else {
		var s string = ""
		return &s
	}
}

// Print bool into true/false, with title(optional).
// If <skipEmpty> is true, will not print title if <data> is empty.
// If <singleLine> is true, <data> will not start on new line.
func ReportStatus(data bool, title string, singleLine bool) {
	r := ReportNew(data, title, false, singleLine)
	r.ModeStatus = true
	fmt.Print(r)
}

// Only print if helper.Debug is true.
// Print bool into true/false, with title(optional), into a string.
// If <skipEmpty> is true, will not print title if <data> is empty.
// If <singleLine> is true, <data> will not start on new line.
// Return a string pointer.
func ReportStatusSp(data bool, title string, singleLine bool) *string {
	r := ReportNew(data, title, false, singleLine)
	r.ModeStatus = true
	return r.StringP()
}

// Setup ReportT with data(optional/nil), title(optional/""), <skipEmpty>, <singleLine>.
// Return the ReportT pointer.
func ReportNew(data any, title string, skipEmpty bool, singleLine bool) *ReportT {
	var r ReportT
	r.Data = data
	r.Title = title
	r.SingleLine = singleLine
	r.SkipEmpty = skipEmpty
	return &r
}

// Print self.Data, self.Title to string
// If self.SkipEmpty is true, will not print self.Title if self.Data is empty.
// If self.SingleLine is true, self.Data will not start on new line.
func (self *ReportT) String() string {
	return *self.StringP()
}

// Only print if helper.Debug is true.
// Print self.Data, self.Title to string
// If self.SkipEmpty is true, will not print self.Title if self.Data is empty.
// If self.SingleLine is true, self.Data will not start on new line.
func (self *ReportT) StringDebug() string {
	return *self.StringPDebug()
}

// Print self.Data, self.Title to string pointer
// If self.SkipEmpty is true, will not print self.Title if self.Data is empty.
// If self.SingleLine is true, self.Data will not start on new line.
func (self *ReportT) StringP() *string {
	var output string

	switch v := self.Data.(type) {
	case Err:
		if DebugReport {
			fmt.Println("case Err")
		}
		output = *StrToJsonIndentSp(v.Error(), true)
	case ErrsT:
		if DebugReport {
			fmt.Println("case ErrsT")
		}
		for _, e := range v {
			output += *StrToJsonIndentSp(e.Error(), true)
		}
	case error:
		if DebugReport {
			fmt.Println("case error")
		}
		output = *StrToJsonIndentSp(v.Error(), true)
	case []error:
		if DebugReport {
			fmt.Println("case []error")
		}
		for _, e := range v {
			output += *StrToJsonIndentSp(e.Error(), true)
		}
	case Warning:
		if DebugReport {
			fmt.Println("case Warning")
		}
		output = *StrPtrToJsonIndentSp(v.StringP(), true)
	case Warnings:
		if DebugReport {
			fmt.Println("case Warnings")
		}
		for _, w := range v {
			output = *StrPtrToJsonIndentSp(w.StringP(), true)
		}
	case *Warning:
		if DebugReport {
			fmt.Println("case *Warning")
		}
		output = *StrPtrToJsonIndentSp(v.StringP(), true)
	case *Warnings:
		if DebugReport {
			fmt.Println("case *Warnings")
		}
		for _, w := range *v {
			output = *StrPtrToJsonIndentSp(w.StringP(), true)
		}
	case string:
		if DebugReport {
			fmt.Println("case string")
		}
		if len(v) > 0 {
			output = *StrPtrToJsonIndentSp(&v, true)
		}
	case *string:
		if DebugReport {
			fmt.Println("case *string")
		}
		if v != nil && len(*v) > 0 {
			output += *StrPtrToJsonIndentSp(v, true)
		}
	case []string:
		if DebugReport {
			fmt.Println("case []string")
		}
		if len(v) > 0 {
			for _, s := range v {
				output += *StrPtrToJsonIndentSp(&s, true)
			}
		}
	case *[]string:
		if DebugReport {
			fmt.Println("case *[]string")
		}
		if len(*v) > 0 {
			for _, s := range *v {
				output += *StrPtrToJsonIndentSp(&s, true)
			}
		}
	case []byte:
		if DebugReport {
			fmt.Println("case []byte")
		}
		output = *JsonIndentSp(&v, true)
	case *[]byte:
		if DebugReport {
			fmt.Println("case *[]byte")
		}
		output = *JsonIndentSp(v, true)
	case bool:
		if DebugReport {
			fmt.Println("case bool")
		}
		if self.ModeStatus {
			output = BoolStatus(v) + "\n"
		} else {
			output = BoolString(v) + "\n"
		}
	case *bool:
		if DebugReport {
			fmt.Println("case *bool")
		}
		if self.ModeStatus {
			output = BoolStatus(*v) + "\n"
		} else {
			output = BoolString(*v) + "\n"
		}
	default:
		if DebugReport {
			fmt.Println("case default")
		}
		output = *AnyToJsonIndentSp(self.Data, true)
	}

	// Title
	var title string
	if len(self.Title) > 0 {
		title = self.Title + ":"
	}

	// Output
	if len(title) > 0 {
		if !self.SkipEmpty && (len(output) == 0 || !self.SingleLine) ||
			!self.SingleLine && len(output) != 0 {
			output = title + "\n" + output
		} else if len(output) > 0 {
			output = title + output
		}
	}

	return &output
}

// Only print if helper.Debug is true.
// Print self.Data, self.Title to string pointer
// If self.SkipEmpty is true, will not print self.Title if self.Data is empty.
// If self.SingleLine is true, self.Data will not start on new line.
func (self *ReportT) StringPDebug() *string {
	if Debug {
		return self.StringP()
	}
	var s string = ""
	return &s
}
