<html>
<head>
    <title></title>
</head>
<body>
<form action="/formRepeat" method="post">
    用户名:<input type="text" name="username">
    密码:<input type="password" name="password">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="登录">
</form>
</body>
</html>