{{define "dataTextareaTemplate"}}

<label class="span span_2of6 span_float_left text text-align-right">{{.Title}}</label>
<div class="span span_4of6 span_float_right" >
    <textarea class="form-control span span_full" style="height:10rem;max-height: 15rem;" row="10" col="30" type="text" name="name_{{.Identifier}}_{{.ID}}" id="name_{{.Identifier}}_{{.ID}}" placeholder="{{.DefaultValue}}"></textarea>
</div>

{{end}}