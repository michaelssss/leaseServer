# 本程序用于家里动态IP暴露值公网固定服务器

随便写写,用于将家庭服务器对外公布,拒绝花生壳的应用  

## 使用方法:  
    go build main.go
    nohup ./main ServerAccessKey &
其中SeverAccessKey是Server与Client通信的密钥