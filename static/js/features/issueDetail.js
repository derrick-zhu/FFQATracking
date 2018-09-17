function didSelectWith(id, type, desc, extID) {

  console.log(didSelectWith.caller);
  console.log('id: ' + id + ', type:' + type + ', param:' + desc);

  setInnerHtmlWithID(id + '-btn', desc);
  setHtmlValueWithID(id, type);

  issueDetailUpdate(extID, id, type);
}

function issueDetailUpdate(issueId, key, value) {

  var param = Object.create(null);
  param[key] = value;

  $.ajax({
    dataType: 'json',
    method: 'post',
    url: '/issuedetail/' + issueId + '/update',
    data: $.param(param),
    success: function (result) {
      if (result == null) {
        trackCallStack();
        console.log("error: no result data");
      } else {
        if (result.Code == 302) {
          window.location.href = result.URL;
        } else if (result.Code == 200) {
          reloadDiv("issue_log_history");
        }
      }
    },
    error: function (result) {
      console.log(result);
    }
  });
}

function issueDetailSubmitNewLog(issueId) {

  $.ajax({
    type: 'POST',
    dataType: 'json',
    url: '/issuedetail/' + issueId + '/newlog',
    data: $('#frmIssueDetail').serialize(),
    success: function (result) {
      if (result == null) {
        trackCallStack();
        console.log("error: no result data");
      } else {
        if (result.Code == 200) {
          reloadDiv("issue_log_history");
          reloadDiv("issue_log_new");
        } else if (result.Code == 302) {
          window.location.href = result.URL;
        }
      }
    },
    error: function (result) {
      console.log('Fails in register account with ' + result);
    }
  });
}