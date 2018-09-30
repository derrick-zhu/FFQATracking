{{define "issueLogItem"}}

<!-- avatar canvas' local variable -->
{{$avatarName := printf "user_avatar_%d" .CreatorID}}

<div class="col span_2of6">
    <div class="log-user-icon">
        <canvas id="{{$avatarName}}" name="{{$avatarName}}" class="log-avatar" width="48px" height="48px"></canvas>
        <script type="text/javascript">
            $('#{{$avatarName}}').ready(function () {
                appendAvatarCanvasCollection(
                    "{{$avatarName}}",
                    "{{.CreatorName}}");
            });
        </script>
    </div>
    <div class="log-user-info">
        <div>
            <p style="font-size:1.75rem;"><a href="/person/{{.CreatorID}}">{{.CreatorName}}</a></p>
        </div>
        <div>
            <p style="font-size:1.5rem;">{{.TimeDisplay}}</p>
        </div>
    </div>
</div>

<div class="col span_4of6">
    {{if eq .Type 1}}
    <p>Issue status {{.StatusTitle}} had been changed from {{BugStatusWithType .PrvStatus}} into {{BugStatusWithType .NewStatus}}</p>
    {{else}}
    <div id="issue_comment_{{.ID}}"></div>
    <script>
        $("#issue_comment_{{.ID}}").ready(function () {
            appendMarkdownCollection(
                "issue_comment_{{.ID}}", 
                "{{.Content}}")
        });
    </script>
    {{end}}

    <script type="text/javascript">
        $('#{{$avatarName}}').ready(function () {
            appendAvatarCanvasCollection(
                "{{$avatarName}}",
                "{{.CreatorName}}");
        });
    </script>
</div>

<hr>

{{end}}