# go-helper [![Paypal donate](https://www.paypalobjects.com/en_US/i/btn/btn_donate_LG.gif)](https://www.paypal.com/donate/?business=HZF49NM9D35SJ&no_recurring=0&currency_code=CAD)

Collections of Golang helper modules.

### Table Of Content
<!-- TOC -->

- [Features](#features)
- [Doc](#doc)
- [Usage](#usage)
- [Test](#test)
- [Used By Project](#used-by-project)
- [Repository](#repository)
- [Contributors](#contributors)
- [Change Log](#change-log)
- [License](#license)

<!-- /TOC -->
<!--more-->

### Features

|Module|Description|Readme|Example|
|---|---|---|---|
|array|simple array template|[README.md](array/README.md)||
|basestruct|a simple struct with 5 common fields to be embedded by other structs|[README.md](basestruct/README.md)|[example/basestruct](/example/basestruct/)|
|cmd|`exec.Command` shell wrapper|[README.md](cmd/README.md)||
|errs|simple error array for stacking error messages with prefix|[README.md](errs/README.md)||
|ezlog|log functions auto detect and apply json marshal indent|[README.md](ezlog/README.md)|[example/ezlog/](/example/ezlog/)|
|file|file/directory functions|[README.md](file/README.md)||
|strany|convert any to *string|[README.md](strany/README.md)|[example/strany/](/example/strany/)|
|str|string/array functions|[README.md](str/README.md)||
|ver|Both package and struct level functions for constructing version string|[README.md](ver/README.md)|[example/ver](/example/ver/)|

### Doc

- https://pkg.go.dev/github.com/J-Siu/go-helper

### Usage

```go
import "github.com/J-Siu/go-helper/v2"
```

### Test

```sh
cd test
# All
go test
# Individual
go test cmd_test.go
go test file_test.go
```

### Used By Project

- [go-gitapi](https://github.com/J-Siu/go-gitapi)
- [go-mygit](https://github.com/J-Siu/go-mygit)

### Repository

- [go-helper](https://github.com/J-Siu/go-helper)

### Contributors

- [John Sing Dao Siu](https://github.com/J-Siu)

### Change Log

- v1
  - 0.0.1
    - Initial Commit
  - 0.9.0
    - Function update
  - 0.9.1
    - Fix git command args
  - v0.9.1
    - Add prefix v for version number
  - v0.9.2
    - Fix MyCmdRunWg() missing wgP.Done()
  - v0.9.3
    - ReportTStringP():
      - fix using wrong var when handling *[]byte
      - add []byte case
  - v0.9.4
    - Report support SingleLine mode
  - v0.9.5
    - Fix ReportT.SpringP() output
  - v0.9.6
    - Fix ReportT.SpringP() skip empty output
  - v0.9.8
    - Fix ReportT.SpringP() logical bug
    - Add report_test.go
  - v0.9.9
    - Add workPath support for gitCmd and myCmd
    - Add GitRoot(), GitRootSubmodule(), GitExecExist(), GitExecPath()
    - Add test
  - v1.0.0
    - file.go
      - Add FullPath()
      - All func return *string
    - gitCmd.go
      - GitPush() correct optionP param type
    - myCmd.go
      - MyCmd.Run() improve debug output
      - MyCmdInit() use full path for WorkDir
    - report.go
      - ReportT.StringP()
        - Add *[]string case
        - case []string, *[]string
          - fix bug only print last line
          - no longer remove empty line
    - string.go
      - StrArrayPtrRemoveEmpty() return new array
      - StrPToArrayP no longer remove empty line
      - func name StrPToArrayP -> StrPtrToArrayPtr
  - v1.0.1
    - Improve comment for godoc
  - v1.1.1
    - Add Git(), GitBranchCurrent(), GitClone(), GitPull()
  - v1.1.2
    - Add IsRegularFile(), IsDir(), SameDir()
    - Add StrPtrToJsonIndentSp(), StrToJsonIndentSp(), AnyToJsonMarshalIndentSp()
    - Add basic error type
    - Add warning type
    - If !Debug, short circuit ReportDebug() ReportSpDebug()
    - Report() support error type
  - v1.1.3
    - Add ErrsType.Empty(), ErrsType.Clear()
    - Add FullPathStr(), FileRemoveExt(), FileInDir(), FileSimplifyName(), FileHasExt()
    - Add MyCmd.Reset(), MyCmd.Init()
    - Add number types/pointers support in ReportT.StringP()
    - Fix #6 - ReportT().StringP add space after title ":"
    - Fix BoxSealAnonymous() decoding length checking
  - v1.1.4
    - ReportT.StringP() handle nil *[]string
  - v1.1.5
    - Go 1.20
  - v1.1.6
    - MyCmd struct
      - Add ExitCode
      - Run() will handle exit code properly
      - Update test cases
  - v1.1.7
    - Use proper receiver var
  - v1.1.8
    - Add TildeEnvExpand()
  - v1.1.9
    - Add error, *error type support in ReportT.StringP()
  - v1.1.10
    - Added
      - ErrsQueue()
      - FileStrArrRead()
      - FileStrArrWrite()
      - NumToStr()
- v2
  - v2.0.0
    - Cleanup
  - v2.1.0
    - Add strany
  - v2.2.0
    - add ezlog, str.Any
  - v2.2.1
    - Remove array.go
    - cmd.go
      - Cmd struct is now exported
      - Add package level functions
    - str.Any
      - New() will enable indent
      - Add []error and *[]error support
  - v2.3.0
    - ezlog
      - improve json indent output support
      - add byte support
    - rename err->errs, as err is too commonly used for var
    - errs - add Clean()
  - v2.3.1
    - str - strIn use *string
  - v2.3.2
    - ezlog.Sp() takes rune instead of any
    - JsonIndent() will trim "\n"
  - v2.3.3
    - add array
    - strany - move to folder, to enable package level usage
    - ezlog
      - add trim support
      - add log level prefix
      - add shorthand L(), Ln(), N(), NLn(), N(), NLn(), S(). T(), Tab(), TxtEnd(), TxtStart()
      - breaking: change Sp() to accept `rune` (previous `any`)
  - v2.3.4
    - ezlog
      - add shorthand Mn(), Nn()
  - v2.4.0
    - add basestruct
    - StrAny add unquote support
  - v2.4.1
    - ezlog
      - fix ezlog.Sp() error
      - move example to /example folder
  - v2.4.2
    - ezlog
      - shorten log level name
      - expose EzLog
      - add LogL()
  - v2.4.3
    - add ver
    - add README for all packages
  - v2.5.0
    - add version const
    - split most packages into pkg, struct files
    - err use array.Array
    - strany
      - add debug output support for `Any`
      - fix output for `array.Array[any]`, `array.Array[error]`
  - v2.5.1
    - ezlog
      - add OK(), Success(), YesNo() shorthand
  - v2.5.2
    - errs
      - add `New`
    - ezlog
      - add `SetSkipEmpty`, `Se`/`SkipEmpty`, `Dump`, `LogPrefix`
      - add all struct functions to package level
      - remove `NLn`, `MLn`, `Ln`
      - rename
        - `Disabled` -> `DISABLED`
        - `Level` -> `ExLogLevel`
        - `Sp` -> `C`
    - file
      - `ArrayRead`, `ArrayWrite` use `*[]string` instead of `[]string`
    - str
      - add case sensitive support for `ArrayContains`, `ArrayContainsSubString`, `ContainsAnySubStrings`, `ContainsAnySubStringsBool`
      - add nil check
      - remove `ArrayPrintLn`

### License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
