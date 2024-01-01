# 初步领悟C指针

## 1.不是很难，先看指针结构

首先，我们看一下指针的结构。

> type *var-name;

这就是C语言的指针。怎么理解他呢？type代表类型，也就是说type为int，float之类的等等。然后往后看*，*在C语言里专业叫法为取消引用运算符，最后一个单词是var-name，就是变量名的意思。OK，一个指针变量的中文翻译结构如下：

> 数据类型 取消引用运算符 变量名;

理解误区：我想，很多人第一次见到这句int *p时，常常以为这是两部分组成——int 和 *p，实际上恰恰相反，它是三部分组成——int， * ，p。甚至你可以理解为两部分——int * 和p。int *即为整型指针的意思，p代表变量，连起来就是整型指针变量。
## 2.现在唯一的问题是`*`应该怎么理解

既然*叫取消引用运算符，我给你解释一下这句汉语，就是说本来引用了一个东西，现在不引用了。这个引用你可以理解成锁的意思，本来有把锁，现在没了。如此说来是否还是抽象？下面放入C指针代码。

```c
#include <stdio.h>
int main(){
    int number = 5;
    int p = *(&number);
    printf("p的值:%d\n",p);
    printf("number的值:%d\n",number);
}
```
运行后结果如下：

```c
[Running] cd "j:\Project\Day2\" && gcc var.c -o var && "j:\Project\Day2\"var
p的值:5
number的值:5

[Done] exited with code=15 in 1.008 seconds
```
现在注意看代码第四行，我们的p首先不是指针变量，只是一个单纯的变量，我们给他赋值为*(&number)。它最后输出结果即为number的值。
### 2.1 一个简单的问题，`&`是什么？

&为取地址运算符。取地址运算符怎么理解呢？我简单来说，一个数据放入计算机，只有两个属性，一个是它的地址，一个是它本身的值。世界上任何一个计算机，本质上只是处理数据的工具，而计算机数据其实只有两种，一个是数据本身，一个是数据地址，数据地址的本质其实也就是数据。
### 2.2 点破指针的最后一层盔甲

```c
#include <stdio.h>
int main(){
    int number = 5;
    int * p = &number;
    printf("p的值:%d\n",p);
    printf("&number的值:%d\n",&number);
    printf("number的地址:%d\n",&number);
}
```
这个代码，我只是把*这个符号从&number的身上去掉，放到p的身上。

运算结果如下：

```c
[Running] cd "j:\Project\Day2\" && gcc var.c -o var && "j:\Project\Day2\"var
p的值:6487572
&number的值:6487572
number的地址:6487572

[Done] exited with code=24 in 0.729 seconds
```
你可以看到，p等于的就是number的地址位置，&你可以理解为锁，这个锁上写着number的门牌号，&number即为number的地址6487572。
### 2.3 最后一击

运算代码：

```c
#include <stdio.h>
int main(){
    int number = 5;
    int * p = &number;
    printf("p的值:%d\n",p);
    printf("&number的值:%d\n",&number);
    printf("number的地址:%d\n",&number);
    printf("解锁&number:%d\n",*p);
}
```
结果如下：

```c
[Running] cd "j:\Project\Day2\" && gcc var.c -o var && "j:\Project\Day2\"var
p的值:6487572
&number的值:6487572
number的地址:6487572
解锁&number:5

[Done] exited with code=16 in 0.511 seconds
```
**至此，对于指针的讲解全部完成。指针只是解锁地址，拿到值地址下值本身的钥匙，&只是取得地址。**