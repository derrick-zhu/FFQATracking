$(function() {

    $.fn.initiativePickerChanged = function(id, type, extID) {
        trackCallStack();
        var queryUri = $().fnRouterForBlackboardPageWith(-1, type, -1);
        window.location.href = queryUri;
    };

    $.fn.milestonePickerValueChanged = function(id, type, extID) {
        trackCallStack();
        console.log(id, type, extID);
        var initiativeID = $('#initiatives').val();
        var queryUri = $().fnRouterForBlackboardPageWith(-1, initiativeID, type);
        window.location.href = queryUri;
    };
});