$(function () {
    $("#txtName").focus();//输入焦点
    $("#pass").keydown(function (event) {
        if (event.which == "13") {//回车键,移动光标到密码框
            $("#pass").focus();
        }
    });
    $("#pass").keydown(function (event) {
        if (event.which == "13") {//回车键，用.ajax提交表单
            $("#btnLogin").trigger("click");
        }
    });
    $("#btnLogin").click(function () { //“登录”按钮单击事件
        //alert(1111)
        //获取用户名称
        var strTxtName = $("#username").val();
        if(strTxtName.length <= 0){
            alert('用户名不能为空');
        }

        var strTxtPass = $("#pass").val();
        if(strTxtPass.length <= 0){
            alert('密码不能为空');
        }
        //获取输入密码
        let sendData ={
            tel: strTxtName,
            password: strTxtPass
        }
        //开始发送数据
        $.ajax({ //请求登录处理页
            url: "http://www.gotit.top/api/user/login", //登录处理页
            //dataType: "html",
            type:'post',
            //contentType: "application/json; charset=utf-8",
            //传送请求数据
            data: `{"tel":"17713575249","password":"123456"}`,
            data: sendData,
            dataType: 'json',
            success: function (res) { //登录成功后返回的数据
                //根据返回值进行状态显示
                if (res.error !=1) {//注意是True,不是true
                    alert('登陆成功');
                   window.location.href=`main.html`;
                }
                else {
                    alert(res.message);
                    //$("#divError").show().html("用户名或密码错误！" + strValue);
                }
            },
            error: function (XMLHttpRequest, textStatus, thrownError) {
                alert('登陆失败，请稍后再试');
            }
        })
    })
});


//一个公用的ajax模板
function loadKind() {
    let sendData = {

    };

    $.ajax({
        type: 'get',
        url : '/api/species',
        //url: '/api/one/new/list?p=1',
        data: sendData,
        dataType: 'json',
        success: function (data) {
            if (data) {
                //请求成功
            } else {
                alert(data.msg);
            }
        },
        error: function () {
            //alert("");
        }
    });
}