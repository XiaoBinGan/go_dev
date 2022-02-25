<!DOCTYPE html>
<html>
<head>
    <!-- set `maximum-scale` for some compatibility issues -->
    <meta
            name="viewport"
            content="width=device-width, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no"
    />
    <meta name="msapplication-tap-highlight" content="no" />
    <meta name="format-detection" content="telephone=no, email=no" />

</head>

<body>
<div id="app">
{{/*  这里注意的是自定义函数的的key一定要在前面  这样后面的.参数才能正常的被传入到hobby自定义函数当中去  */}}
   My name is {{hobby . }}
</div>
</body>
</html>
