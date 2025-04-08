// 无关紧要的用来模拟恶意用户的代码，因为不怎么明白JS的语法，便请kimi代写了
// 引入axios库，用于发送HTTP请求
const axios = require('axios');
// 引入uuid库，用于生成唯一的请求ID
const { v4: uuidv4 } = require('uuid');

// 配置目标URL
const BASE_URL = 'http://localhost:8080/miaosha/'; // 后端服务的URL
const PRODUCT_NAME = 'vivo50'; // 替换为你的产品名称
const REQUESTS_PER_SECOND = 100; // 每秒发送的请求数量
const TOTAL_REQUESTS = 10000; // 总共发送的请求数量

let requestCount = 0; // 已发送的请求数量计数器

// 定义发送请求的函数
function sendRequest() {
	const requestId = uuidv4(); // 生成唯一的请求ID
	axios.put(`${BASE_URL}${PRODUCT_NAME}`, { // 发送PUT请求
		// 这里可以添加你的请求体内容
		requestId: requestId, // 请求ID
		timestamp: Date.now() // 当前时间戳
	})
		.then(response => { // 请求成功的回调
			console.log(`Request ${requestId} succeeded:`, response.data); // 打印成功响应
		})
		.catch(error => { // 请求失败的回调
			console.error(`Request ${requestId} failed:`, error); // 打印失败信息
		})
		.finally(() => { // 请求完成后的回调
			requestCount++; // 增加已发送请求数量
			if (requestCount < TOTAL_REQUESTS) { // 如果未达到总请求数量
				sendRequest(); // 继续发送请求
			}
		});
}

// 控制请求频率
setInterval(() => { // 每秒执行一次
	for (let i = 0; i < REQUESTS_PER_SECOND; i++) { // 每秒发送指定数量的请求
		if (requestCount < TOTAL_REQUESTS) { // 如果未达到总请求数量
			sendRequest(); // 发送请求
		}
	}
}, 100); // 间隔1000毫秒（1秒）

console.log(`Starting to send ${TOTAL_REQUESTS} requests at ${REQUESTS_PER_SECOND} requests per second.`);