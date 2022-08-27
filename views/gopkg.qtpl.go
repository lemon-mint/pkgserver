// Code generated by qtc from "gopkg.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/gopkg.qtpl:1
package views

//line views/gopkg.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/gopkg.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/gopkg.qtpl:1
func StreamGoPackage(qw422016 *qt422016.Writer, pkg, vcs, url string) {
//line views/gopkg.qtpl:1
	qw422016.N().S(`
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="`)
//line views/gopkg.qtpl:6
	qw422016.E().S(pkg)
//line views/gopkg.qtpl:6
	qw422016.N().S(` `)
//line views/gopkg.qtpl:6
	qw422016.E().S(vcs)
//line views/gopkg.qtpl:6
	qw422016.N().S(` `)
//line views/gopkg.qtpl:6
	qw422016.N().S(url)
//line views/gopkg.qtpl:6
	qw422016.N().S(`">
<meta http-equiv="refresh" content="0; url=https://pkg.go.dev/`)
//line views/gopkg.qtpl:7
	qw422016.E().S(pkg)
//line views/gopkg.qtpl:7
	qw422016.N().S(`">
</head>
<body>
<a href="https://pkg.go.dev/`)
//line views/gopkg.qtpl:10
	qw422016.E().S(pkg)
//line views/gopkg.qtpl:10
	qw422016.N().S(`">Redirecting to documentation...</a>
</body>
</html>
`)
//line views/gopkg.qtpl:13
}

//line views/gopkg.qtpl:13
func WriteGoPackage(qq422016 qtio422016.Writer, pkg, vcs, url string) {
//line views/gopkg.qtpl:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/gopkg.qtpl:13
	StreamGoPackage(qw422016, pkg, vcs, url)
//line views/gopkg.qtpl:13
	qt422016.ReleaseWriter(qw422016)
//line views/gopkg.qtpl:13
}

//line views/gopkg.qtpl:13
func GoPackage(pkg, vcs, url string) string {
//line views/gopkg.qtpl:13
	qb422016 := qt422016.AcquireByteBuffer()
//line views/gopkg.qtpl:13
	WriteGoPackage(qb422016, pkg, vcs, url)
//line views/gopkg.qtpl:13
	qs422016 := string(qb422016.B)
//line views/gopkg.qtpl:13
	qt422016.ReleaseByteBuffer(qb422016)
//line views/gopkg.qtpl:13
	return qs422016
//line views/gopkg.qtpl:13
}
