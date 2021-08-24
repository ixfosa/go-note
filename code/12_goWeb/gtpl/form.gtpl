
<html>
<head>
    <title>parameter</title>
</head>
<body>
<form action="/parameter" method="post">
    用户:<input type="text" name="username"> <br>
    密码:<input type="password" name="password"> <br>
    性别:<input type="radio" name="sex" value="女">女
    <input type="radio" name="sex" value="男">男  <br>
    爱好:<input type="checkbox" name="hobby" value="睡觉">睡觉
    <input type="checkbox" name="hobby" value="吃饭">吃饭
    <input type="checkbox" name="hobby" value="打游戏">打游戏 <br>
    城市:<select name="city">
        <option value="南昌">南昌</option>
        <option value="九江</">九江</option>
        <option value="瑞昌">瑞昌</option>
        <option value="夏畈">夏畈</option>
    </select> <br><br>
    简介:<textarea name="jianjie" rows="8" cols="20"></textarea><br><br>
    <input type="submit" value="登录">
</form>
</body>
</html>