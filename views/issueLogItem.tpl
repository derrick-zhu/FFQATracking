{{define "issueLogItem"}}

<!-- avatar canvas' local variable -->
{{$avatarName := printf "user_avatar_%d" .CreatorID}}

<div class="card clearfix">

    <div class="col span span_1of6">
        <div class="log-user-icon">
            <canvas class="avatar" id="{{$avatarName}}" name="{{$avatarName}}" width="48px" height="48px"></canvas>
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

    <div class="col span span_5of6 card-content">

        <!-- 只有提交评论的用户才能够删除当前comment -->
        {{if .IsViewersComment}}
        <div class="card-button-area">
            <a id="btnDeleteComment" value="{{.ID}}" type="submit" class="btn btn-xs btn-danger" onclick="$().fnDeleteComment({{.IssueID}}, {{.ID}});">
                <span class="glyphicon glyphicon-remove" aria-hidden="true"></span>
            </a>
        </div>
        {{end}}

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
    </div>

</div>

{{end}}