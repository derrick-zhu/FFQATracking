{{define "dataIssueTableTemplate"}}
<div>
    <span class="btn btn-success fileinput-button" style="width:120px">
        <i class="glyphicon glyphicon-plus"></i>
        <span>New Issue...</span>
        <input id="btnAttachImage" name="attachImage" type="button" hidden="true">
    </span>
</div>
<div>
    <table class="table table-hover table-condensed">
        <thead>
            <tr>
                <th></th>
                <th>#</th>
                <th class="th-big">Title</th>
                <th>Status</th>
                <th>Priority</th>
                <th>Version</th>
                <th>Creator</th>
                <th>Assignor</th>
                <th>Create Date</th>
            </tr>
        </thead>

        <tbody>
            {{template "issueListTableRow" .}}
        </tbody>
    </table>
</div>
{{end}}