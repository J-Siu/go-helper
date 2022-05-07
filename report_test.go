package helper

import "testing"

func TestReportEmptyEmptytruetrue(t *testing.T) {
	var info string = ""
	var title string = ""
	var skip bool = true
	var single bool = true
	var want string = ""
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyEmptyfalsetrue(t *testing.T) {
	var info string = ""
	var title string = ""
	var skip bool = false
	var single bool = true
	var want string = ""
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyEmptytruefalse(t *testing.T) {
	var info string = ""
	var title string = ""
	var skip bool = true
	var single bool = false
	var want string = ""
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyEmptyfalsefalse(t *testing.T) {
	var info string = ""
	var title string = ""
	var skip bool = false
	var single bool = false
	var want string = ""
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportInfoEmptytruetrue(t *testing.T) {
	var info string = "This is info"
	var title string = ""
	var skip bool = true
	var single bool = true
	var want string = info + "\n"
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportInfoEmptyfalsetrue(t *testing.T) {
	var info string = "This is info"
	var title string = ""
	var skip bool = false
	var single bool = true
	var want string = info + "\n"
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportInfoEmptytruefalse(t *testing.T) {
	var info string = "This is info"
	var title string = ""
	var skip bool = true
	var single bool = false
	var want string = info + "\n"
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportInfoEmptyfalsefalse(t *testing.T) {
	var info string = "This is info"
	var title string = ""
	var skip bool = false
	var single bool = false
	var want string = info + "\n"
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyTitletruetrue(t *testing.T) {
	var info string = ""
	var title string = "title"
	var skip bool = true
	var single bool = true
	var want string = ""
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyTitlefalsetrue(t *testing.T) {
	var info string = ""
	var title string = "title"
	var skip bool = false
	var single bool = true
	var want string = title + ":\n"
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyTitletruefalse(t *testing.T) {
	var info string = ""
	var title string = "title"
	var skip bool = true
	var single bool = false
	var want string = ""
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportEmptyTitlefalsefalse(t *testing.T) {
	var info string = ""
	var title string = "title"
	var skip bool = false
	var single bool = false
	var want string = title + ":\n"
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}

func TestReportInfoTitletruetrue(t *testing.T) {
	var info string = "This is info"
	var title string = "title"
	var skip bool = true
	var single bool = true
	var want string = title + ":" + info + "\n"
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}

func TestReportInfoTitlefalsetrue(t *testing.T) {
	var info string = "This is info"
	var title string = "title"
	var skip bool = false
	var single bool = true
	var want string = title + ":" + info + "\n"
	msg := *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportInfoTitletruefalse(t *testing.T) {
	var info string = "This is info"
	var title string = "title"
	var skip bool = true
	var single bool = false
	var want string = title + ":\n" + info + "\n"
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
func TestReportInfoTitlefalsefalse(t *testing.T) {
	var info string = "This is info"
	var title string = "title"
	var skip bool = false
	var single bool = false
	var want string = title + ":\n" + info + "\n"
	var msg string = *ReportSp(info, title, skip, single)
	if msg != want {
		t.Fatalf(`ReportSp(%s, %s, %t, %t) = "%s"`, info, title, skip, single, msg)
	}
}
