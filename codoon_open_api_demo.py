# -*- coding: UTF-8 -*-

import os, urlparse, subprocess, urllib2
from urllib import urlencode
    

if __name__ == "__main__":
    # 第一步，通过网页注册或者邮件，注册名字，回调地址，并获得client_id


    # 第二步，使用client_id，回调地址作为参数，调用codoon服务器认证授权页面，以获取AccessCode（这里的的回调地址和注册的回调地址要一致）：
    callback_url = "http://api.yourdomain.com/get_code"
    params = {
        'client_id': 'xxxxxxxxxxxxx',
        'redirect_uri': callback_url,
        'response_type': 'code'
    }

    print "https://openapi.codoon.com/authorize?scope=user,sports&" + urlencode(params)


    # 第三步，当用户登录并确认授权后，咕咚回调注册的回调地址，传回AccessCode，如： code=7e1d0072232572a5642514b064a551c8


    # 第四步，通过AccessCode换取Access Token
    import requests, base64

    client_key = "xxxxx"
    client_secret = "xxxxx"
    params = dict(

        client_id="xxxxx",
        grant_type="authorization_code",
        code="xxxxxxxxxxxxxxxxxxxxxx",
        redirect_uri="http://api.yourdomain.com/get_code",
        scope="user,sports"
    )

    headers = {'Authorization': 'Basic %s' % base64.b64encode(client_key + ":" + client_secret)}
    response = requests.post("https://openapi.codoon.com/token",data=params, headers=headers)
    print response.text

    {"access_token": {access_token}, "token_type": "bearer", "expire_in": 3600, "refresh_token": {refresh_token}, "scope": {scope}}
    # urllib2.urlopen("https://openapi.codoon.com/authorize?scope=user,sports&" + urlencode(params)) 

    
