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


<span class="span span_2of6 span_float_left text text-align-right" style="line-height: 1.425">{{$title}}</span>
<div class="btn-group span span_4of6 span_float_right" style="line-height: 1.425">
    <!-- <button id="{{$id}}-btn" name="{{$id}}-btn" type="button" class="btn btn-normal" style="width:90%;text-align:left;">{{$defaultContent}}</button> -->
    
    <button id="{{$id}}-btn" name="{{$id}}-btn" type="button" class="btn btn-normal dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" style="width:100%; min-width:1rem; text-align:left; padding-left: 4px; padding-right: 10px;">
        {{$defaultContent}}<span class="caret" style="position:absolute; vertical-align:middle; right: 1rem; top:50%"></span>
    </button>
    <input type="hidden" class="form-control" id="{{$id}}" name="{{$id}}" value="{{$defaultValue}}" style="width:95%;">

    <ul class="dropdown-menu" style="width:90%; max-height:20em; overflow-y: scroll">
        {{range $index, $item := $data}}
            {{$itemValue := GetTypeFromModel $item}}
            {{$itemContent := GetBriefTitleFromModel $item}}
            <li><a onclick="$().fnDataPickerDidChangeValue('{{$.Identifier}}', '{{$itemValue}}', '{{$itemContent}}', '{{$extID}}');" value="{{$itemValue}}">{{$itemContent}}</a></li>
        {{end}}
    </ul>
</div>

{{end}}