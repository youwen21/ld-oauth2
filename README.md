# LD-Oauth2
本项目目的： 帮佬友快速对接 Linux.do Oauth2, 简化对接linux.do登录流程

本系统帮业务系统简化的工作：
 - token 过期维护
 - token 置换 用户信息
 - OAuth2 验证

## 角色说明
- 用户： 访问业务系统的用户，需要获取linux.do用户信息
- 业务系统： 需要对接linux.do的 系统。
- ld-auth: 本应用。
- linux.do：linux.do站（简称：L站）。

## 对接文档

## 执行流程
业务系统 -> ld-auth -> ld-auth -> L站 -> ld-auth -> 业务系统

- 用户在业务系统 请求 linux.do 用户信息授权， 业务系统准备参数， 跳转到 ld-auth
- ld-auth 自跳转一次， 设置cookie 信息 （没有数据库，用 cookie保存信息， 所以需要本次跳转）
- ld-auth 跳转到 linux.do, 用户 授权用户信息
- linux.do 回调跳转到 ld-auth
- ld-auth 回调跳转到 业务系统
- 业务系统 接收用户信息

## 业务系统工作
 - 组装请求ld-auth 的参数
 - 准备一个回调地址，接收linux.do 用户信息

`请求参数`
```text
	redirect_to string `json:"redirect_to" form:"redirect_to"` // 业务系统回调地址
	check-param-x string `json:"check-param-x" form:"check-param-x"` // 非必须
```


`业务系统回调地址 - 回调参数` 
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


## 运行原理

### 角色
- Linux.do站: www.linux.do
- A站 ： 您要接入linux.do账号的站
- LD-Oauth2服务 ： 本仓库代码服务

LD-Oauth2 是A站和Linux.do站的中间服务，免去A站获取code,  保管token, 刷新token等工作， 直接获取linux.do的用户信息。

A站请求LD-Oauth2服务，LD-Oauth2服务从Linux.do站获取用户信息，成功后 跳转回A站，并且携带用户信息。


