{{define "navbar"}}

<div>
    <ul class="nav navbar-nav">
        <li>
            <svg data-test="ff-logo-icon" class="ff-logo-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 500 67" width="12em" height="4em">
                <a xlink:href="/" href="/" data-ffref="chk_hd_logo" data-tstid="Go_HomePage" aria-label="Farfetch Logo"> <g
                        fill="#222222" fill-rule="evenodd">
                        <title>Farfetch Logo</title>
                        <rect style="stroke:none; fill: #fffff; fill-opacity:0" x="0" y="0" height="100%" width="100%"></rect>
                        <polyline points="13.35 13.34 40 13.34 40 0 13.35 0 13.35 0 0 13.34 0 66.54 13.35 66.54 13.35 39.95 33.36 39.95 33.36 26.61 13.35 26.61 13.35 13.34"></polyline>
                        <path d="M77.63,26.61 L90.9,26.61 L90.9,13.34 L77.63,13.34 L77.63,26.61 Z M90.9,0 L77.63,0 L64.28,13.34 L64.28,66.54 L77.63,66.54 L77.63,39.95 L90.9,39.95 L90.9,66.54 L104.25,66.54 L104.25,13.34 L90.9,0 Z"></path>
                        <polyline points="459.52 0 459.52 66.54 472.87 66.54 472.87 39.95 486.14 39.95 486.14 66.54 499.49 66.54 499.49 0 486.14 0 486.14 26.61 472.87 26.61 472.87 0 459.52 0"></polyline>
                        <polyline points="146.63 26.61 146.63 13.34 160.04 13.34 160.04 26.61 173.39 26.61 173.39 13.34 160.04 0 133.33 0 133.33 66.54 146.63 66.54 146.63 39.95 160.04 39.95 160.04 66.54 173.39 66.54 173.39 39.95 160.04 26.61 146.63 26.61"></polyline>
                        <polyline points="215.73 13.34 242.38 13.34 242.38 0 215.72 0 202.38 13.34 202.38 66.54 215.72 66.54 215.72 39.95 235.74 39.95 235.74 26.61 215.72 26.61 215.72 13.34 215.73 13.34"></polyline>
                        <polyline points="370.96 0 331 0 330.95 13.34 344.28 13.34 344.28 66.54 357.63 66.54 357.63 13.34 370.96 13.34 370.96 0"></polyline>
                        <polyline points="277.65 53.22 277.63 53.22 277.63 39.95 297.65 39.95 297.65 26.61 277.63 26.61 277.63 13.34 304.3 13.34 304.3 0 277.63 0 264.3 13.34 264.29 13.34 264.29 53.22 264.31 53.22 277.63 66.54 304.3 66.54 304.3 53.22 277.65 53.22"></polyline>
                        <polyline points="408.59 53.22 408.59 13.34 435.23 13.34 435.23 0 408.58 0 395.24 13.34 395.24 53.22 408.58 66.54 435.24 66.54 435.24 53.22 408.59 53.22"></polyline>
                    </g>
                </a>
            </svg>
        </li>
        <li {{if .IsHome}} class="active" {{end}}><a href="/">Home</a></li>
        <li {{if .IsIssueList}} class="active" {{end}}><a href="/issuelist">Bugs</a></li>
        <li {{if .IsAnalysis}} class="active" {{end}}><a href="/analysis">Analysis</a></li>
    </ul>
</div>

<div class="pull-right">
    <ul class="nav navbar-nav">
        {{if .IsLogin}} {{if .LoggedInAccount}}
        <li><a href="/account">Welcome, {{.LoggedInAccount}}</a></li>
        {{end}}
        <li><a href="/login/exit">Logout</a></li>
        {{else}}
        <li><a href="/register">Register</a></li>
        <li><a href="/login">Login</a></li>
        {{end}}
    </ul>
</div>

{{end}}