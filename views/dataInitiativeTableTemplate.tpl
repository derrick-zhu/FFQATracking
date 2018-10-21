{{define "dataIntiativeTemplate"}}

<table class="table table-hover table-condensed">
    <thead>
        <tr>
            <th></th>
            <th>#</th>
            <th>Title</th>
            <th>Creator</th>
            <th>Assignor</th>
        </tr>
    </thead>
    <tbody>
        {{range $initiative := .}}
        <tr>
            <td></td>
            <td>{{$initiative.ID}}</td>
            <td>{{$initiative.Name}}</td>
            <td>{{$initiative.CreatorName}}</td>
            <td>{{$initiative.AssignorName}}</td>
        </tr>
        {{end}}
    </tbody>
</table>


{{end}}