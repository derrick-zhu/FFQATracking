{{define "dataPickerTemplate"}}    
<!-- 
    param:
        title           string
        Identifier      string <out>
        DefaultValue    int64       // this is a index,
        collection      []interface{}
 -->
 
 {{$title := .Title}}
 {{$defaultContent := (index .Collection .DefaultValue | GetBriefTitleFromModel)}}
 {{$defaultValue := (index .Collection .DefaultValue | GetTypeFromModel)}}
 {{$id := .Identifier}}
 {{$data := .Collection}}

 <div class="form-group">
    <div class="col-xs-3" style="margin:0.5px;"> 
        <label class="right label-ff-standard" style="width:100px">{{$title}}</label>
        <div class="btn-group">
            <button id="{{$id}}-btn" name="{{$id}}-btn" type="button" class="btn btn-normal">{{$defaultContent}}</button>
            <input type="hidden" class="form-control" id="{{$id}}" name="{{$id}}" value="{{$defaultValue}}">
            <button type="button" class="btn btn-normal dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                <span class="caret"></span>
                <span class="sr-only">Toggle Dropdown</span>
            </button>

            <ul class="dropdown-menu" style="max-height: 15em; overflow-y: scroll">
                {{range $index, $item := $data}}
                    {{$itemValue := ($item | GetTypeFromModel)}}
                    {{$itemContent := ($item | GetBriefTitleFromModel)}}
                    <li><a onclick="return didSelectWith('{{$.Identifier}}', '{{$itemValue}}', '{{$itemContent}}');" value="{{$itemValue}}">{{$itemContent}}</a></li>
                {{end}}
            </ul>
        </div>
    </div>
</div>


{{end}}