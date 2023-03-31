<<<<<<< HEAD
<<<<<<< HEAD
### 友情提示：近阶段微信风控严重，封号几率超级大，请大家勿用

# 基于chatGPT wechatbot **(已适配GPT3.5)** 
> 将个人微信化身GPT机器人，
> 项目基于[openwechat](https://github.com/eatmoreapple/openwechat) 开发。
> 大家记得登录完微信之后会生成一个 stro..json 文件，每次登录记得删除这个文件哈，不然会加载历史会话

### chatGPT网页版（已续费 欢迎使用）
http://gpt.wxredcover.cn/

>欢迎大家start!!!!!!

## 公众号版本
https://gitee.com/lmuiotctf/chat_wxmp

### 功能
 * GPT机器人模型热度可配置
 * 提问增加上下文
 * 机器人群聊@回复
 * 机器人私聊回复
 * 好友添加自动通过

# 注意事项
> * 项目仅供娱乐，滥用可能有微信封禁的风险，请勿用于商业用途。
> * 请注意收发敏感信息，本项目不做信息过滤。


# 项目部署
```sh
# 复制配置文件，根据自己实际情况，调整配置里的内容
修改 config.json

其中配置文件参考下边的配置文件说明。

# 快速开始


# 启动项目
go run main.go

````

# 配置文件说明
````
{
  "api_key": "your api key",
  "auto_pass": true,
  "session_timeout": 60,
  "max_tokens": 1024,
  "model": "text-davinci-003",
  "temperature": 1,
  "reply_prefix": "来自机器人回复：",
  "session_clear_token": "清空会话"
}

api_key：openai api_key
auto_pass:是否自动通过好友添加
session_timeout：会话超时时间，默认60秒，单位秒，在会话时间内所有发送给机器人的信息会作为上下文。
max_tokens: GPT响应字符数，最大2048，默认值512。max_tokens会影响接口响应速度，字符越大响应越慢。
model: GPT选用模型，默认text-davinci-003，具体选项参考官网训练场
temperature: GPT热度，0到1，默认0.9。数字越大创造力越强，但更偏离训练事实，越低越接近训练事实
reply_prefix: 私聊回复前缀
session_clear_token: 会话清空口令，默认`下一个问题`
````
# chatGpt key获取教程
首先需要在chatgpt官网:https://openai.com/

注册一个账号，这里我就不多说了，注册完成之后登录即可。
https://beta.openai.com/overview

然后在右上角的 View Api KeY 进行创建查看


# 使用示例
### 向机器人发送`清空会话`，清空会话信息。
### 私聊
<img width="300px" src="https://gitee.com/lmuiotctf/chatGpt_wechat/raw/master/image/no1.png"/>

### 群聊@回复
<img width="300px" src="https://gitee.com/lmuiotctf/chatGpt_wechat/raw/master/image/no2.png"/>

### 添加微信（备注: wechabot）进群交流

**如果二维码图片没显示出来，请添加微信号 留言**

<img width="210px"  src="https://gitee.com/lmuiotctf/chatGpt_wechat/raw/master/image/wechat.png" align="left">

=======
=======
>>>>>>> 425bd6828877888580bc966f789bf5573ba6843e
# ChatWechat

#### 介绍
微信机器人开发

#### 软件架构
软件架构说明


#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
<<<<<<< HEAD
>>>>>>> 425bd6828877888580bc966f789bf5573ba6843e
=======
>>>>>>> 425bd6828877888580bc966f789bf5573ba6843e
