{{define "dataTextareaTemplate"}}

<div class="span span_2of6 span_float_left text text-align-right">{{.Title}}</div>
<div class="span span_4of6 span_float_right" >
    <textarea style="width:100%" type="text" name="name_{{.Identifier}}_{{.ID}}" id="name_{{.Identifier}}_{{.ID}}" placeholder="{{.DefaultValue}}"></textarea>
</div>

{{end}}