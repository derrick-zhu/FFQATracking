$(function () {

    $('#date_startDate').datepicker({
        format: 'yyyy-mm-dd',
        startDate: '-3d',
        autoclose: true,
    }).on('changeDate', function (ev) {
        $('#value_startDate').val(ev.timeStamp);
    });

    $('#date_endDate').datepicker({
        format: 'yyyy-mm-dd',
        startDate: '-3d',
        autoclose: true,
    }).on('changeDate', function (ev) {
        $('#value_endDate').val(ev.timeStamp);
    });

    $('#btnCommitNewProject').click(function(){
        trackCallStack();
        console.log($('#frmProjectProperties').serialize());
        $.post(
            '/blackboard/newinitiative', 
            {
                title: $('#name_title_0').val(),
                description: $('#name_description_0').val(),
                creator: $('#creator').val(),
                assignor: $('#assignor').val(),
                startDate: $('#value_startDate').val(),
                endDate: $('#value_endDate').val(),
            },
            function(data, status) {
                if ((data.Code >= 200) && (data.Code < 400)) {
                    $.toast({
                        text: "Success in create new project!",
                        icon: "success",
                        hideAfter: 1200,
                        position: "bottom-center",
                        afterHidden: function() {
                            $('#bbNewInitiativeModal').modal("hide");
                        }
                    });
                }
                else {
                    $.toast({
                        text: "Fails in submit new project!",
                        icon: "error",
                        hideAfter: 2000,
                        position: "bottom-center",
                        afterHidden: function() {
                            // window.location.href = data.URL;
                        }
                    });
                }
            }
        );
    });
});