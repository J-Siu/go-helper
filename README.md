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

### License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
