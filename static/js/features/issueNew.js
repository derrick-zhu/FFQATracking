$(function () {
  $.fn.fnDataPickerDidChangeValue = function (id, type, desc, extID) {
    trackCallStack()
    console.log('id: ' + id + ', type:' + type + ', param:' + desc)

    setInnerHtmlWithID(id + '-btn', desc)
    setHtmlValueWithID(id, type)
  };

  $('#btnCommitNewIssue').click(function () {
    trackCallStack();
    console.log($('#frmSubmitNewIssue').serialize());
    $.post(
        '/blackboard/newissue',
        {

        },
        function(data, status){
            window.location.href = "#";
        }
    );
  });
});

function newIssueCheckInputContent () {
  if (checkInputLength('issueTitle', 1, 1024) == false) {
    alert('invalid title');
    return false;
  }

  if (checkInputLength('issueContent', 1, 3096) == false) {
    alert('invalid description');
    return false;
  }
}
