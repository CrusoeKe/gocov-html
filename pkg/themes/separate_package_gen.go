// Code generated by "go run generator.go". DO NOT EDIT.

// TODO kehuishu001
package themes

import (
	"text/template"
	"time"
)

func (t separatePackageTheme) Data() *templateData {
	td := &templateData{
		When:       time.Now().Format(time.RFC1123),
		ProjectURL: ProjectURL,
	}

	td.Style = "Ym9keSB7CiAgICBiYWNrZ3JvdW5kLWNvbG9yOiAjZmZmOwogICAgZm9udC1mYW1pbHk6ICJIZWx2ZXRpY2EgTmV1ZSIsIEhlbHZldGljYSwgQXJpYWwsIHNhbnMtc2VyaWY7Cn0KCnRhYmxlIHsKICAgIG1hcmdpbi1sZWZ0OiAxMHB4OwogICAgYm9yZGVyLWNvbGxhcHNlOiBjb2xsYXBzZTsKfQoKdGQgewogICAgYmFja2dyb3VuZC1jb2xvcjogI2ZmZjsKICAgIHBhZGRpbmc6IDJweDsKfQoKdGFibGUub3ZlcnZpZXcgdGQgewogICAgcGFkZGluZy1yaWdodDogMjBweDsKfQoKdGQucGVyY2VudCwKdGQubGluZWNvdW50IHsKICAgIHRleHQtYWxpZ246IHJpZ2h0Owp9CgpkaXYucGFja2FnZSwKI3RvdGFsY292IHsKICAgIGNvbG9yOiAjZmZmOwogICAgYmFja2dyb3VuZC1jb2xvcjogIzM3NWVhYjsKICAgIGZvbnQtc2l6ZTogMTZweDsKICAgIGZvbnQtd2VpZ2h0OiBib2xkOwogICAgcGFkZGluZzogMTBweDsKICAgIGJvcmRlci1yYWRpdXM6IDVweCA1cHggNXB4IDVweDsKfQoKZGl2LnBhY2thZ2UsCiN0b3RhbGNvdiB7CiAgICBmbG9hdDogcmlnaHQ7CiAgICByaWdodDogMTBweDsKfQoKI3RvdGFsY292IHsKICAgIHRvcDogMTBweDsKICAgIHBvc2l0aW9uOiByZWxhdGl2ZTsKICAgIGJhY2tncm91bmQtY29sb3I6ICNmZmY7CiAgICBjb2xvcjogIzAwMDsKICAgIGJvcmRlcjogMXB4IHNvbGlkICMzNzVlYWI7CiAgICBjbGVhcjogYm90aDsKfQoKI3N1bW1hcnlXcmFwcGVyIHsKICAgIHBvc2l0aW9uOiBmaXhlZDsKICAgIHRvcDogMTBweDsKICAgIGZsb2F0OiByaWdodDsKICAgIHJpZ2h0OiAxMHB4OwoKfQoKc3Bhbi5wYWNrYWdlVG90YWwgewogICAgZmxvYXQ6IHJpZ2h0OwogICAgY29sb3I6ICMwMDA7Cn0KCiNkb2N0aXRsZSB7CiAgICBiYWNrZ3JvdW5kLWNvbG9yOiAjZmZmOwogICAgZm9udC1zaXplOiAyNHB4OwogICAgbWFyZ2luLXRvcDogMjBweDsKICAgIG1hcmdpbi1sZWZ0OiAxMHB4OwogICAgY29sb3I6ICMzNzVlYWI7CiAgICBmb250LXdlaWdodDogYm9sZDsKfQoKI2Fib3V0IHsKICAgIG1hcmdpbi1sZWZ0OiAxOHB4OwogICAgZm9udC1zaXplOiAxMHB4Owp9CgouZnVuY3RpdGxlLAouZnVuY25hbWUgewogICAgdGV4dC1hbGlnbjogY2VudGVyOwogICAgZm9udC1zaXplOiAyMHB4OwogICAgZm9udC13ZWlnaHQ6IGJvbGQ7CiAgICBjb2xvcjogIzM3NWVhYjsKfQoKLmZ1bmNuYW1lIHsKICAgIHRleHQtYWxpZ246IGxlZnQ7CiAgICBtYXJnaW4tdG9wOiAyMHB4OwogICAgbWFyZ2luLWxlZnQ6IDEwcHg7CiAgICBtYXJnaW4tYm90dG9tOiAyMHB4OwogICAgcGFkZGluZzogMnB4IDVweCA1cHg7CiAgICBiYWNrZ3JvdW5kLWNvbG9yOiAjZTBlYmY1Owp9Cgp0YWJsZS5saXN0aW5nIHsKICAgIG1hcmdpbi1sZWZ0OiAxMHB4Owp9Cgp0YWJsZS5saXN0aW5nIHRkIHsKICAgIHBhZGRpbmc6IDBweDsKICAgIGZvbnQtc2l6ZTogMTJweDsKICAgIGJhY2tncm91bmQtY29sb3I6ICNlZWU7CiAgICB2ZXJ0aWNhbC1hbGlnbjogdG9wOwogICAgcGFkZGluZy1sZWZ0OiAxMHB4OwogICAgYm9yZGVyLWJvdHRvbTogMXB4IHNvbGlkICNmZmY7Cn0KCnRhYmxlLmxpc3RpbmcgdGQ6Zmlyc3QtY2hpbGQgewogICAgdGV4dC1hbGlnbjogcmlnaHQ7CiAgICBmb250LXdlaWdodDogYm9sZDsKICAgIHZlcnRpY2FsLWFsaWduOiBjZW50ZXI7Cn0KCnRhYmxlLmxpc3RpbmcgdHIubWlzcyB0ZCB7CiAgICBiYWNrZ3JvdW5kLWNvbG9yOiAjRkZCQkI4Owp9Cgp0YWJsZS5saXN0aW5nIHRyOmxhc3QtY2hpbGQgdGQgewogICAgZm9udC13ZWlnaHQ6IG5vcm1hbDsKICAgIGNvbG9yOiAjMDAwOwp9Cgp0YWJsZS5saXN0aW5nIHRyOmxhc3QtY2hpbGQgdGQ6Zmlyc3QtY2hpbGQgewogICAgZm9udC13ZWlnaHQ6IGJvbGQ7Cn0KCi5pbmZvIHsKICAgIG1hcmdpbi1sZWZ0OiAxMHB4Owp9CgouaW5mbyBjb2RlIHt9CgpwcmUgewogICAgbWFyZ2luOiAxcHg7Cn0KCnByZS5jbWQgewogICAgYmFja2dyb3VuZC1jb2xvcjogI2U5ZTllOTsKICAgIGJvcmRlci1yYWRpdXM6IDVweCA1cHggNXB4IDVweDsKICAgIHBhZGRpbmc6IDEwcHg7CiAgICBtYXJnaW46IDIwcHg7CiAgICBsaW5lLWhlaWdodDogMThweDsKICAgIGZvbnQtc2l6ZTogMTRweDsKfQoKYSB7CiAgICB0ZXh0LWRlY29yYXRpb246IG5vbmU7CiAgICBjb2xvcjogIzM3NWVhYjsKfQoKYTpob3ZlciB7CiAgICB0ZXh0LWRlY29yYXRpb246IHVuZGVybGluZTsKfQoKcCB7CiAgICBtYXJnaW4tbGVmdDogMTBweDsKfQ=="

	return td
}

func (t separatePackageTheme) Template() *template.Template {
	tmpl := `{{define "theme"}}
<html>
	<head>
		<title>Coverage Report</title>
		<meta charset="utf-8" />
        {{if .Style}}
        <style type="text/css">
        {{.Style}}
        </style>
        {{end}}
	</head>
	<body>
		<div id="doctitle">Coverage Report</div>
        {{if not .Packages}}
		<p>no test files in package.</p>"
        {{else}}
        <div id="about">Generated on {{.When}} with <a href="{{.ProjectURL}}">gocov-html</a></div>
        {{/* Report overview/summary available? */}}
        {{if .Overview}}
        <div class="funcname">Report Overview</div>
            <table class="overview">
            {{range $k,$rp := .Packages}}
            <tr id="s_pkg_{{$rp.Pkg.Name}}">
                <td><code><a href="{{$rp.HtmlFilePath}}">{{$rp.Pkg.Name}}</a></code></td>
                <td class="percent"><code>{{printf "%.1f%%" $rp.PercentageReached}}</code></td>
                <td class="linecount"><code>{{printf "%d" $rp.ReachedStatements}}/{{printf "%d" $rp.TotalStatements}}</code></td>
            </tr>
            {{end}}
            </table>
            <p>
            This coverage report has been generated with the following command:
            </p>
            <pre class="cmd">{{.Command}}</pre>
        </div>
        {{end}}

        <div id="summaryWrapper">
        {{if not .Overview}}
            {{$rp := index .Packages 0}}
            <div class="package">{{$rp.Pkg.Name}}</div>
            <div id="totalcov">{{printf "%.1f%%" $rp.PercentageReached}}</div>
        {{else}}
            <div class="package">{{.Overview.Pkg.Name}}</div>
            <div id="totalcov">{{printf "%.1f%%" .Overview.PercentageReached}}</div>
        {{end}} {{/* if overview end */}}
        </div>
        {{end}} {{/* range if end */}}
        {{if .Script}}
        <script type="text/javascript">
        {{.Script}}
        </script>
        {{end}}
	</body>
</html>
{{end}}`
	p := template.Must(template.New("theme").Parse(tmpl))
	return p
}

func (t separatePackageTheme) PackageTemplate() *template.Template {
	tmpl := `{{define "theme"}}
<html>
	<head>
		<title>Coverage Report</title>
		<meta charset="utf-8" />
        {{if .Style}}
        <style type="text/css">
        {{.Style}}
        </style>
        {{end}}
	</head>
	<body>
		<div id="doctitle">Package Coverage Report</div>
        {{if not .Pkg}}
		<p>no test files in package.</p>
        {{end}}

        <div id="pkg_{{.Pkg.Name}}" class="funcname">
            Package Overview: {{.Pkg.Name}}
            <span class="packageTotal">{{printf "%.1f%%" .ReachedPercentage}}</span>
        </div>
        <p>Please select a function to see what's left for testing.</p>

        <table class="overview">
        {{range $k,$f := .Functions}}
            <tr id="s_fn_{{$f.Name}}">
                <td>
                    <code><a href="#fn_{{$f.Name}}">{{$f.Name}}(...)</a></code>
                </td>
                <td>
                    <code>{{$.Pkg.Name}}/{{$f.ShortFileName}}</code>
                </td>
                <td class="percent">
                    <code>{{printf "%.1f%%" $f.CoveragePercent}}</code>
                </td>
                <td class="linecount">
                    <code>{{$f.StatementsReached}}/{{len $f.Statements}}</code>
                </td>
            </tr>
        {{end}}
        </table>

        {{/* Functions source code here */}}
        {{range $k,$f := .Functions}}
        <div class="funcname" id="fn_{{$f.Name}}">func {{$f.Name}}</div>
        <div class="info">
            <a href="#s_fn_{{$f.Name}}">Back</a>
            <p>In <code>{{$f.File}}</code>:</p>
        </div>
        <table class="listing">
            {{range $p,$info := $f.Lines}}
            <tr{{if $info.Missed}} class="miss"{{end}}>
                <td>{{$info.LineNumber}}</td>
                <td>
                    <code><pre>{{$info.Code}}</pre></code>
                </td>
            </tr>
            {{end}}
        </table>
        {{end}} {{/* range function lines */}}

        <!--    Can be parsed by external script
                PACKAGE:{{.Pkg.Name}} DONE:{{printf "%.1f" .ReachedPercentage}}
        -->
	</body>
</html>
{{end}}`
	p := template.Must(template.New("theme").Parse(tmpl))
	return p
}
