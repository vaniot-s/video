$(document).ready(function() {
   //
    DEFAULT_COOKIE_EXPIRE_TIME = 30;

    uname = '';
    session = '';
    uid = 0;
    currentVideo = null;//当前video
    listedVideos = null;//videolist

    //信息
    session = getCookie('session');
    uname = getCookie('username');


    // home page event registry
    $("#regbtn").on('click', function(e) {
        $("#regbtn").text('Loading...')
        e.preventDefault()
        registerUser(function(res, err) {
            if (err != null) {
                $('#regbtn').text("Register")
                popupErrorMsg('encounter an error, pls check your username or pwd');
                return;
            }

            var obj = JSON.parse(res);
            setCookie("session", obj["session_id"], DEFAULT_COOKIE_EXPIRE_TIME);
            setCookie("username", uname, DEFAULT_COOKIE_EXPIRE_TIME);
            $("#regsubmit").submit();
        });
    });

    $("#siginbtn").on('click', function(e) {

        $("#siginbtn").text('Loading...')
        e.preventDefault();
        signinUser(function(res, err) {
            if (err != null) {
                $('#siginbtn').text("Sign In");
                //window.alert('encounter an error, pls check your username or pwd')
                popupErrorMsg('encounter an error, pls check your username or pwd');
                return;
            }

            var obj = JSON.parse(res);
            setCookie("session", obj["session_id"], DEFAULT_COOKIE_EXPIRE_TIME);
            setCookie("username", uname, DEFAULT_COOKIE_EXPIRE_TIME);
            $("#siginsubmit").submit();
        });
    });

    $("#signinhref").on('click', function() {
        $("#regsubmit").hide();
        $("#siginsubmit").show();
    });

    $("#registerhref").on('click', function() {
        $("#regsubmit").show();
        $("#siginsubmit").hide();
    });

});


function setCookie(cname, cvalue, exmin) {
    var d = new Date();
    d.setTime(d.getTime() + (exmin * 60 * 1000));
    var expires = "expires="+d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

function getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for(var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

// Async ajax methods
// User operations
function registerUser(callback) {
    var username = $("#username").val();
    var pwd = $("#pwd").val();
    var apiUrl = window.location.hostname + '/api';

    if (username == '' || pwd == '') {
        callback(null, err);
    }

    var reqBody = {
        'user_name': username,
        'pwd': pwd
    }

    var dat = {
        'url': 'http://'+ window.location.hostname + ':9003/user',
        'method': 'POST',
        'req_body': JSON.stringify(reqBody)
    };




    $.ajax({
        url  : 'http://' + window.location.hostname + '/api',
        type : 'post',
        data : JSON.stringify(dat),
        statusCode: {
            500: function() {
                callback(null, "internal error");
            }
        },
        complete: function(xhr, textStatus) {
            if (xhr.status >= 400) {
                callback(null, "Error of Signin");
                return;
            }
        }
    }).done(function(data, statusText, xhr){
        if (xhr.status >= 400) {
            callback(null, "Error of register");
            return;
        }

        uname = username;
        callback(data, null);
    });
}

function signinUser(callback) {
    var username = $("#susername").val();
    var pwd = $("#spwd").val();
    var apiUrl = window.location.hostname + '/api';

    if (username == '' || pwd == '') {
        callback(null, err);
    }

    var reqBody = {
        'user_name': username,
        'pwd': pwd
    }

    var dat = {
        'url': 'http://'+ window.location.hostname + ':9003/user/' + username,
        'method': 'POST',
        'req_body': JSON.stringify(reqBody)
    };

    $.ajax({
        url  : 'http://' + window.location.hostname + '/api',
        type : 'post',
        data : JSON.stringify(dat),
        statusCode: {
            500: function() {
                callback(null, "Internal error");
            }
        },
        complete: function(xhr, textStatus) {
            if (xhr.status >= 400) {
                callback(null, "Error of Signin");
                return;
            }
        }
    }).done(function(data, statusText, xhr){
        if (xhr.status >= 400) {
            callback(null, "Error of Signin");
            return;
        }
        uname = username;

        callback(data, null);
    });
}



