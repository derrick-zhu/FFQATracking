
function didSelectWith(id, type, desc, extID) {

}


function issueDetailUpdate(issueId) {
    
}

function issueDetailSubmitNewLog(issueId) {
    
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/issuedetail/" + issueId + "/newlog",
        data: $('#frmIssueDetail').serialize(),
        success: function (result) {
            console.log(result);
            if (result && (result.Code == 200 || result.Code == 302)) {
                
            }
        },
        error: function (result) {
            console.log("Fails in register account with " + result);
        }
    });
}