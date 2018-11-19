{{define "bbNewInitiativeTemplate"}}
<div class="modal fade" id="bbNewInitiativeModal" tabindex="-1" role="dialog" aria-labelledby="bbNewInitiativeModalLabel">
    <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                <h4 class="modal-title" id="bbNewInitiativeModalLabel">New Project</h4>
            </div>
            <div class="modal-body">

                <div>
                    <div class="clearfix">
                        <div class="span span_full">
                            <form id="frmProjectProperties" method="post" action="#" onsubmit="return false;">
                                {{range $index, $section := .}}
                                <div class="span span_full clearfix">
                                    <!-- {{$section}} -->
                                    {{- $ctrlType := ControllerTypeOfTemplateData $section}}

                                    {{- if eq $ctrlType 1}}
                                    {{template "dataFieldTemplate" $section}}

                                    {{- else if eq $ctrlType 2}}
                                    {{template "dataTextareaTemplate" $section}}

                                    {{- else if eq $ctrlType 3}}
                                    {{template "dataPickerTemplate" $section}}

                                    {{- else if eq $ctrlType 4}}
                                    {{template "dataDatePickerTemplate" $section}}

                                    {{- end}}
                                </div>
                                {{end}}
                            </form>
                        </div>
                    </div>
                </div>

            </div>
            <div class="modal-footer">
                <!-- <button type="button" class="btn btn-default" data-dismiss="modal">Close</button> -->
                <input id="btnCommitNewProject" name="btnCommitNewProject" type="submit" class="btn btn-danger" style="width:150px;" value="Create Project">
                <!-- <button type="button" class="btn btn-primary">Save changes</button> -->
            </div>
        </div>
    </div>
</div>
{{end}}