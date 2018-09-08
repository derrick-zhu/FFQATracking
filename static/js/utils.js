
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