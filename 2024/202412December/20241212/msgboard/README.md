# 留言板实战小作业

## 前言

	这个小项目来源于学长布置的作业，实现的为一个简单的留言板功能。

## 正文


### 接口
该项目一共实现了10个接口，包括Register - POST（用户注册）； 2. Login - GET（用户登录）； 3. Publish - POST（发布留言）； 4. Reply - POST（回复留言）； 5. DelUser - DELETE（删除用户）； 6. CloseMsg - PUT（关闭留言板）； 7. GetAllMsg - GET（获取所有留言）； 8. DeleteMsg - DELETE（删除留言）； 9. OpenMsg - PUT（开放留言板）；10. AlterMsg - PUT (修改留言)。

下面对这几个接口进行简述

#### 1. 注册接口

- **请求方式**：POST
- **接口路径**：/register
- **作用**：用户注册
- **请求参数**：
  - `username`: string (必填) - 用户名（登录用）
  - `nickname`: string (必填) - 用户昵称
  - `password`: string (必填) - 用户密码
- **请求示例**：

```json
{
    "username": "111",
    "nickname": "anon酱",
    "password": "123456"
}
```



- **返回参数**：
  - `id`: string - 用户ID
- **状态码**：
  - 200: 注册成功
  - 400: 用户名已存在
  - 502: 注册失败
- **返回示例**：

```json
{
    "code": 200,
    "你的用户id": "123456"
}
```



```json
{
    "code": 400,
    "message": "该用户名已经使用过了~"
}
```



```json
{
    "code": 502,
    "message": "注册失败"
}
```



#### 2. 登录接口

- **请求方式**：GET
- **接口路径**：/login
- **作用**：用户登录
- **请求参数**：
  - `username`: string (必填) - 用户名
  - `password`: string (必填) - 用户密码
- **请求示例**：

```json
{
    "username": "111",
    "password": "123456"
}
```



- **返回参数**：
  - `message`: string - 登录成功消息
- **状态码**：
  - 200: 登录成功
  - 502: 账号不存在或密码错误
- **返回示例**：

```json
{
    "code": 200,
    "message": "登陆成功 It‘s My GO！！！！！"
}
```



```json
{
    "code": 502,
    "message": "账号不存在或是账号密码错了？"
}
```



#### 3. 发布信息接口

- **请求方式**：POST
- **接口路径**：/Publish
- **作用**：发布信息
- **请求参数**：
  - `username`: string (必填) - 用户名
  - `password`: string (必填) - 用户密码
  - `message`: string (必填) - 发布的消息内容
- **请求示例**：

```json
{
    "username": "111",
    "password": "123456",
    "message": "It‘s My Go!!!!!"
}
```



- **返回参数**：
  - `id`: string - 消息ID
- **状态码**：
  - 200: 发布成功
  - 501: 账号密码错误
  - 502: 发布失败
- **返回示例**：

```json
{
    "code": 200,
    "id": "78910",
    "message": "发布成功！"
}
```



```json
{
    "code":    501,
    "message": "账号不存在或密码错误辣~",
}
```



```json
{
    "code": 502,
    "message": "发布消息出问题了~"
}
```



#### 4. 回复信息接口

- **请求方式**：POST
- **接口路径**：/Reply
- **作用**：回复信息
- **请求参数**：
  - `username`: string (必填) - 用户名
  - `password`: string (必填) - 用户密码
  - `message`: string (必填) - 回复的消息内容
  - `parent_id`: string (必填) - 被回复消息的ID
- **请求示例**：

```json
{
    "username": "111",
    "password": "123456",
    "message": "Girls Band Cry~",
    "parent_id": "12345"
}
```



- **返回参数**：
  - `id`: string - 回复的消息ID
- **状态码**：
  - 200: 回复成功
  - 400: 回复的消息已不存在
  - 502: 回复失败
  - 503: 留言板关闭了
- **返回示例**：

```json
{
    "code": 200,
    "id": "0d000721",
    "message": "留言成功！"
}
```



```json
{
    "code": 400,
    "message": "你回复的消息已经不存在了哦~"
}
```



```json
{
    "code": 502,
    "message": "留言失败~"
}
```



```json
{
    "code":    503,
    "message": "看来你来晚了呢，这个此处的回复功能已经被关闭了~",
}
```



#### 5. 关闭消息接口

- **请求方式**：PUT
- **接口路径**：/CloseMsg
- **作用**：关闭消息
- **请求参数**：
  - `username`: string (必填) - 用户名
  - `password`: string (必填) - 用户密码
  - `msg_id`: string (必填) - 要关闭的消息ID
- **请求示例**：

```json
{
    "username": "111",
    "password": "123456",
    "msg_id": "12345"
}
```



- **返回参数**：
  - `message`: string - 关闭结果消息
- **状态码**：
  - 200: 关闭成功
  - 501: 账号不对
  - 502: 关闭失败
  - 503: 消息不存在，或是账号消息不匹配
- **返回示例**：

```json
{
    "code": 200,
    "message": "留言关闭成功！"
}
```



```json
{
    "code":    501,
    "message": "关闭失败，账号密码错误",
}
```



```json
{
    "code": 502,
    "message": "留言关闭失败了~"
}
```



```json
{
    "code":    503,
    "message": "不存在的消息或是用户不匹配！",
}
```



#### 6. 获取所有消息接口

- **请求方式**：GET
- **接口路径**：/GetAllMsg
- **作用**：获取所有消息
- **请求参数**：
  - `username`: string (必填) - 用户名
  - `password`: string (必填) - 用户密码
- **请求示例**：

```json
{
    "username": "111",
    "password": "123456"
}
```



- **返回参数**：
  - `messages`: array - 消息列表
- **状态码**：
  - 200: 获取成功
  - 502: 账号不存在或密码错误
- **返回示例**：

```json
{
    "code": 200,
    "messages": [
        {"id": "1", "content": "我已经只有母鸡卡了"},
        {"id": "2", "content": "不喜欢00的都是神人！！"}
    ]
}
```



```json
{
    "code": 502,
    "message": "账号不存在或密码错误辣~"
}
```



#### 7. 删除消息接口

- **请求方式**：DELETE
- **接口路径**：/DelMsg
- **作用**：删除消息
- **请求参数**：
  - `username`: string (必填) - 用户名
  - `password`: string (必填) - 用户密码
  - `msg_id`: string (必填) - 要删除的消息ID
- **请求示例**：

```json
{
    "username": "111",
    "password": "123456",
    "msg_id": "12345"
}
```



- **返回参数**：
  - `message`: string - 删除结果消息
- **状态码**：
  - 200: 删除成功
  - 500: 消息不存在
  - 501: 账号密码错误
  - 502: 删除失败
- **返回示例**：

```json
{
    "code": 200,
    "message": "你的这句话已经永远消失在留言板上了~"
}
```



```json
{
    "code":    500,
    "message": "不存在的消息！！删除失败！",
}
```



```json
{
    "code":    501,
    "message": "账号不存在或密码错误辣~",
}
```



```json
{
    "code": 502,
    "message": "不存在的消息！！删除失败！"
}
```



#### 8. 打开消息接口

- **请求方式**：PUT
- **接口路径**：/OpenMsg
- **作用**：打开消息
- **请求参数**：
  - `username`: string (必填) - 用户名
  - `password`: string (必填) - 用户密码
  - `msg_id`: string (必填) - 要打开的消息ID
- **请求示例**：

```json
{
    "username": "111",
    "password": "123456",
    "msg_id": "12345"
}
```



- **返回参数**：
  - `message`: string - 打开成功或失败的提示
- **状态码**：
  - 200: 打开成功
  - 501: 用户与留言不匹配
  - 503: 密码错误或账号不存在
  - 502: 打开失败
- **返回示例**：

```json
{
    "code": 200,
    "message": "已经成功打开了留言板，快去通知你的小伙伴来留言吧！"
}
```



```json
{
    "code": 502,
    "message": "开门失败了~"
}
```



```json
{
    "code":    503,
    "message": "账号不存在或密码错误辣~",
}
```



```json
{
    "code":    501,
    "message": "用户不匹配或是你开了一个空门~",
}
```

#### 9. 删除用户接口

- **请求方式**：DELETE
- **接口路径**：/DelUser
- **作用**：删除用户
- **请求参数**：
  - `username`: string (必填) - 用户名
  - `password`: string (必填) - 用户密码
- **请求示例**：

```json
{
    "username": "111",
    "password": "123456"
}
```



- **返回参数**：
  - `message`: string - 删除结果消息
- **状态码**：
  - 200: 删除成功
  - 502: 删除失败
- **返回示例**：

```json
{
    "code": 200,
    "message": "删除成功"
}
```



```json
{
    "code": 502,
    "message": "删除用户失败"
}
```

#### 10. 修改留言接口
- **请求方式**： PUT
- **接口路径**：/AlterMsg
- **作用**：修改留言的信息
- **请求参数**：
  - `username`: string (必填) - 用户名
  - `password`: string (必填) - 用户密码
  - `msg_id`: string(必填) - 修改对应的留言的id
  - `message`: string(必填) - 修改的内容
- **请求示例**：

```json
{
    "username": "111",
    "password": "123456",
    "msg_id": "114514",
    "message": "关注塔菲谢谢喵"
}
```



- **返回参数**：
  - `message`: string - 删除结果消息
- **状态码**：
  - 200: 修改成功
  - 500: 账号或密码不对
  - 501: 修改的留言对象不存在，或者用户与留言不匹配
  - 502: 修改失败
- **返回示例**：

```json
{
    "code": 200,
    "message": "修改成功！"
}
```




```json
{
    "code": 500,
    "message": "你输入的账号不存在或者密码错误了！"
}
```



```json
{
    "code": 501,
    "message": "用户不匹配或者留言已经没了"
}
```



```json
{
    "code": 502,
    "message": "修改失败"
}
```



### MySQL的表格

- **用户信息的表**
```mysql
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY, //用户的标识id
  nickname VARCHAR(255)  //绰号
  username VARCHAR(255) NOT NULL UNIQUE, //用户名，具有唯一性
  password VARCHAR(255) NOT NULL,  //密码
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP, //注册时间
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE 
CURRENT_TIMESTAMP //更新时间
);
```
- **留言信息的表**
```mysql
CREATE TABLE messages (
    id INT PRIMARY KEY AUTO_INCREMENT, //唯一id，自增
    user_id INT NOT NULL,	//发布留言的人的id
    host_id INT NOT NULL,	//当前留言所在的主题的发布人的id
    context TEXT NOT NULL,	//留言内容
    create_at DATETIME DEFAULT CURRENT_TIMESTAMP,//发布的时间
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, //修改的时间
    is_closed TINYINT(1) DEFAULT 0, //是否关闭
    parent_id INT DEFAULT NULL, //这条留言指向的留言，就是他爹
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE, 
    FOREIGN KEY (parent_id) REFERENCES messages(id) ON DELETE CASCADE	//外键，不必多说，谁被删，这个也被删
);
```
## 结语

第一次写这么多代码，关于状态码，有很多都不是用的原本的含义，但是也懒得改了，下次再我一定会好好查清楚的！
	
总的来说，这个小项目实现了一个简单的留言板功能，数据库采用MySQL，写完代码的时候脑子还是嗡嗡的，关于这个文档，写着也是很难受，因为设置了好多状态码，好难找)关于这点我会好好反思如何改进的。