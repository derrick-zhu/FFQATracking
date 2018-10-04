// markdown editor
var gMDEditor = null;
// 所有需要被刷新的(element id, content)
var gAllAvatarCanvasSet = new Set();
var gAllMarkDownSet = new Set();

// model for lazy loading
class jsLazyLoadModel {
  constructor(elemId, content) {
    this.elemId = elemId;
    this.content = content;
  }
}


$(function () {
  $('#issue_detail_last_div').ready(function() {
    initMarkdownEditorInstance();
    refreshAllMarkdown();
    refreshAllAvatar();
  });

  /**
   * method for uploading and insert attachement file into comment
   */
  $('#btnAttachImage').fileupload({
    dataType: 'json',
    type: 'POST',
    loadImageFileTypes: /^image\/(gif|jpeg|jpg|png|svg\+xml)$/, // ?? 貌似目前没有作用
    disableImageResize: /Android(?!.*Chrome)|Opera/.test(window.navigator.userAgent),
    imageMaxWidth: 800,
    imageMaxHeight: 800,
    imageCrop: false, // Force cropped images
    done: function (e, data) {
      if (null != data.result && null != data.result.UserInfo &&
        0 < data.result.UserInfo.length) {
        oldComment = gMDEditor.value();
        newComment = oldComment + '![' + data.files[0].name + '](' +
          data.result.UserInfo + ')';
        gMDEditor.value(newComment);
      }
    }
  });

  /**
   * 提交评论
   */
  $('#btnCommitComment').click(function () {
    $.post(
      '/issuedetail/' + this.name + '/newlog', {
        issue_comment: gMDEditor.value(),
      },
      function (data, status) {
        window.location.href = window.location.href;
      });
  });
});



// on finish loading
function  initMarkdownEditorInstance() {
  trackCallStack();
  if (null == gMDEditor) {
    gMDEditor = new SimpleMDE({
      elements: $('#issue_comment')[0],
      toolbar: [
        'bold', 'italic', 'heading', 'code', 'unordered-list', 'ordered-list',
        '|', 'link', 'table', 'horizontal-rule', '|', 'preview', 'side-by-side',
        'fullscreen'
      ]
    });
  }
}

// add elemId and content into avatar render collection
function appendAvatarCanvasCollection(elemId, content) {
  trackCallStack();
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
  trackCallStack();
  for (let item of gAllAvatarCanvasSet) {
    AvatarDrawCanvasWith(item.content, item.elemId);
  }
}

function appendMarkdownCollection(elemId, content) {
  trackCallStack();
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
  trackCallStack();
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
  trackCallStack();

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
    success: function (result) {
      if (result == null) {
        trackCallStack();
        console.log('error: no result data');
      } else {
        if (result.Code == 302) {
          window.location.href = result.URL;
        } else if (result.Code == 200) {
          reloadDiv('issue_log_history'); // issue log history section
          reloadDiv('issue-level-band'); // colour band at top
        }
      }
    },
    error: function (result) {
      console.log(result);
    }
  });
}