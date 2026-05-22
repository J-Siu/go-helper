module github.com/J-Siu/go-helper/v2

go 1.26.3

require (
	github.com/charlievieth/strcase v0.0.5
	github.com/edwardrf/symwalk v0.1.0
)

retract (
	v2.8.3 // Published accidentally
)

require golang.org/x/sys v0.45.0 // indirect
