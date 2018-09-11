{{define "issueListTableRow"}}

    {{range $issue := .allIssue}}
    <tr>
        <td><input type="checkbox" name="" id=""></td>
        <td>{{$issue | PropertyInIssue "ID"}}</td>
        <td><a href="#">{{$issue | PropertyInIssue "Title"}}</a></td>
        <td>{{$issue | PropertyInIssue "Status"}}</td>
        <td>{{$issue | PropertyInIssue "Priority"}}</td>
        <td>{{$issue | PropertyInIssue "Version"}}</td>
        <td>{{$issue | PropertyInIssue "Creator"}}</td>
        <td>{{$issue | PropertyInIssue "Assignor"}}</td>
        <td>{{$issue | PropertyInIssue "CreateDate"}}</td>
    </tr>
    {{end}}

{{end}}