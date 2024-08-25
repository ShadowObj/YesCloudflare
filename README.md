# YesCloudflare

查询Cloudflare反代节点小工具

程序原型由 [Joey Huang](https://t.me/Joeyblog/) 开发

Telegram反馈群: https://t.me/+ft-zI76oovgwNmRh/

### 用法 Usage

```
Usage of ShadowObj/yescloudflare:

-c/-config string
          指定配置文件 (默认config.toml)
          注意: 指定配置文件不会覆盖命令行参数。
-o/-output ip.txt
          指定输出文件 (默认ip.txt)
-A/-auto
          自动获取下一页内容 (默认需要确认)
-key apikey
          指定APIKEY
-norepeat
          自动去除重复IP (默认不去除)
-port port
          指定端口 (默认全部, 可用英文逗号分隔)
-asn asn1,asn2
          指定ASN (默认全部, 可用英文逗号分隔)
-region CN,HK,JP,KR,TW
          指定地区ISO3166二字码
          (默认全部, 可用英文逗号分隔)
-page 1-10
	  指定需要查询的页面
	  (默认1-10, 范围应为第1到100页)
```

附: 
1. [config.toml 配置文件完整示例](https://github.com/ShadowObj/YesCloudflare/blob/main/config.toml)
2. [小白向 - 如何获取Cloudflare反代IP列表指北](https://telegra.ph/%E5%A6%82%E4%BD%95%E8%8E%B7%E5%8F%96Cloudflare%E5%8F%8D%E4%BB%A3IP%E5%88%97%E8%A1%A8%E6%8C%87%E5%8C%97-08-21)

### 如何获取APIKEY

进入censys的Search API汇总页
https://search.censys.io/api#/hosts/searchHosts

如果未注册：

点击右上角Register，按流程进行(需要常见邮箱，不支持临时邮箱)，进行邮箱验证，**登入时选择"Censys Search"**，重新点击上方网址

如果已注册：进入上方网址

在上述流程完成后，来到下方任意一个API块，点击Try it out，再点击Execute，在Responses中会出现用于curl的一串命令，其中"Authorization: Basic"后的一长串字符即为APIKEY。

*注意：不要复制额外的空格哟！(＾Ｕ＾)ノ~*
