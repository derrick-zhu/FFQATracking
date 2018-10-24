$(function () {
    $('#btnCommitNewIssue').click(function () {
        trackCallStack();
        console.log($('#frmSubmitNewIssue').serialize());
        $.post(
            '/blackboard/newissue',
            {

            },
            function (data, status) {
                window.location.href = "#";
            }
        );
    });
});