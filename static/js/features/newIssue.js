
function newIssueCheckInputContent() 
{    
    if (checkInputLength("issueTitle", 1, 1024) == false) {
        alert("invalid title");
        return false;
    }

    if (checkInputLength("issueContent", 1, 3096) == false) {
        alert("invalid description");
        return false;
    }
}