package global

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
	"miaoshaSystem/sql"
	"time"
)

/*
	func NewWriter() *kafka.Writer {
		w := &kafka.Writer{
			Addr:     kafka.TCP("localhost:9092"), // Kafka broker 地址
			Topic:    "my_topic",                  // Kafka 主题
			Balancer: &kafka.LeastBytes{},         // 平衡策略
		}
		return w
	}
*/
func SendToKafka(productName, username string) error {
	// Kafka 配置
	writer := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"), // Kafka 服务器地址
		Topic:        "seckill_requests",          // Kafka 主题
		BatchSize:    100,                         //批次大小
		BatchTimeout: 100 * time.Millisecond,
	}

	// 创建秒杀请求消息
	message := map[string]string{
		"productName": productName,
		"username":    username,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	// 发送到 Kafka
	err = writer.WriteMessages(context.Background(), kafka.Message{
		Value: messageBytes,
	})
	if err != nil {
		return fmt.Errorf("failed to send message to Kafka: %v", err)
	}
	return nil
}
func StartKafkaConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "seckill_requests",
		GroupID:  "seckill_group",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Failed to read message: %v\n", err)
			continue
		}

		var request map[string]string
		if err := json.Unmarshal(m.Value, &request); err != nil {
			fmt.Printf("Failed to unmarshal message: %v\n", err)
			continue
		}

		productName := request["productName"]
		username := request["username"]

		HandleSeckill(productName, username, c)

	}
}

var c *gin.Context

func HandleSeckill(productName, username string, c *gin.Context) {
	err := sql.DB.Model(&Product{}).Where("name = ?", productName).Update("num", gorm.Expr("num - ?", 1))
	if err != nil {
		fmt.Printf("创建订单失败: %v\n", err)
	} //先减库存
	errr := sql.R.Set(context.Background(), username, productName, 60*60*time.Second).Err()
	if errr != nil {
		fmt.Printf("Failed to set order: %v\n", err)
		return
	} //创建订单，保存在数据库中
	//接下来返回订单给前端
	c.JSON(400, gin.H{"success": "订单创建成功",
		"time":         time.Now(),
		"username":     username,
		"product name": productName,
		"注意":"未支付的订单将在一个小时之后失效",
	})
}
