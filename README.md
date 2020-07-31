### 使用的框架
```text
gin : web 开发框架
gorm ： 数据库开发框架
viper ： 配置文件读取
```

### 各个包说明

```text
config : 配置文件存放路径和加载
db ： 数据库配置
handler : web 请求具体执行类
model : 数据库orm 对象
nt : 国金公链的接口调用
response: 统一格式的数据返回
router : 路由管理器
sign : 参数验签
```
## Install

```sh
$ skaffold dev
```
## Usage

### local test
```sh
POST http://localhost:8089/api/v1/createFakeChainAddr
Content-Type: application/x-www-form-urlencoded

phone=17888888882&id_card=330619199106122021&eid=12123123&timestamp=1234567890123&nonce=12345678901&sign=5e47dced7af727563ba68f598c47ecd9ac94c799&device_num=12313&mobile_os=1231&mobile_os_version=1231&lng=0.0&lat=0.0


POST http://localhost:8080/api/v1/syncNt
Content-Type: application/x-www-form-urlencoded

phone=17888888881&timestamp=1234567890123&nonce=12345678901&sign=5e47dced7af727563ba68f598c47ecd9ac94c799&device_num=12313&mobile_os=1231&mobile_os_version=1231&lng=0.0&lat=0.0


GET http://tt.nt-chain.cn/service/api/v1/users/detail?attributes=invitation_code,node_type,created_at,is_identity_validated,avatar,nickname,phone_number,spaced_used,chain_addr,nt&phone_number=17888888881
X-Access-Key: DJ
X-Time: 1594908085
X-Signature: 0b6e4667c179fd6f531951a613ee02b3
Content-Type: application/json
X-App-Type: Android

POST http://tt.nt-chain.cn/api/v1/third/register
X-Access-Key: DJ
X-Time: 1594823089
X-Signature: 8d5a8d1b9d361fb6edbf2a3b05fa4206
Content-Type: application/json
X-App-Type: Android

{"mobile":"18867171234","invitation_code": "123"}
```
