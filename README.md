# YesCloudflare

查询Cloudflare反代节点小工具

程序原型由 [Joey Huang](https://t.me/Joeyblog/) 开发

Telegram反馈群: https://t.me/+ft-zI76oovgwNmRh/

点一下Star是最大的支持！谢谢！

### 用法 Usage

```
Usage of ShadowObj/yescloudflare:
  
  -A    自动获取下一页内容 (默认需要确认)
  --asn int
        指定ASN
  --auto
        自动获取下一页内容 (默认需要确认)
  --file string
        指定输出文件 (默认ip.txt) (default "ip.txt")
  --key string
        指定API密钥
  --norepeat
        自动去除重复IP (默认不去除)
  --port int
        指定端口 (默认全部)
  --region string
        指定地区
  --help 
        获取帮助
```

### 如何获取APIKEY

进入censys的Search API汇总页
https://search.censys.io/api#/hosts/searchHosts

如果未注册：

点击右上角Register，按流程进行(需要常见邮箱，不支持临时邮箱)，进行邮箱验证，**登入时选择"Censys Search"**，重新点击上方网址

如果已注册：进入上方网址

在上述流程完成后，来到下方任意一个API块，点击Try it out，再点击Execute，在Responses中会出现用于curl的一串命令，其中"Authorization: Basic"后的一长串字符即为APIKEY。

*注意：不要复制额外的空格哟！(＾Ｕ＾)ノ~*
