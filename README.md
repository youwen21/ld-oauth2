# LD-Oauth2

快速接入 Linux.do Oauth2 账号。  

LD-Oauth2目标：简化接入Linux.do账号。

LD-Oauth2工作范围： 获取用户授权code, 保管token, 刷新token. 

## 运行原理

### 角色交代
  - Linux.do站: www.linux.do
  - A站 ： 您要接入linux.do账号的站, 
  - LD-Oauth2服务 ： 对接linux.do的服务，简化A站接入Linux.do账号。

LD-Oauth2 是A站和Linux.do站的中间服务，免去A站获取code,  保管token, 刷新token等工作， 直接获取linux.do的用户信息。


## A站对接文档

### 对接准备
 一个回调地址， url query接收 参数
```text
	Id         int    `json:"id" form:"id"`
	Username   string `json:"username" form:"username"`
	Name       string `json:"name" form:"name"`
	Active     bool   `json:"active" form:"active"`
	TrustLevel int    `json:"trust_level" form:"trust_level"`
	Silenced   bool   `json:"silenced" form:"silenced"`
```

### 对接

第一步，在A站,需要获取linux.do用户信息时，跳转到链接：

> {{host}}/oauth2/auth?redirect_to=/auth_finish.html?check-param-x=check_in_your_service"

#### host地址
- https://ldauth.himyou.com/
- 自己的服务器搭建: 于安全考虑，建议使用此方式。

#### 参数说明 
 - check-param-x ： 自定义key, 回调时会携带此key, 用于安全验证。

为增强安全，可自行开发修改，使用公钥私钥加密解密方式确保回调是安全的。


## 案例展示

https://player.himyou.com/demo.html


### 解析
player.himyou.com 是一个静态站，没有后台代码， 对接LD-Oauth2服务 也可以获取linux.do user信息。

## 自搭服务文档
下载源码编译或者下载执行包, 

运行命令：
```bash

export LD_CLIENT_ID=xxx
export LD_CLIENT_SECRET=xxx
ldauth

// 或者

ldauth -ld_client_id="" -ld_client_secret=""


```
优先读取args参数, 不存在时读取env

仓库地址：
 > https://github.com/youwen21/ld-oauth2







