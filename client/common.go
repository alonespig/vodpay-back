package client

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"vodpay/form"
	"vodpay/utils"

	"github.com/google/uuid"
)

func CheckTimestamp(timestamp int64) bool {
	time := math.Abs(float64(time.Now().UnixMilli() - timestamp))
	if time > 600000 {
		return false
	}
	return true
}

func NewUuid() string {
	return utils.MD5(uuid.New().String())
}

func CheckOrderSign(orderForm form.OrderForm, secretKey string) string {
	var builder strings.Builder
	builder.WriteString("appid=")
	builder.WriteString(orderForm.Appid)
	builder.WriteString("&product_code=")
	builder.WriteString(strconv.Itoa(int(orderForm.ProductCode)))
	builder.WriteString("&account=")
	builder.WriteString(orderForm.Mobile)
	builder.WriteString("&order_no=")
	builder.WriteString(orderForm.ChannelOrderNo)
	builder.WriteString("&timestamp=")
	builder.WriteString(strconv.FormatInt(orderForm.Timestamp, 10))
	builder.WriteString("&secret_key=")
	builder.WriteString(secretKey)
	log.Println(builder.String())
	return utils.MD5(builder.String())
}

func CheckOrderQuerySign(orderForm form.OrderQueryForm, secretKey string) string {
	var builder strings.Builder
	builder.WriteString("appid=")
	builder.WriteString(orderForm.Appid)
	builder.WriteString("&product_code=")
	builder.WriteString(strconv.Itoa(int(orderForm.ProductCode)))
	builder.WriteString("&order_no=")
	builder.WriteString(orderForm.ChannelOrderNo)
	builder.WriteString("&timestamp=")
	builder.WriteString(strconv.FormatInt(orderForm.Timestamp, 10))
	builder.WriteString("&secret_key=")
	builder.WriteString(secretKey)
	log.Println(builder.String())
	return utils.MD5(builder.String())
}
