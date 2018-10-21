{{define "dataIssueTableTemplate"}}

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
    <script type="text/javascript">
        function didIssueClicked(params) {
            if (params.length > 0) {
                var url = "/issuedetail/" + params;
                console.log(url);
                window.location.href = url;
            }
        }
    </script>

    <tbody>
        {{template "issueListTableRow" .}}
    </tbody>
</table>

{{end}}