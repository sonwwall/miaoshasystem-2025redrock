// 无关紧要的用来模拟普通用户的代码，因为不怎么明白JS的语法，便请kimi代写了
// 引入axios库，用于发送HTTP请求
const axios = require('axios');

// 配置目标URL
const BASE_URL = 'http://localhost:8080/miaosha/'; // 后端服务的URL
const PRODUCT_NAME = 'vivo50'; // 抢购的产品名称
const JWT_TOKEN = 'YOUR_JWT_TOKEN_HERE'; // 替换为你的JWT Token

// 模拟用户小明的抢购请求
function simulateUserMiaosha() {
    axios.put(`${BASE_URL}${PRODUCT_NAME}`, { // 发送PUT请求
        // 这里可以添加你的请求体内容
        userId: 'xiaoming', // 模拟用户ID
        timestamp: Date.now() // 当前时间戳
    }, {
        headers: {
            'Authorization': `Bearer ${JWT_TOKEN}` // 设置JWT Token
        }
    })
        .then(response => { // 请求成功的回调
            console.log('Miaosha request succeeded:', response.data); // 打印成功响应
        })
        .catch(error => { // 请求失败的回调
            console.error('Miaosha request failed:', error); // 打印失败信息
        });
}

// 模拟用户小明的抢购行为
simulateUserMiaosha();