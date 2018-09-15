function regiseter() {
    $.ajax({
        type: "POST",
        dataType: "json",
        url: "/register",
        data: $('#frmRegister').serialize(),
        success: function (result) {
            console.log(result);
            if (result && (result.Code == 200 || result.Code == 302)) {
                window.location.href = result.URL;
            }
        },
        error: function (result) {
            console.log("Fails in register account with " + result);
        }
    });
}