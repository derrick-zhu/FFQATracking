{{define "navbar"}}

<a class="navbar-brand" href="/">{{.Title}}</a>
<div>
    <ul class="nav navbar-nav">
        <li {{if .IsHome}} class="active" {{end}}><a href="/">Home</a></li>
        <li {{if .IsBugs}} class="active" {{end}}><a href="/bugs">Bugs</a></li>
        <li {{if .IsAnalysis}} class="active" {{end}}><a href="/analysis">Analysis</a></li>
    </ul>
</div>

<div class="pull-right">
    <ul class="nav navbar-nav">
        {{if .IsLogin}}
        <li><a href="/login/exit">Logout</a></li>
        {{else}}
        <li><a href="/register">Register</a></li>
        <li><a href="/login">Login</a></li>
        {{end}}
    </ul>
</div>

{{end}}
