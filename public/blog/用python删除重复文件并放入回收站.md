点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# 用<img src="https://percheung.github.io/blogImg/Python.png" width="50px" alt="C" />python删除重复文件并放入回收站

## 1.需求

执行删除当前文件夹（就一层，没有子文件夹）里的重复文件，将文件移动到回收站。

## 2.实现

可以使用第三方库`send2trash`来实现。使用命令安装该库。

```bash
pip install send2trash
```

python代码如下。

```python
import os
import hashlib
from send2trash import send2trash

def calculate_hash(file_path):
    with open(file_path, 'rb') as f:
        content = f.read()
        file_hash = hashlib.md5(content).hexdigest()
        return file_hash

def delete_duplicate_files():
    current_folder = os.getcwd()  # 获取当前文件夹路径
    
    hash_dict = {}
    
    for file_name in os.listdir(current_folder):
        file_path = os.path.join(current_folder, file_name)
        
        if os.path.isfile(file_path):
            file_hash = calculate_hash(file_path)
            
            if file_hash in hash_dict:
                print('删除重复文件:', file_path)
                send2trash(file_path)  # 移动文件到回收站
            else:
                hash_dict[file_hash] = file_path

delete_duplicate_files()
```

在上述代码中，我们导入了`send2trash`库，并使用`send2trash(file_path)`函数来将重复文件移动到回收站。

使用`send2trash`库可以更安全地删除文件，因为它提供了回收站的保护，可以在需要时还原已删除的文件。这样，你可以在删除文件后从回收站中恢复它们，而不是永久删除。

## 3.使用

命名python文件为`删除重复文件.py`，只需要在需要删除重复文件的目录下执行以下命令，就会去重并放入回收站。

```bash
python 删除重复文件.py
```

