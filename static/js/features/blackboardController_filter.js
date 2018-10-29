$(function() {

    $.fn.initiativePickerChanged = function(id, type, extID) {
        trackCallStack();
        console.log(id, type, extID);
        $.get(
            "/blackboard/filter/change/?initiative_id=" + type,
            {},
            function (data, result) {
                console.log(result);
                $('#datapicker-versions').html(data.UserInfo);
            });
    };

    $.fn.milestonePickerValueChanged = function(id, type, extID) {
        trackCallStack();
        console.log(id, type, extID);
    };
});