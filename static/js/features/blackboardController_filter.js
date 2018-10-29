$(function() {

    $.fn.initiativePickerChanged = function(id, type, extID) {
        trackCallStack();
        console.log(id, type, extID);
        $.get(
            "/blackboard/filter/change/?initiative_id=" + type,
            {},
            function (data, result) {
                
                if (data.UserInfo != undefined) {
                    if (data.UserInfo.Param.versions != undefined) {
                        $('#datapicker-versions').html(data.UserInfo.Param.versions);
                    }
                    
                    if (data.UserInfo.Param.issues != undefined) {
                        $('#issue-table').html(data.UserInfo.Param.issues);
                    }
                }
                
            });
    };

    $.fn.milestonePickerValueChanged = function(id, type, extID) {
        trackCallStack();
        console.log(id, type, extID);
        console.log($('#initiatives').val());
    };
});