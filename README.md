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
var email = ""       //Trimepay账户
var password = ""    //密码
var method = "1"     //1:支付宝  2:微信
var supportTip = 0.3 //赞助小费，单位元，可为0
```

使用`crontab -e`添加定时提现任务

每天可提现两次，所以设置两个时间点，请勿使用整点，否则大概率会导致提现撞车从而Tony钱包瞬间变空然后提现失败

例：每日几时几分提现
```
分 时 * * * go run 刚刚下载的文件路径/AutoTrimepay.go
```

加入crontab前建议先手动执行一遍，错误日志会保存在`AutoTrimepay.log`中
