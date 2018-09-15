function login() {
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/login/signin",
        data: $('#frmLogin').serialize(),
        success: function (result) {
            console.log(result);
            if (result && (result.Code == 200 || result.Code == 302)) {
                window.location.href = result.URL;
            }
        },
        error: function (result) {
            console.log(result);
            alert('error: ' + result);
        }
    });
}