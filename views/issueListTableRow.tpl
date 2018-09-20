{{define "issueListTableRow"}}

    {{$accounts := .allAccount}}

    {{range $issue := .allIssue}}
    
        {{$creator := AccountForIDInArray $accounts $issue.Creator}}
        {{$assignor := AccountForIDInArray $accounts $issue.Assignor}}
    
        <tr class='{{$issue | IssueCSSWithPriority}}' onclick="didIssueClicked('{{$issue.ID}}')">

            <td></td>
            <td>{{$issue.ID}}</td>
            <td class="tr-title">{{$issue.Title}}</td>
            <td>{{PropertyInIssue "Status" $issue}}</td>
            <td>{{PropertyInIssue "Priority" $issue}}</td>
            <td><a href='#'>{{PropertyInIssue "Version" $issue}}</a></td>
            <td><a href='/account/{{$creator.ID}}/'>{{$creator.Name}}</a></td>
            <td><a href='/account/{{$assignor.ID}}/'>{{$assignor.Name}}</a></td>
            <td>{{PropertyInIssue "CreateDate" $issue}}</td>

        </tr>
    {{end}}

{{end}}