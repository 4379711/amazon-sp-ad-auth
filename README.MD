### 亚马逊sp-api和亚马逊广告api的授权换取密钥
1. ad目录是Advertising的授权代码
2. sp目录是SellingPartner的授权代码
3. 修改ad/default_engine.go,sp/default_engine.go里的相关信息,换成你的app信息
4. 以下操作均在test里
   - 运行TestAuthUrl函数,获取授权地址,在店铺里访问后,获取授权码
   - 修改TestRefreshToken参数,参数为授权码(过期时间很短),运行后换取refresh token
   - 运行TestAccessToken函数,获取access token