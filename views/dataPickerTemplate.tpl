{{define "dataPickerTemplate"}}    
<!-- 
    param:
        ID              int64       // external index for anything
        title           string
        Identifier      string <out>
        DefaultValue    int64       // this is a index,
        collection      []interface{}
 -->
 
 {{$extID := .ID}}
 {{$title := .Title}}
 {{$defaultContent := (index .Collection .DefaultValue | GetBriefTitleFromModel)}}
 {{$defaultValue := (index .Collection .DefaultValue | GetTypeFromModel)}}
 {{$id := .Identifier}}
 {{$data := .Collection}}


<label class="span span_2of6 span_float_left text text-align-right">{{$title}}</label>
<div class="btn-group span span_4of6 span_float_right">
    <button id="{{$id}}-btn" name="{{$id}}-btn" type="button" class="btn btn-normal" style="width:90%;text-align:left;">{{$defaultContent}}</button>
    <input type="hidden" class="form-control" id="{{$id}}" name="{{$id}}" value="{{$defaultValue}}" style="width:95%;">
    <button type="button" class="btn btn-normal dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" style="width:8%;min-width:1rem;padding-left: 4px;padding-right: 10px;">
        <span class="caret"></span>
        <span class="sr-only">Toggle Dropdown</span>
    </button>

    <ul class="dropdown-menu" style="width:90%; max-height:20em; overflow-y: scroll">
        {{range $index, $item := $data}}
            {{$itemValue := GetTypeFromModel $item}}
            {{$itemContent := GetBriefTitleFromModel $item}}
            <li><a onclick="return didSelectWith('{{$.Identifier}}', '{{$itemValue}}', '{{$itemContent}}', '{{$extID}}');" value="{{$itemValue}}">{{$itemContent}}</a></li>
        {{end}}
    </ul>
</div>

{{end}}