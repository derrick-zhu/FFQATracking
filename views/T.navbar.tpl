{{define "navbar"}}

{{$account := .AccountData}}

<div>
    <ul class="nav navbar-nav">
        <li><img src="/static/img/ftech.jpg" height="50em"></li>
        <li {{if .IsHome}} class="active" {{end}}><a href="/">Home</a></li>
        <li {{if .IsBlackBoard}} class="active" {{end}}><a href="/blackboard">Blackboard</a></li>
        <li {{if .IsAnalysis}} class="active" {{end}}><a href="/analysis">Analysis</a></li>
    </ul>
</div>

<div class="pull-right">
    <ul class="nav navbar-nav">
        {{if .IsLogin}} {{if $account}}
        <li><a href="/account/{{$account.ID}}">Welcome, {{$account.Email}}</a></li>
        {{end}}
        <li><a href="/login/exit">Logout</a></li>
        {{else}}
        <li><a href="/register">Register</a></li>
        <li><a href="/login">Login</a></li>
        {{end}}
    </ul>
</div>

{{end}}