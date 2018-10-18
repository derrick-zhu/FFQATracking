{{define "dataDatePickerTemplate"}}

<span class="span span_2of6 span_float_left text text-align-right">{{.Title}}</span>
<div id="date_{{.Identifier}}" name="date_{{.Identifier}}" class="text input-group date" data-provide="datepicker">
    <input type="text" class="form-control">
    <input id="{{.Identifier}}" name="{{.Identifier}}" type="hidden" value="">
    <div class="input-group-addon">
        <span class="glyphicon glyphicon-th"></span>
    </div>
</div>

{{end}}