{{define "issueListTableRow"}}

    {{range $issue := .allIssue}}
    {{$issueId := PropertyInIssue "ID" $issue}}
        <tr class='{{$issue | IssueCSSWithPriority}}' onclick="didIssueClicked('{{$issueId}}')">

            <td></td>
            <td>{{$issue | PropertyInIssue "ID"}}</td>
            <td class="tr-title">{{$issue | PropertyInIssue "Title"}}</td>
            <td>{{$issue | PropertyInIssue "Status"}}</td>
            <td>{{$issue | PropertyInIssue "Priority"}}</td>
            <td><a href='#'>{{$issue | PropertyInIssue "Version"}}</a></td>
            <td><a href='/account/{{$issue | PropertyInIssue "Creator"}}/'>{{$issue | PropertyInIssue "Creator"}}</a></td>
            <td><a href='/account/{{$issue | PropertyInIssue "Assignor"}}'>{{$issue | PropertyInIssue "Assignor"}}</a></td>
            <td>{{$issue | PropertyInIssue "CreateDate"}}</td>

        </tr>
    {{end}}

{{end}}