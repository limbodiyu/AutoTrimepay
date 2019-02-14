# AutoTrimepay
Trimepay自动提现

## 用法

安装golang
```
CentOS:
yum install go -y

Debian/Ubuntu:
apt-get install golang -y
```

下载源代码
```
curl -O https://raw.githubusercontent.com/CGDF-Github/AutoTrimepay/master/AutoTrimepay.go
```

自行修改源代码中的以下内容
```
var email = "!!!change to your email!!!"
var password = "!!!change to your password!!!"
var method = "1" //1 : alipay  2 : wechat
```

使用`crontab -e`添加定时提现任务
每天可提现两次，所以设置两个时间点，请勿使用整点，否则大概率会导致提现撞车从而Tony钱包瞬间变空然后提现失败

例：每天9点52提现：
```
52 9 * * * go run 刚刚下载的文件路径/AutoTrimepay.go
```
