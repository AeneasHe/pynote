pynote 派记: 一个markdown笔记管理软件
=================================

## 派记开发目标和原则
- 内容和样式分离
    - 针对office系列
    - 采用html,css为主，json,xml为辅的文档体系
    - 内容必须要便于人编辑修改阅读，同时便于软件批量处理，机器大数据处理
    - 样式必需便于切换，常见文档默认提供数款样式，其他样式由开发者自行决定收费模式

- 内容所有权属于用户，样式所有权属于派记及生态伙伴
    - 用户在计算机本地资源浏览器中，派记客户端中，派记云端中，使用体验一致
    - 笔记默认不上传云端，只在用户分享和主动上传云端时才会同步到云端
    - 支持在线托管和发布，类似gitbook
    - 支持内容的存证（hash）保存到区块链，类似电子签章功能
    - 支持样式以nft的形式发布到区块链
    - 支持有限的内容发布到区块链

- 支持多种常见开源协议文档
    - markdown
    - 思维导图
    - 示意图（mermaid的协议方案）

- 走mongodb的开源发展路径

## 说明
- 需要用户已经安装了 Chrome/Chromium >= 70
- 目前过滤了其他格式的文档，只支持.md结尾的markdown
- 目前还没有样式
- 目前只能浏览不能编辑
- 目前目录遍历采用深度优先遍历，后面需要改成广度优先

## 安装
dist文件夹下有mac安装包，拖访达的应用程序即可

## 开发启动
- 配置文件  
跟路径下的 config.json  
默认初始目录RootPath为: "/Users/aeneas/Github/Cofepy/youdao"  

- 启动
``` shell
go run main.go
```

## 打包发布
发布时需要将main.go按如下修改

开发时： 
``` golang
//  main.go

go server.StartFromTemplate(configPath) //从template中读取静态文件
//go server.StartFromAsset() //从嵌入式asset读取静态文件
```
部署时： 
``` golang
//  main.go

//go server.StartFromTemplate(configPath) //从template中读取静态文件
go server.StartFromAsset() //从嵌入式asset读取静态文件
```

打包脚本,目前只验证过mac
```
gen.sh              将template文件夹下的内容生成嵌入式文件
build-macos.sh      mac
build-linux.sh      linux
build-windows.bat   windows
```

打包以后mac下的配置文件在,修改对应配置即可:  
dist/mac/Pynote/app/Contents/MacOS/config.json  