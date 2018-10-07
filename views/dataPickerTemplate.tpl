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

<div class="col-xs-3" style="margin:0.5px;"> 
    <label class="left standard-ff">{{$title}} :</label>
    <div class="btn-group">
        <button id="{{$id}}-btn" name="{{$id}}-btn" type="button" class="btn btn-normal" style="max-width:100%">{{$defaultContent}}</button>
        <input type="hidden" class="form-control" id="{{$id}}" name="{{$id}}" value="{{$defaultValue}}" style="max-width:100%">
        <button type="button" class="btn btn-normal dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" style="max-width:100%">
            <span class="caret"></span>
            <span class="sr-only">Toggle Dropdown</span>
        </button>

        <ul class="dropdown-menu" style="max-height: 20em; overflow-y: scroll">
            {{range $index, $item := $data}}
                {{$itemValue := GetTypeFromModel $item}}
                {{$itemContent := GetBriefTitleFromModel $item}}
                <li><a onclick="return didSelectWith('{{$.Identifier}}', '{{$itemValue}}', '{{$itemContent}}', '{{$extID}}');" value="{{$itemValue}}">{{$itemContent}}</a></li>
            {{end}}
        </ul>
    </div>
</div>


{{end}}