{{define "issueListTableRow"}}

    {{$accounts := .allAccount}}

    {{range $issue := .allIssue}}
    
        {{$creator := AccountForIDInArray $accounts $issue.Creator}}
        {{$assignor := AccountForIDInArray $accounts $issue.Assignor}}
    
        <tr class='{{$issue | IssueCSSWithPriority}}' onclick="didIssueClicked('{{$issue.ID}}')">

            <td></td>
            <td>{{$issue.ID}}</td>
            <td class="tr-title">{{$issue.Title}}</td>
            <td>{{$issue | PropertyInIssue "Status"}}</td>
            <td>{{$issue | PropertyInIssue "Priority"}}</td>
            <td><a href='#'>{{$issue | PropertyInIssue "Version"}}</a></td>
            <td><a href='/account/{{$creator.ID}}/'>{{$creator.Name}}</a></td>
            <td><a href='/account/{{$assignor.ID}}/'>{{$assignor.Name}}</a></td>
            <td>{{$issue | PropertyInIssue "CreateDate"}}</td>

        </tr>
    {{end}}

{{end}}