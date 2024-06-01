点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# <img src="https://percheung.github.io/blogImg/Python.png" width="50px" alt="C" />python聊天室

> 下面是一个简单的示例，包含了chat_client.py和chat_server.py的代码。

## chat_server

> chat_server.py监听指定的端口，并接收来自客户端的消息，并将消息广播给所有连接到服务器的客户端

```python
import socket
import threading

def handle_client(client_socket, client_address):
    # 处理每个客户端的函数
    while True:
        try:
            data = client_socket.recv(1024).decode('utf-8')  # 接收客户端发送的消息
            if not data:
                break
            print(f'{client_address[0]}:{client_address[1]} - {data}')  # 打印客户端的消息和地址
            broadcast_data = f'{client_address[0]}:{client_address[1]} - {data}'.encode('utf-8')  # 准备广播消息
            broadcast(broadcast_data, client_socket)  # 广播消息给所有连接到服务器的客户端
        except:
            break

    client_socket.close()

def broadcast(message, sender_socket):
    # 广播消息给所有连接到服务器的客户端
    for client in clients:
        if client != sender_socket:
            client.send(message)

def start_server(port):
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.bind(('0.0.0.0', port))  # 绑定服务器的 IP 地址和端口号
    server_socket.listen(10)  # 监听客户端连接，允许最多 5 个同时连接
    print(f'服务器已启动，监听端口 {port}')

    while True:
        client_socket, client_address = server_socket.accept()  # 接受客户端的连接请求
        clients.append(client_socket)  # 将客户端添加到客户端列表
        print(f'客户端已连接：{client_address[0]}:{client_address[1]}')
        client_thread = threading.Thread(target=handle_client, args=(client_socket, client_address))
        client_thread.start()  # 启动一个线程处理客户端

clients = []

if __name__ == '__main__':
    port = int(input('请输入端口号：'))
    start_server(port)
```

## chat_client

> chat_client.py连接到服务器，并允许用户发送消息

```python
import socket
import threading

def receive_messages(client_socket):
    # 接收服务器发送的消息
    while True:
        try:
            data = client_socket.recv(1024).decode('utf-8')  # 接收消息
            print(data)  # 打印消息
        except:
            break

    client_socket.close()

def send_message(client_socket):
    # 发送消息给服务器
    while True:
        message = input()  # 输入消息
        client_socket.send(message.encode('utf-8'))  # 发送消息到服务器

    client_socket.close()

def start_client(server_ip, server_port):
    client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    try:
        client_socket.connect((server_ip, server_port))  # 连接到服务器
        print(f'已连接到 {server_ip}:{server_port}')
        receive_thread = threading.Thread(target=receive_messages, args=(client_socket,))
        receive_thread.start()  # 启动一个线程接收服务器的消息
        send_thread = threading.Thread(target=send_message, args=(client_socket,))
        send_thread.start()  # 启动一个线程发送消息
    except:
        print('连接失败。')

if __name__ == '__main__':
    server_ip = input('请输入服务器 IP 地址：')
    server_port = int(input('请输入服务器端口号：'))
    start_client(server_ip, server_port)
```

