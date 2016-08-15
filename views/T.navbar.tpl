{{define "navbar"}}
 <!-- 将导航条由竖变为横 -->
 <a class="navbar-brand" href="/">我的博客</a>
<div>
    <ul class="nav navbar-nav">
         <li {{ if .IsHome}} class="active" {{ end}}><a href="/">首页</a></li>
         <li {{ if .IsCategory}} class="active" {{ end}}><a href="/category">分类</a></li>
         <li {{ if .IsTopic}} class="active" {{ end}}><a href="/topic">文章</a></li>
         <li {{ if .IsPhotograph}} class="active" {{ end}}><a href="/photograph">图片</a></li>       
    </ul>
 </div>

 <div class="pull-right">
     <ul class="nav navbar-nav">
     {{if .IsLogin}}
     <li><a href="/login?exit=true">退出登录</a></li> 
     <li><a href="/signup">注册新用户</a></li> 
     {{else}}
     <li><a href="/login">管理员登录</a></li>   
     <li><a href="/signup">注册新用户</a></li>  	
     {{end}}
     </ul>
 	 
 </div>

{{end}}