{{define "bbNewIssueTemplate"}}

<div class="modal fade" id="bbNewIssueModal" tabindex="-1" role="dialog" aria-labelledby="bbNewIssueModalLabel">
    <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                <h4 class="modal-title" id="bbNewIssueModalLabel">New Issue</h4>
            </div>
            <div class="modal-body clearfix">

                <div>
                    <form role="form" id="frmSubmitNewIssue" method="post" action="#" onsubmit="return false;">
                        <div class="form-group">
                            <label>Title: *</label>
                            <input id="Title" name="Title" type="text" class="form-control" placeholder="Issue Title">
                        </div>

                        <div class="form-group">
                            <label>Description: *</label>
                            <textarea id="Description" name="Description" class="form-control" placeholder="Please input the description here."
                                rows="10"></textarea>
                        </div>

                        <!-- bug的各种状态设置 -->
                        {{range $idx, $issueTpl := .}}
                        <div class="span span_3of6 span_float_left" style="margin:0.5px;">
                            {{$ctrlType := ControllerTypeOfTemplateData $issueTpl}}
                            {{if eq $ctrlType 3}}
                            {{template "dataPickerTemplate" $issueTpl}}
                            {{end}}
                        </div>
                        {{end}}

                        <!-- 分割线 -->
                        <!-- <div class="form-group col-xs-12">
                            <hr class="seperate-line" width="80%" color=#987cb9 SIZE=3>
                        </div> -->

                        <!-- 提交按钮 -->
                        <!-- <div class="form-group col-xs-12">
                            <input class="btn btn-danger center-block" style="width:150px;" type="submit" value="Submit"
                                onclick="return newIssueCheckInputContent();">
                        </div> -->

                    </form>
                </div>

            </div>
            <div class="modal-footer">
                <input id="btnCommitNewIssue" name="btnCommitNewIssue" class="btn btn-danger" style="width:150px;" type="submit" value="Submit">
                    <!-- <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                    <button type="button" class="btn btn-primary">Save changes</button> -->
            </div>
        </div>
    </div>
</div>

{{end}}