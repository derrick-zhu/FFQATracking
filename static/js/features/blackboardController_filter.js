$(function() {

    $.fn.initiativePickerChanged = function(id, type, extID) {
        trackCallStack();
        console.log(id, type, extID);
        $.get(
            "/blackboard/filter/change/?initiative_id=" + type,
            {},
            function (data, result) {
                console.log(result);
                var divToLoad = ' #datapicker-initiatives';
                $(divToLoad).load(window.location.href + divToLoad);
            });
    };

    $.fn.milestonePickerValueChanged = function(id, type, extID) {
        trackCallStack();
        console.log(id, type, extID);
    };
});