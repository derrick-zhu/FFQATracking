{{define "modalPopupTemplate"}}

<div type="button" class="btn btn-primary" data-toggle="modal" data-target={{.JSCmd.ID}} value="{{.JSCmd.Name}}">
    <i class="glyphicon glyphicon-plus" style="vertical-align: middle;"></i>
</div>

{{end}}