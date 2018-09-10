{{define "issueFormCreate"}}    
<!-- 
    param:
        title           string
        Identifier      string <out>
        defaultValue    string
        collection      []string
 -->
 <div class="form-group">
    <div class="col-xs-3" style="margin:0.5px;"> 
        <label class="right label-ff-standard" style="width:100px">{{.Title}}</label>
        <div class="btn-group">
            <button id="{{.Identifier}}-btn" name="{{.Identifier}}-btn" type="button" class="btn btn-normal">{{index .Collection .DefaultValue}}</button>
            <input type="hidden" class="form-control" id="{{.Identifier}}" name="{{.Identifier}}" value="{{index .Collection .DefaultValue}}">
            <button type="button" class="btn btn-normal dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                <span class="caret"></span>
                <span class="sr-only">Toggle Dropdown</span>
            </button>

            <ul class="dropdown-menu" style="max-height: 15em; overflow-y: scroll">
                {{range .Collection}}
                <li><a onclick="return didSelectWith('{{$.Identifier}}', '{{.}}');">{{.}}</a></li>
                {{end}}
            </ul>
        </div>
    </div>
</div>


{{end}}