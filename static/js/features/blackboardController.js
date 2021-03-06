
$(function() {

    $.fn.fnRouterForBlackboardPageWith = function(sprint, initiative, milestone) {
        var query = [];

        if (sprint >= 0) {
            query.push('sprint=' + sprint);
        }

        if (initiative >= 0) {
            query.push('proj=' + initiative);
        }

        if (milestone >= 0) {
            query.push('ms=' + milestone);
        }

        if (query.length > 0) {
            return "/blackboard/?" + query.join('&');
        } else {
            return "/blackboard";
        }
    };

});