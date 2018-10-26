$(function() {

    $.fn.initiativePickerChanged = function(id, type, extID) {
        trackCallStack();
        console.log(id, type, extID);
        $.get(
            "/blackboard/filter/change/?initiative_id=" + type,
            {},
            function (data, result) {
                console.log(result);
                // window.location.href = '#';
                // var divToLoad = ' #datapicker-versions';
                $('#datapicker-versions').html('sdas');
                // $('#issue-table').load(window.location.href + ' #issue-table');
            });
    };

    $.fn.milestonePickerValueChanged = function(id, type, extID) {
        trackCallStack();
        console.log(id, type, extID);
    };
});