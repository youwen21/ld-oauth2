package ginresp

var pageRefresh = `
<!DOCTYPE HTML>
<html>
<head>
<meta http-equiv="Content-Type" context="text/html;charset=utf8"/>
<meta http-equiv="Refresh" content="%v;url=%s"/>
<title>%s</title>
</head>

<body>
<h3>%s</h3>

<p>%v秒后页面跳转。</p>

</body>
</html>
`

var pageErr = `
<!DOCTYPE HTML>
<html>
<head>
<meta http-equiv="Content-Type" context="text/html;charset=utf8"/>
<title>%s</title>
</head>

<body>
<h3>%s</h3>

</body>
</html>
`
