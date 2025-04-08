# 秒杀系统

### 一、高并发处理
我通过中间件 **ban 掉高频发送请求的脚本哥**，通过**Kafka 消息队列** 的缓冲机制，处理高并发请求


### 二、库存问题
优先进行 **减少库存操作**，这样就能防止超卖。即使后续服务未能成功，也只会少卖不会多卖，从而避免因超卖引发的业务问题。

### 三、订单存储
之前寒假考核的时候，我还不太会用非关系型数据库，存订单相当恼火。这次通过使用 **Redis** 存储键值对，解决了这个问题，爽。
##接口文档
以下是基于你提供的代码的接口文档，使用Markdown格式编写：

# 秒杀系统接口文档

## 1. 注册接口

### 接口地址
`POST /register`

### 请求参数
| 参数名 | 类型 | 是否必填 | 示例值 | 描述 |
| --- | --- | --- | --- | --- |
| name | string | 是 | "张三" | 用户名 |
| age | int | 是 | 20 | 年龄 |
| address | string | 是 | "北京市海淀区" | 地址 |
| avatar | string | 是 | "http://example.com/avatar.jpg" | 头像链接 |
| password | string | 是 | "123456" | 密码 |

### 请求体示例
```json
{
    "name": "张三",
    "age": 20,
    "address": "北京市海淀区",
    "avatar": "http://example.com/avatar.jpg",
    "password": "123456"
}
```

### 返回结果
| 状态码 | 返回值 | 描述 |
| --- | --- | --- |
| 200 | {"message": "user registered successfully"} | 注册成功 |
| 400 | {"error": "error message"} | 请求参数错误 |
| 500 | {"error": "user creation failed"} | 服务器内部错误 |

## 2. 登录接口

### 接口地址
`POST /login`

### 请求参数
| 参数名 | 类型 | 是否必填 | 示例值 | 描述 |
| --- | --- | --- | --- | --- |
| name | string | 是 | "张三" | 用户名 |
| pass | string | 是 | "123456" | 密码 |

### 请求体示例
```json
{
    "name": "张三",
    "pass": "123456"
}
```

### 返回结果
| 状态码 | 返回值 | 描述 |
| --- | --- | --- |
| 200 | {"token": "token_string"} | 登录成功，返回JWT令牌 |
| 400 | {"error": "error message"} | 请求参数错误 |
| 401 | {"error": "Invalid token"} | 令牌无效 |

## 3. 创建秒杀产品接口

### 接口地址
`POST /createmiaosha`

### 请求参数
| 参数名 | 类型 | 是否必填 | 示例值 | 描述 |
| --- | --- | --- | --- | --- |
| name | string | 是 | "iPhone 13" | 产品名称 |
| num | int | 是 | 100 | 产品数量 |
| producter | string | 是 | "Apple" | 生产商 |
| time_begintokill | int64 | 是 | 1640995200 | 秒杀开始时间（时间戳） |
| time_endkill | int64 | 是 | 1641081600 | 秒杀结束时间（时间戳） |

### 请求体示例
```json
{
    "name": "iPhone 13",
    "num": 100,
    "producter": "Apple",
    "time_begintokill": 1640995200,
    "time_endkill": 1641081600
}
```

### 返回结果
| 状态码 | 返回值 | 描述 |
| --- | --- | --- |
| 200 | {"message": "Product creation failed"} | 创建成功 |
| 400 | {"error": "error message"} | 请求参数错误 |
| 500 | {"error": "Product creation failed"} | 服务器内部错误 |

## 4. 秒杀接口

### 接口地址
`POST /miaosha/:productName`

### 请求参数
| 参数名 | 类型 | 是否必填 | 示例值 | 描述 |
| --- | --- | --- | --- | --- |
| productName | string | 是 | "iPhone 13" | 产品名称 |
| token | string | 是 | "token_string" | JWT令牌 |

### 请求体示例
```json
{
    "token": "token_string"
}
```

### 返回结果
| 状态码 | 返回值 | 描述 |
| --- | --- | --- |
| 200 | {"success": "订单创建成功", "time": "current_time", "username": "username", "product name": "productName", "注意": "未支付的订单将在一个小时之后失效"} | 秒杀成功 |
| 400 | {"error": "error message"} | 请求参数错误 |
| 401 | {"error": "Invalid token"} | 令牌无效 |
| 500 | {"error": "Failed to send request to Kafka"} | 服务器内部错误 |

## 5. Kafka 消费者接口（内部接口，不对外暴露）

### 接口地址
无

### 功能描述
该接口用于处理 Kafka 队列中的秒杀请求，将请求发送到后端进行处理。

### 请求参数
无

### 返回结果
无
