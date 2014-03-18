ueditor4go
==========

刚开始学golang，在开发项目的时候遇到编辑器问题，然后又没找到一个golang的在线编辑器，就写了个支持go的ueditor
获取远程图片及涂鸦部分没有完成，其他功能都已实现，不影响大部分使用。
controller中暂时没有session验证，有需要的可以自己加入

使用方法：

1.在你页面head内，引用ueditor文件、配置文件和语言包文件。插入下面的代码，修改引用路文件的路径。
<!-- 配置文件 -->
<script type="text/javascript" src="/ueditor/ueditor.config.js"></script>
<!-- 编辑器源码文件 -->
<script type="text/javascript" src="/ueditor/ueditor.all.js"></script>
<!-- 语言包文件(建议手动加载语言包，避免在ie下，因为加载语言失败导致编辑器加载失败) -->
<script type="text/javascript" src="/ueditor/lang/zh-cn/zh-cn.js"></script>

2.然后在代码文件中设置编辑器容器，并添加编辑器的实例化代码。具体代码示例如下
<textarea id="content" name="content">{{.Article.Content}}</textarea>
<script type="text/javascript">
    var editor = UE.getEditor('content')
</script>

3.配置好ueditor目录下的ueditor.config.js的参数

4.实现controller，对上传模块的处理（见controller中的ue.go，ue.go中一些存储目录和网址信息需要根据自身情况填写）

5.实现路由（这里因为是基于beego的，路由信息在routers下的router.go中）

欢迎交流！
