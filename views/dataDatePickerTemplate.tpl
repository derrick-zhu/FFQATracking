{{define "dataDatePickerTemplate"}}

<span class="span span_2of6 span_float_left text text-align-right">{{.Title}}</span>
<div id="date_{{.Identifier}}" name="date_{{.Identifier}}" class="text input-group date" style="padding: 6px 0px 6px 12px;" data-provide="datepicker">
    <input id="{{.Identifier}}" name="{{.Identifier}}" type="text" class="form-control">
    <div class="input-group-addon">
        <span class="glyphicon glyphicon-th"></span>
    </div>
</div>
<input id="value_{{.Identifier}}" name="value_{{.Identifier}}" type="hidden" value="">

{{end}}