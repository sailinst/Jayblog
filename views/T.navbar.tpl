{{define "navbar"}}
 <!-- 将导航条由竖变为横 -->

</head>

<body>
<div id="top">
   <a href="https://github.com/sailinst">我的Github网站</a>
   <a href="http://user.qzone.qq.com/919874817/main">我的QQ空间</a> 
   <a href="/myvideo">视频</a>
   {{if .IsLogin}}
   <a href="/login?exit=true">退出</a>
   {{else}}
   <a href="/login">管理员登录</a>
   {{end}}
</div>  
<div id="header" >
 
   <h1> <a class="navbar-brand"  href="/">我的博客</a> </h1>
       <div class="adver0" >
         <ul  class="nav navbar-nav" style="list-style-type:none">
         <li {{ if .IsHome}} class="active" {{ end}}><a href="/">首页</a></li>
         <li {{ if .IsCategory}} class="active" {{ end}}><a href="/category">分类</a></li>
         <li {{ if .IsTopic}} class="active" {{ end}}><a href="/topic">文章</a></li>
         <li {{ if .IsPhotograph}} class="active" {{ end}}><a href="/photograph">图片</a></li>       
         </ul>

       </div>
    
       <div class="adver"> 
        <object type="application/x-shockwave-flash" data="../static/music/DewPlayer/dewplayer-bubble.swf"  width="650" height="65" 
        id="dewplayer-bubble" >
        <param name="wmode" value="transparent">
        <param name="movie" value="../static/music/DewPlayer/dewplayer-bubble-vol.swf" />
        <param name="flashvars" value="mp3=../static/music/DewPlayer/mp3/让一切随风.mp3|../static/music/DewPlayer/mp3/不再犹豫.mp3|../static/music/DewPlayer/mp3/光辉岁月.mp3|../static/music/DewPlayer/mp3/test2.mp3&autostart=0&showtime=1&randomplay=1&volume=80&autoreplay=1" />
        </object>
      
       </div>

 
</div>
 
{{end}}