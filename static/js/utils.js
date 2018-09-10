
function checkInputLength(id, min, max) 
{
    if (id.length <= 0)  {
        return false;
    }

    var eleValue = document.getElementById(id);
    var lenValue = eleValue.value.length;

    return (lenValue >= min && lenValue <= max);
}

function checkWithRegex(id, pattern) 
{
    var idValue = document.getElementById(id);
    var result = idValue.value.match(pattern);
    return result != null;
}

function setInnerHtmlWithID(id, params) {

    console.log('setInnerHtmlWithID('+ id + ', ' + params + ')');
    
    var x = document.getElementById(id);
    x.innerHTML = params;
}

function setHtmlValueWithID(id, params) {

    console.log('setValueWithID('+ id + ', ' + params + ')');

    var x = document.getElementById(id);
    x.value = params;
}
