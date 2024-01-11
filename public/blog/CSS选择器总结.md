点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({ top: 0, behavior: 'smooth' });" style="background-color: transparent; position: fixed; bottom: 20px; right: 40px; padding: 10px 10px 5px 10px; cursor: pointer; z-index: 10; border-radius: 13%; box-shadow: 0.5px 3px 7px rgba(0, 0, 0, 0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color: transparent; width: 30px;"></div>

# CSS选择器总结

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/css.png" width="30%" alt="maven" />
</div>

## 1. 类别选择器

类选择器根据类名来选择，前面以”.”来标志：
```css
.demoDiv{
color:#FF0000;
}
```

```html
<div class="demoDiv">
```
## 2. 标签选择器

完整的HTML页面是有很多不同的标签组成，而标签选择器，则是决定哪些标签采用相应的CSS样式：
```css
p{
font-size:12px;
background:#900;
color:090;
}
```

```html
<p>写入文字</p>
```
## 3. ID选择器

ID 选择器可以为标有特定 ID 的 HTML 元素指定特定的样式。 根据元素ID来选择元素，具有唯一性：
```css
#demoDiv{
color:#FF0000;
}
```
```html
<div id="demoDiv">
这个区域字体颜色为红色
</div>
```

## 4. 后代选择器

后代选择器也称为包含选择器，用来选择特定元素或元素组的后代，将对父元素的选择放在前面，对子元素的选择放在后面，中间加一个空格分开：
```css
.father.child{
color:#0000CC;
}
```
```html
<p class="father">
黑色
<label class="child">蓝色
<b>也是蓝色</b>
</label>
</p>
```
## 5. 子选择器

子选择器与后代选择器的区别，子选择器（child selector）仅是指它的直接后代，或者你可以理解为作用于子元素的第一个后代。而后代选择器是作用于所有子后代元素。后代选择器通过空格来进行选择，而子选择器是通过“>”进行选择：
```css
#links a {color:red;}
#links > a {color:blue;}
```
```html
<p id="links">
<a href="#">Div+CSS教程</a>>
<span><a href="#">CSS布局实例</a></span>
<span><a href="#">CSS2.0教程</a></span>
</p>
```
## 6. 伪类选择器

**伪：我理解为基于元素但是又超越元素之外所提取出的抽象。**

用文档以外的其他条件来应用元素的样式，比如鼠标悬停等。这时候我们就需要用到伪类了：

```css
a:link{
color:#999999;
}
a:visited{
color:#FFFF00;
}
a:hover{
color:#006600;
}
input:focus{
background:#E0F1F5;
}
```
Link表示链接在没有被点击时的样式。Visited表示链接已经被访问时的样式。Hover表示当鼠标悬停在链接上面时的样式。

## 7. 通用选择器

通用选择器用*来表示：
```css
*{
font-size: 12px;
}
```
表示所有的元素的字体大小都是12px。

同时通用选择器还可以和后代选择器组合。例如：
```css
p *{
……
}
```
表示所有p元素后代的所有元素都应用这个样式。

## 8. 群组选择器

当几个元素样式属性一样时，可以共同调用一个声明，元素之间用逗号分隔。如：
```css
p, td, li {
line-height:20px;
color:#c00;
}
#main p, #sider span {
color:#000;
line-height:26px;
}
.main p span {
color:#f60;
}
.text1 h1,#sider h3,.art_title h2 {
font-weight:100;
}
```
## 9. 相邻同胞选择器

我们除了上面的子选择器与后代选择器，我们可能还希望找到兄弟两个当中的一个，如一个标题h1元素后面紧跟了两个段落p元素，我们想定位第一个段落p元素，对它应用样式。我们就可以使用相邻同胞选择器。看下面的代码：
```css
h1 + p {color:blue}
```
```html
<h1>一个非常专业的CSS站点</h1>
<p>Div+CSS教程中，介绍了很多关于CSS网页布局的知识。</p>
<p>CSS布局实例中，有很多与CSS布局有关的案例。</p>
```
我们将会看到第一个段落“Div+CSS教程中，介绍了很多关于CSS网页布局的知识。”文字颜色将会是蓝色。而第二段则不受此CSS样式的影响。

## 10. 属性选择器

可以用判断html标签的某个属性是否存在的方法来定义css。属性选择器，是根据元素的属性来匹配的，其属性可以是标准属性也可以是自定义属性，当然，也可以同时匹配多个属性：
```css
　　[attr]
　　[title] {margin-left: 10px}
　　/*选择具有 title 属性的所有元素；*/
　　
　　[attr=val]
　　[title = 'this'] {margin-right: 10px}
　　/* 选择属性 title 的值等于 this 的所有元素 */
　　
　　[attr^=val]
　　[title ^= 'this'] {margin-left: 15px}
　　/*选择属性title的值以this开头的所有元素*/
　　
　　[attr$=val]
　　[title $= 'this'] {margin-right: 15px}
　　/*选择属性title的值以this结尾的所有元素*/
　　
　　[attr*=val]
　　[title *= 'this'] {margin: 10px}
　　/*选择属性title 的值包含 this 的所有元素*/
　　
　　[attr~=val]
　　[title ~= 'this'] {margin-top: 10px}
　　/*选择属性 title 的值包含一个以空格分隔的词为 this 的所有元素，即 title 的值里必须要有 this 这个单词并且this要与其他单词之间有空格分隔*/
　　[attr|=val]
　　[title |= 'this'] {margin-bottom: 10px}
　　/*选择属性 title 的值等于this，或值以 this- 开头的所有元素*/
```

## 11. 伪元素选择器

所有伪元素选择器都必须放在出现该伪元素的选择器的最后面，也就是说伪元素选择器不能跟任何派生选择器，如：p:first-letter em {} 这就是不合法的，ie6不支持:
```css
    :first-letter，设置块元素首字母样式，行内元素转换成块元素和行内块元素也支持；
　　div p:first-letter {font-size: 20px}
　　//选择div元素里所有的p元素的第一个字母或汉字，如果把块元素转换成行内元素则就不支持了；
　　:first-line，设置第一个文本行样式；
　　.box .main:first-line {color: #f00}
　　//只有部分属性允许first-line：所有font属性、color、所有background属性、word-spacing、letter-spacing、text-decoration、vertical-align、text-transform、line-height
　　:before，设置之前的样式，可以插入生成的内容，并设置其样式；
　　body:before {content: 'The Start:'; display: block}
　　//在body元素前插入文本内容'The Start:'，并设置其为块元素
　　:after，设置之后的样式，可以插入生成的内容，并设置其样式；
　　body:after {content: 'The End.'; display: block}
　　//在body元素最后插入文本内容'The End.'，并设置其为块元素
```
## 12. 结构性伪类选择器

HTML CODE:
```html
　　1.<div class="box">
　　2. <span>First span</span>
　　3. <p class="ft">First p</p>
　　4. <div>First div<strong class="ft">Strong text</strong></div>
　　5. <p class="ft">Second p</p>
　　6. <div class="ft">Second div <span>Second span</span><span>Third span</span></div>
　　7.</div>
```
结构性伪类选择器的冒号前边可以跟一个其他选择器做为限定；
带括号的选择器，里面一定要有参数；
匹配子元素，同时也会匹配孙子元素，因为子元素是孙子元素的父元素；
```css
    :first-of-type，选择相对父元素里同类型子元素中的第一个，!lte8
　　.box :first-of-type {color: #f00}
　　//匹配2.3.4以及4里面的strong和6里面的第一个span，因为这个span是6里的第一个span子元素
　　.box .ft:first-of-type {color: #ff0}
　　//匹配3和4里面的strong，因为3是box里面的第一个p且class=”ft”，而4里只有一个strong且class=”ft”，而5和6虽然class=”ft”但是他们相对于box的同类型中不是第一个出现的；
:last-of-type，选择相对父元素里同类型子元素中的最后一个，!lte8
　　.box :last-of-type {color: #f00}
　　//匹配2.5.6以及4里的strong和6里的最后一个span
:only-of-type，选择相对父元素里同类型子元素中只有一个的元素，!lte8
　　.box :only-of-type {color: #f00}
　　//匹配2以及4里的strong，类为box里同类型元素只有一个的只有span
　　.box .ft:only-of-type {color: #f00}
　　//只匹配4里的strong
　　:only-child，选择的元素相对于其父元素是唯一的子元素，!lte8
　　.box :only-child {color: #f00}
　　//只匹配4里的strong
:nth-child(n)，选择其父元素的第n个子元素或多个子元素，索引从1开始，当n用于表达式时索引从0开始!lte8
　　.box :nth-child(3) {color: #f00}
　　//匹配第三个子元素即这里的4
　　.box :nth-child(odd) {color: #f00} 等价于 .box :nth-child(2n + 1) {color: #f00}
　　//匹配奇数即这里的2.4.6以及4里的strong和6里的第一个span
　　.box :nth-child(even) {color: #f00} 等价于 .box :nth-child(2n + 2) {color: #f00}和.box :nth-child(2n)
　　//匹配偶数即这里的3.5以及6里的第二个span
　　.box :nth-child(n + 1) {color: #f00}
　　//匹配 n + 1开始的所有子元素即.box里所有的子元素以及子孙元素，因为这里n是从1开始的即：
　　n = 0 ----> n + 1 = 0 + 1 = 1，即这里的2
　　n = 1 ----> n + 1 = 1 + 1 = 2，即这里的3
　　... ...
　　n = 4 ----> n + 1 = 4 + 1 = 5，即这里的6
:nth-last-child(n)，跟:nth-child(n)使用类似，只是索引是从最后开始往前数，!lte8
　　.box :nth-last-child(3) {color: #f00}
　　//匹配倒数第三个子元素即这里的4
:nth-of-type(n)，选择父元素的第n个或多个同类型的子元素，!lte8
　　.box :nth-of-type(2) {color: #f00}
　　//匹配5和6以及6里面的第二个span
:nth-last-of-type(n)，同上，只是从最后开始往前数，!lte8
　　.box :nth-last-of-type(2) {color: #f00}
　　//匹配3和4以及6里面的第一个span
　　
　　:first-child，选择父元素里的第一个子元素，!ie6
　　.box :first-child {color: #f00}
　　//匹配2和4里的strong以及6里的第一个span
:last-child，选择父元素里的最后一个子元素，!lte8
　　.box :last-child {color: #f00}
　　//匹配6和6里的最后一个span以及4里的strong
:root，选择文档的根元素，在HTML中就是指<html>标签，!lte8
　　:empty，选择没有任何内容的元素，那怕是有一个空格也不行，!lte8
　　table td:empty {background-color: #ffc}
　　//匹配表格里没有内容的td
　　:target，选择当前活动的元素，指锚点点击后跳转到的那一个元素，!lte8
　　:not(selector)，选择排除selector以外的其他所有元素，!lte8
　　.box *:not(div) {background-color: #ffc}
　　//选择box里除div以外的所有后代元素，如果div里有其他非div元素，也会选择上，如上的HTML CODE就会选择上div里面的span和strong
```

## 13. UI元素状态伪类选择器

```css
:enabled，指定元素处于可用状态时的样式，一般用于input，select和textarea
:disabled，指定元素处于不可用状态时的样式，一般用于input，select和textarea
:read-only，指定元素为只读状态时的样式，FF为-moz-read-only，一般用于input和textarea
:read-write，指定元素为只可写状态时的样式，FF为-moz-read-write，一般用于input和textarea
:checked，指定元素被选中状态时的样式，FF为-moz-checked一般用于checkbox和radio
:default，指定元素默认选中的样式，一般用于checkbox和radio
:indeterminate，指定默认一组单选或复选都没被选中的样式，只要有一个被选中则样式被取消，一般用于checkbox和radio
::selection，指定元素处理选中状态时的样式，可用于所有元素，上面的几个基本上只用于表单元素；
FF为::-moz-selection，不能用群组选择器来写；
::selection {background-color: #ffc; color: #fff}
::-moz-selection {background-color: #ffc; color: #fff}
```