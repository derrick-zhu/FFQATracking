$(function() {
    // 添加新的milestone： 
    // 1, project(initiative)的当前版本号
    // 2, 输入的milestone值
    // 3, 调用 添加milestone 方法
    $('#btnCommitNewMilestone').click(function() {

        var projID = $("#initiatives").val();
        var newMilestoneID = $("#name_version_0").val();

        if ((projID.length == 0) || newMilestoneID.length == 0) {
            $.toast({
                text: 'invalid initiative or milestone.',
                icon: 'error',
                hideAfter: "2500",
                position: "top-center",
            });
            return;
        }

        $.post(
            '/blackboard/newmilestone',
            {
                project: projID,
                milestone: newMilestoneID,
            },
            function(data, status) {
                if (data.Code >= 200 && data.Code < 400) {
                    $.toast({
                        text: "Success in create new milestone!",
                        icon: "success",
                        hideAfter: 1200,
                        position: "bottom-center",
                        afterHidden: function() {

                            $('#bbNewMilestoneModal').modal("hide");

                            uri = $().fnGotoBlackboardPage(-1, data.UserInfo.Param.proj, data.UserInfo.Param.msid);
                            if (uri.length > 0) {
                                window.location.href = '/blackboard?' + uri;
                            } else {
                                window.location.href = '/blackboard/#';
                            }
                        }
                    });
                } else {
                    $.toast({
                        text: "Fails in submit new milestone version",
                        icon: "error",
                        hideAfter: false,
                        position: "bottom-center",
                    });
                }
            }
        );
    });
});