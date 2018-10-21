{{define "issueListTableRow"}}

{{$accounts := .allAccount}}

{{range $issue := .allIssue}}

    {{$creatorIdx := AccountIndexOfID $accounts $issue.Creator}}
    {{$assignorIdx := AccountIndexOfID $accounts $issue.Assignor}}
    
    {{if ge $creatorIdx 0}}{{if ge $assignorIdx 0}}
    
    {{$creator := AccountForIDInArray $accounts $issue.Creator}}
    {{$assignor := AccountForIDInArray $accounts $issue.Assignor}}

    <tr class='{{$issue | IssueCSSWithPriority}}'>

        <td></td>
        <td>{{$issue.ID}}</td>
        <td class="tr-title"><a href="/issuedetail/{{$issue.ID}}">{{$issue.Title}}</a></td>
        <td>{{PropertyInIssue "Status" $issue}}</td>
        <td>{{PropertyInIssue "Priority" $issue}}</td>
        <td><a href='#'>{{PropertyInIssue "Version" $issue}}</a></td>
        <td><a href='/account/{{$creator.ID}}/'>{{$creator.Name}}</a></td>
        <td><a href='/account/{{$assignor.ID}}/'>{{$assignor.Name}}</a></td>
        <td>{{PropertyInIssue "CreateDate" $issue}}</td>

    </tr>

    {{end}}{{end}}

{{end}}

{{end}}