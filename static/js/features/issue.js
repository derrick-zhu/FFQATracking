
function didSelectWith(id, type, desc) 
{
    console.log(didSelectWith.caller);
    console.log('id: ' + id + ', type:' + type + ', param:' + desc);

    setInnerHtmlWithID(id + '-btn', desc);
    setHtmlValueWithID(id, type);
}

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