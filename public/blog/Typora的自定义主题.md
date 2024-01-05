点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}

# Typora的自定义主题

> 修改我们的Typora主题。这个主题是苹果味的，主要用到的字体为PingFang SC，Times New Roman，Microsoft YaHei Mono。

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/Typora.png" width="30%" alt="Typora" />
</div>

## 1.加入my.css

> 在文件夹`C:\Users\Peter\AppData\Roaming\Typora\themes`中，加入下面的css文件即可，命名为my.css。

**`注意：Typora的默认路径是AppData\Roaming\Typora\themes下，\Users\Peter\部分会因为你用户名不同而不同。`**

### my.css内容

```css
:root {
    --side-bar-bg-color: #fafafa;
    --control-text-color: #777;
}

html {
    font-size: 16px;
    -webkit-font-smoothing: antialiased;
	font-family: 'Times New Roman','PingFang SC';
}

body {
    font-family: 'Times New Roman','PingFang SC';
    color: rgb(51, 51, 51);
    line-height: 1.6;
}

#write {
    max-width: 860px;
  	margin: 0 auto;
  	padding: 30px;
    padding-bottom: 100px;
	font-family: 'Times New Roman','PingFang SC';
}

@media only screen and (min-width: 1400px) {
	#write {
		max-width: 1024px;
		font-family: 'Times New Roman','PingFang SC';
	}
}

@media only screen and (min-width: 1800px) {
	#write {
		max-width: 1200px;
		font-family: 'Times New Roman','PingFang SC';
	}
}

#write > ul:first-child,
#write > ol:first-child{
    margin-top: 30px;
	font-family: 'Times New Roman','PingFang SC';
}

a {
    color: #4183C4;
	font-family: 'PingFang SC';
}
h1,
h2,
h3,
h4,
h5,
h6 {
    position: relative;
    margin-top: 1rem;
    margin-bottom: 1rem;
    font-weight: bold;
    line-height: 1.4;
    cursor: text;
	font-family: 'Times New Roman','PingFang SC';
}
h1:hover a.anchor,
h2:hover a.anchor,
h3:hover a.anchor,
h4:hover a.anchor,
h5:hover a.anchor,
h6:hover a.anchor {
    text-decoration: none;
	font-family: 'Times New Roman','PingFang SC';
}
h1 tt,
h1 code {
    font-size: inherit;
	font-family: 'PingFang SC';
}
h2 tt,
h2 code {
    font-size: inherit;
	font-family: 'PingFang SC';
}
h3 tt,
h3 code {
    font-size: inherit;
	font-family: 'PingFang SC';
}
h4 tt,
h4 code {
    font-size: inherit;
	font-family: 'PingFang SC';
}
h5 tt,
h5 code {
    font-size: inherit;
	font-family: 'PingFang SC';
}
h6 tt,
h6 code {
    font-size: inherit;
	font-family: 'PingFang SC';
}
h1 {
    font-size: 2.25em;
    line-height: 1.2;
    border-bottom: 1px solid #eee;
	font-family: 'PingFang SC';
}
h2 {
    font-size: 1.75em;
    line-height: 1.225;
    border-bottom: 1px solid #eee;
	font-family: 'PingFang SC';
}

/*@media print {
    .typora-export h1,
    .typora-export h2 {
        border-bottom: none;
        padding-bottom: initial;
    }

    .typora-export h1::after,
    .typora-export h2::after {
        content: "";
        display: block;
        height: 100px;
        margin-top: -96px;
        border-top: 1px solid #eee;
    }
}*/

h3 {
    font-size: 1.5em;
    line-height: 1.43;
	font-family: 'PingFang SC';
}
h4 {
    font-size: 1.25em;
	font-family: 'PingFang SC';
}
h5 {
    font-size: 1em;
	font-family: 'PingFang SC';
}
h6 {
   font-size: 1em;
    color: #777;
	font-family: 'PingFang SC';
}
p,
blockquote,
ul,
ol,
dl,
table{
    margin: 0.8em 0;
	font-family: 'Times New Roman','PingFang SC';
}
li>ol,
li>ul {
    margin: 0 0;
	font-family: 'Times New Roman','PingFang SC';
}
hr {
    height: 2px;
    padding: 0;
    margin: 16px 0;
    background-color: #e7e7e7;
    border: 0 none;
    overflow: hidden;
    box-sizing: content-box;
	font-family: 'Times New Roman','PingFang SC';
}

li p.first {
    display: inline-block;
	font-family: 'Times New Roman','PingFang SC';
}
ul,
ol {
    padding-left: 30px;
	font-family: 'Times New Roman','PingFang SC';
}
ul:first-child,
ol:first-child {
    margin-top: 0;
	font-family: 'Times New Roman','PingFang SC';
}
ul:last-child,
ol:last-child {
    margin-bottom: 0;
	font-family: 'Times New Roman','PingFang SC';
}
blockquote {
    border-left: 4px solid #dfe2e5;
    padding: 0 15px;
    color: #777777;
	font-family: 'Times New Roman','PingFang SC';
}
blockquote blockquote {
    padding-right: 0;
	font-family: 'Times New Roman','PingFang SC';
}
table {
    padding: 0;
	font-family: 'Times New Roman','PingFang SC';
    word-break: initial;
}
table tr {
    border: 1px solid #dfe2e5;
	font-family: 'Times New Roman','PingFang SC';
    margin: 0;
    padding: 0;
}
table tr:nth-child(2n),
thead {
    background-color: #f8f8f8;
	font-family: 'Times New Roman','PingFang SC';
}
table th {
    font-weight: bold;
    border: 1px solid #dfe2e5;
    border-bottom: 0;
    margin: 0;
    padding: 6px 13px;
	font-family: 'Times New Roman','PingFang SC';
}
table td {
    border: 1px solid #dfe2e5;
    margin: 0;
    padding: 6px 13px;
	font-family: 'Times New Roman','PingFang SC';
}
table th:first-child,
table td:first-child {
    margin-top: 0;
	font-family: 'Times New Roman','PingFang SC';
}
table th:last-child,
table td:last-child {
    margin-bottom: 0;
	font-family: 'Times New Roman','PingFang SC';
}

.CodeMirror-lines {
    padding-left: 4px;
	font-family: 'Microsoft YaHei Mono';
}

.code-tooltip {
    box-shadow: 0 1px 1px 0 rgba(0,28,36,.3);
    border-top: 1px solid #eef2f2;
	font-family: 'Times New Roman','PingFang SC';
}

.md-fences,
code,
tt {
    border: 1px solid #e7eaed;
    background-color: #f8f8f8;
    border-radius: 3px;
    padding: 0;
    padding: 2px 4px 0px 4px;
    font-size: 0.9em;
	font-family: 'Times New Roman','PingFang SC';
}

code {
    background-color: #f3f4f4;
    padding: 0 2px 0 2px;
	font-family: 'Times New Roman','PingFang SC';
}

.md-fences {
    margin-bottom: 15px;
    margin-top: 15px;
    padding-top: 8px;
    padding-bottom: 6px;
	font-family: 'Times New Roman','PingFang SC';
}


.md-task-list-item > input {
  margin-left: -1.3em;
  font-family: 'Times New Roman','PingFang SC';
}

@media print {
    html {
        font-size: 13px;
		font-family: 'Times New Roman','PingFang SC';
    }
    table,
    pre {
        page-break-inside: avoid;
		font-family: 'Times New Roman','PingFang SC';
    }
    pre {
        word-wrap: break-word;
		font-family: 'Times New Roman','PingFang SC';
    }
}

.md-fences {
	background-color: #f8f8f8;
	font-family: 'Times New Roman','PingFang SC';
}
#write pre.md-meta-block {
	padding: 1rem;
    font-size: 85%;
    line-height: 1.45;
    background-color: #f7f7f7;
    border: 0;
    border-radius: 3px;
    color: #777777;
    margin-top: 0 !important;
	font-family: 'Times New Roman','PingFang SC';
}

.mathjax-block>.code-tooltip {
	font-family: 'Times New Roman','PingFang SC';
	bottom: .375rem;
}

.md-mathjax-midline {
	font-family: 'Times New Roman','PingFang SC';
    background: #fafafa;
}

#write>h3.md-focus:before{
	font-family: 'Times New Roman','PingFang SC';
	left: -1.5625rem;
	top: .375rem;
}
#write>h4.md-focus:before{
	font-family: 'Times New Roman','PingFang SC';
	left: -1.5625rem;
	top: .285714286rem;
}
#write>h5.md-focus:before{
	font-family: 'Times New Roman','PingFang SC';
	left: -1.5625rem;
	top: .285714286rem;
}
#write>h6.md-focus:before{
	font-family: 'Times New Roman','PingFang SC';
	left: -1.5625rem;
	top: .285714286rem;
}
.md-image>.md-meta {
	font-family: 'Times New Roman','PingFang SC';
    /*border: 1px solid #ddd;*/
    border-radius: 3px;
    padding: 2px 0px 0px 4px;
    font-size: 0.9em;
    color: inherit;
}

.md-tag {
	font-family: 'Times New Roman','PingFang SC';
    color: #a7a7a7;
    opacity: 1;
}

.md-toc { 
	font-family: 'Times New Roman','PingFang SC';
    margin-top:20px;
    padding-bottom:20px;
}

.sidebar-tabs {
	font-family: 'Times New Roman','PingFang SC';
    border-bottom: none;
}

#typora-quick-open {
	font-family: 'Times New Roman','PingFang SC';
    border: 1px solid #ddd;
    background-color: #f8f8f8;
}

#typora-quick-open-item {
	font-family: 'Times New Roman','PingFang SC';
    background-color: #FAFAFA;
    border-color: #FEFEFE #e5e5e5 #e5e5e5 #eee;
    border-style: solid;
    border-width: 1px;
}

/** focus mode */
.on-focus-mode blockquote {
	font-family: 'Times New Roman','PingFang SC';
    border-left-color: rgba(85, 85, 85, 0.12);
}

header, .context-menu, .megamenu-content, footer{
    font-family: 'Times New Roman','PingFang SC';
}

.file-node-content:hover .file-node-icon,
.file-node-content:hover .file-node-open-state{
    visibility: visible;
}

.mac-seamless-mode #typora-sidebar {
    background-color: #fafafa;
    background-color: var(--side-bar-bg-color);
}

.md-lang {
	font-family: 'Times New Roman','PingFang SC';
    color: #b4654d;
}

/*.html-for-mac {
    --item-hover-bg-color: #E6F0FE;
}*/

#md-notification .btn {
	font-family: 'Times New Roman','PingFang SC';
    border: 0;
}

.dropdown-menu .divider {
	font-family: 'Times New Roman','PingFang SC';
    border-color: #e5e5e5;
    opacity: 0.4;
}

.ty-preferences .window-content {
	font-family: 'Times New Roman','PingFang SC';
    background-color: #fafafa;
}

.ty-preferences .nav-group-item.active {
	font-family: 'Times New Roman','PingFang SC';
    color: white;
    background: #999;
}

.menu-item-container a.menu-style-btn {
	font-family: 'Times New Roman','PingFang SC';
    background-color: #f5f8fa;
    background-image: linear-gradient( 180deg , hsla(0, 0%, 100%, 0.8), hsla(0, 0%, 100%, 0)); 
}
```

## 2.选择你的主题为my

注意：如果你的字体没有完全像我一样，是因为你的电脑上没有苹果字体`PingFang SC`，不过你文件里数字和英文一定会变化的，因为大家都有`Times New Roman`字体，至于中文字体的解决办法是你可以将my.css文件中的PingFang SC替换成微软雅黑。至于编程字体我用的`Microsoft YaHei Mono`。如果你希望文件的字体是苹方，想要字体，可以私信我发给你。