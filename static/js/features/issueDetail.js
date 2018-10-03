// markdown editor
var gMDEditor;
// 所有需要被刷新的(element id, content)
var gAllAvatarCanvasSet = new Set();
var gAllMarkDownSet = new Set();


window.onload = function() {
  // sorry about that, description should be refreshed after the whole page
  // content were loadded.
  refreshAllMarkdown();
};


$(function() {

  /**
   * initialize the markdown editor 
   */
  $('#issue_comment').ready(function () {
    initMarkdownEditorInstance();
  });
  
  /**
   * method for uploading and insert attachement file into comment
   */
  $('#attachImage').fileupload({
    dataType: 'json',
    url: "/issuedetail/{{$issue.ID}}/newattach",
    type: 'POST',
    disableImageResize: /Android(?!.*Chrome)|Opera/.test(window.navigator.userAgent),
    imageMaxWidth: 800,
    imageMaxHeight: 800,
    imageCrop: false, // Force cropped images
    done: function (e, data) {
        if (null != data.result && null != data.result.UserInfo && 0 < data.result.UserInfo.length) {
            oldComment = gMDEditor.value();
            newComment = oldComment + '![' + data.files[0].name + '](' + data.result.UserInfo + ')';
            gMDEditor.value(newComment);
        }
    } 
  });
  
});



class jsLazyLoadModel {
  constructor(elemId, content) {
    this.elemId = elemId;
    this.content = content;
  }
}

// on finish loading
function initMarkdownEditorInstance() {
  if (null == gMDEditor) {
    gMDEditor = new SimpleMDE({elements: $('#issue_comment')[0]});
  }
}

// add elemId and content into avatar render collection
function appendAvatarCanvasCollection(elemId, content) {
  if (content.length > 0 && elemId.length > 0) {
    for (let item of gAllAvatarCanvasSet) {
      if (item.elemId == elemId) {
        return;
      }
    }

    var newModel = new jsLazyLoadModel(elemId, content);
    gAllAvatarCanvasSet.add(newModel);
  }
}

// render all avatars which listed in avatar render collection.
function refreshAllAvatar() {
  for (let item of gAllAvatarCanvasSet) {
    AvatarDrawCanvasWith(item.content, item.elemId);
  }
}

function appendMarkdownCollection(elemId, content) {
  if (content.length > 0 && elemId.length > 0) {
    for (let item of gAllMarkDownSet) {
      if (item.elemId == elemId) {
        return;
      }
    }

    var newModel = new jsLazyLoadModel(elemId, content);
    gAllMarkDownSet.add(newModel);
  }
}

// render all markdown script into html script by markdown collection. this will
// be runs after finishing load whole page.
function refreshAllMarkdown() {
  var result = '';
  for (let item of gAllMarkDownSet) {
    result = gMDEditor.markdown(item.content);
    if (result.length <= 0) {
      result = item.content;
    }
    document.getElementById(item.elemId).innerHTML = result;
  }
}

// change issue property by click the drop-down menu
function didSelectWith(id, type, desc, extID) {
  console.log(didSelectWith.caller);
  console.log('id: ' + id + ', type:' + type + ', param:' + desc);

  setInnerHtmlWithID(id + '-btn', desc);
  setHtmlValueWithID(id, type);

  issueDetailUpdate(extID, id, type);
}

// issue property change event
function issueDetailUpdate(issueId, key, value) {
  var param = Object.create(null);
  param[key] = value;

  $.ajax({
    dataType: 'json',
    method: 'post',
    url: '/issuedetail/' + issueId + '/update',
    data: $.param(param),
    success: function(result) {
      if (result == null) {
        trackCallStack();
        console.log('error: no result data');
      } else {
        if (result.Code == 302) {
          window.location.href = result.URL;
        } else if (result.Code == 200) {
          reloadDiv('issue_log_history');  // issue log history section
          reloadDiv('issue-level-band');   // colour band at top
        }
      }
    },
    error: function(result) {
      console.log(result);
    }
  });
}

// the event about adding issue's new log
function issueDetailSubmitNewLog(issueId) {
  var strOriginMD = gMDEditor.value();
  var arguData = {issue_comment: strOriginMD};

  $.ajax({
    type: 'POST',
    dataType: 'json',
    url: '/issuedetail/' + issueId + '/newlog',
    data: arguData,
    success: function(result) {
      if (result == null) {
        trackCallStack();
        console.log('error: no result data');
      } else {
        if (result.Code == 200) {
          reloadDiv('issue_log_history');
          reloadDiv('issue_log_new');
        } else if (result.Code == 302) {
          window.location.href = result.URL;
        }
      }
    },
    error: function(result) {
      console.log('Fails in register account with ' + result);
    }
  });
}