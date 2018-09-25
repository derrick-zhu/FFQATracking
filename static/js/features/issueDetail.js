class cAvatarModel {
  constructor(elemId, content) {
    this.elemId = elemId;
    this.content = content;
  }
}

// 所有需要被刷新的(element id, content)
var gAllAvatarCanvasSet = new Set();

function appendAvatarCanvasCollection(elemId, content) {
  if (content.length > 0 && elemId.length > 0) {

    for (let item of gAllAvatarCanvasSet) {
      if (item.elemId == elemId) {
        return;
      }
    }

    var newModel = new cAvatarModel(elemId, content);
    gAllAvatarCanvasSet.add(newModel);
  }
}

function refreshAllAvatar() {
  for (let item of gAllAvatarCanvasSet) {
    AvatarDrawCanvasWith(item.content, item.elemId);
  }
}


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
          reloadDiv("issue_log_history"); // issue log history section
          reloadDiv("issue-level-band"); // colour band at top
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