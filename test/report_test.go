package helper

import (
	"fmt"
	"testing"

	"github.com/J-Siu/go-helper"
)

func TestReportEmptyEmptyTrueTrue(t *testing.T) {
	var info string = ""
	var title string = ""
	var skip bool = true
	var single bool = true
	var want string = ""
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyEmptyFalseTrue(t *testing.T) {
	var info string = ""
	var title string = ""
	var skip bool = false
	var single bool = true
	var want string = ""
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyEmptyTrueFalse(t *testing.T) {
	var info string = ""
	var title string = ""
	var skip bool = true
	var single bool = false
	var want string = ""
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyEmptyFalseFalse(t *testing.T) {
	var info string = ""
	var title string = ""
	var skip bool = false
	var single bool = false
	var want string = ""
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportInfoEmptyTrueTrue(t *testing.T) {
	var info string = "This is info"
	var title string = ""
	var skip bool = true
	var single bool = true
	var want string = info + "\n"
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportInfoEmptyFalseTrue(t *testing.T) {
	var info string = "This is info"
	var title string = ""
	var skip bool = false
	var single bool = true
	var want string = info + "\n"
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportInfoEmptyTrueFalse(t *testing.T) {
	var info string = "This is info"
	var title string = ""
	var skip bool = true
	var single bool = false
	var want string = info + "\n"
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportInfoEmptyFalseFalse(t *testing.T) {
	var info string = "This is info"
	var title string = ""
	var skip bool = false
	var single bool = false
	var want string = info + "\n"
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyTitleTrueTrue(t *testing.T) {
	var info string = ""
	var title string = "title"
	var skip bool = true
	var single bool = true
	var want string = ""
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyTitleFalseTrue(t *testing.T) {
	var info string = ""
	var title string = "title"
	var skip bool = false
	var single bool = true
	var want string = title + ": \n"
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyTitleTrueFalse(t *testing.T) {
	var info string = ""
	var title string = "title"
	var skip bool = true
	var single bool = false
	var want string = ""
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyTitleFalseFalse(t *testing.T) {
	var info string = ""
	var title string = "title"
	var skip bool = false
	var single bool = false
	var want string = title + ": \n"
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}

func TestReportInfoTitleTrueTrue(t *testing.T) {
	var info string = "This is info"
	var title string = "title"
	var skip bool = true
	var single bool = true
	var want string = title + ": " + info + "\n"
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}

func TestReportInfoTitleFalseTrue(t *testing.T) {
	var info string = "This is info"
	var title string = "title"
	var skip bool = false
	var single bool = true
	var want string = title + ": " + info + "\n"
	msg := *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportInfoTitleTrueFalse(t *testing.T) {
	var info string = "This is info"
	var title string = "title"
	var skip bool = true
	var single bool = false
	var want string = title + ": \n" + info + "\n"
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportInfoTitleFalseFalse(t *testing.T) {
	var info string = "This is info"
	var title string = "title"
	var skip bool = false
	var single bool = false
	var want string = title + ": \n" + info + "\n"
	var msg string = *helper.ReportSp(info, title, skip, single)
	fmt.Print(msg)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
