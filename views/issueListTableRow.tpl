{{define "issueListTableRow"}}

    {{range $issue := .allIssue}}
    <tr style="background-color:#99ff95">
        <td><input type="checkbox" name="" id=""></td>
        <td><a href='/issue_detail?id={{$issue | PropertyInIssue "ID"}}'>{{$issue | PropertyInIssue "ID"}}</a></td>
        <td><a href='/issue_detail?id={{$issue | PropertyInIssue "ID"}}'>{{$issue | PropertyInIssue "Title"}}</a></td>
        <td>{{$issue | PropertyInIssue "Status"}}</td>
        <td>{{$issue | PropertyInIssue "Priority"}}</td>
        <td><a href='#'>{{$issue | PropertyInIssue "Version"}}</a></td>
        <td><a href='/account/{{$issue | PropertyInIssue "Creator"}}/'>{{$issue | PropertyInIssue "Creator"}}</a></td>
        <td><a href='/account/{{$issue | PropertyInIssue "Assignor"}}'>{{$issue | PropertyInIssue "Assignor"}}</a></td>
        <td>{{$issue | PropertyInIssue "CreateDate"}}</td>
    </tr>
    {{end}}

{{end}}