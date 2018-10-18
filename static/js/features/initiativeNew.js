$(function () {

    /**
     * issue的属性选择器value发生变化时
     * @param {*} id 
     * @param {*} type
     * @param {*} desc 
     * @param {*} extID
     */
    $.fn.fnDataPickerDidChangeValue = function (ID, type, desc, extID) {
        trackCallStack();

        setInnerHtmlWithID(ID + '-btn', desc);
        setHtmlValueWithID(ID, type);
    };

    $('#date_startDate').datepicker({
        format: 'yyyy-mm-dd',
        startDate: '-3d',
        autoclose: true,
    }).on('changeDate', function (ev) {
        
        console.log('start date: ' + ev.timeStamp);
    });

    $('#date_endDate').datepicker({
        format: 'yyyy-mm-dd',
        startDate: '-3d',
        autoclose: true,
    }).on('changeDate', function (ev) {
        console.log('end date: ' + ev.timeStamp);
    });

    $('#btnCommitNewProject').click(function(){
        trackCallStack();
        console.log($('#frmProjectProperties').serialize());
        $.post(
            '/initiative/new', 
            {
                title: $('#name_title_0').val(),
                description: $('#name_description_0').val(),
                creator: $('#creator').val(),
                assignor: $('#assignor').val(),
                startdate: $('#date_startDate').val(),
                endDate: $('#date_endDate').val(),
            },
            function(data, status) {
                console.log(status);
                window.location.href = window.location.href;
            }
        );
    });
});