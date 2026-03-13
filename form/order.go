package form

type baseForm struct {
	Appid     string `form:"appid" binding:"required"`
	Timestamp int64  `form:"timestamp" binding:"required"`
	Sign      string `form:"sign" binding:"required"`
}

type OrderForm struct {
	baseForm
	Callback       string `form:"callback"`
	ProductCode    uint   `form:"product_code"`
	Mobile         string `form:"account" binding:"required"`
	ChannelOrderNo string `form:"order_no" binding:"required"`
	Ip             string
}

type OrderQueryForm struct {
	baseForm
	ProductCode    uint   `form:"product_code"`
	ChannelOrderNo string `form:"order_no" binding:"required"`
}

type PageForm struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"pageSize" binding:"required"`
}

type OrderListQueryForm struct {
	AccountID      string `form:"accountID"`
	ChannelOrderNo string `form:"channelOrderNo"`
	ChannelID      int    `form:"channelID"`
	ProjectID      int    `form:"projectID"`
	ProductID      int    `form:"productID"`
	BrandSkuSpecID int    `form:"brandSkuSpecID"`
	BrandID        int    `form:"brandID"`
	SKUID          int    `form:"skuID"`
	SpecID         int    `form:"specID"`
	SupplierID     int    `form:"supplierID"`
	StartTime      int    `form:"startTime"`
	EndTime        int    `form:"endTime"`
	Status         *int   `form:"status"`
	Page           int    `form:"page" binding:"required"`
	PageSize       int    `form:"pageSize" binding:"required"`
}

type Order struct {
	ID              int64  `json:"id"`
	ProductID       int64  `json:"productID"`
	ProductName     string `json:"productName"`
	SupplierID      int64  `json:"supplierID"`
	SupplierName    string `json:"supplierName"`
	SupplierPrice   int    `json:"supplierPrice"`
	ProjectName     string `json:"projectName"`
	ChannelName     string `json:"channelName"`
	ChannelOrderNo  string `json:"channelOrderNo"`
	Price           int    `json:"price"`
	PlatFromOrderNo string `json:"platFromOrderNo"`
	SelfOrderNo     string `json:"selfOrderNo"`
	AccountID       string `json:"accountID"`
	Msg             string `json:"msg"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
	Status          int    `json:"status"`
}

type OrderListResp struct {
	Total  int64    `json:"total"`
	Orders []*Order `json:"orders"`
}

type OrderQueryResponse struct {
	Code            int    `json:"code"`
	Msg             string `json:"msg"`
	OrderStatus     int    `json:"order_status"`
	PlatformOrderNo string `json:"platform_order_no"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ChannelLineChartQueryForm struct {
	ChannelID int   `form:"channelID"`
	ProjectID int   `form:"projectID"`
	ProductID int   `form:"productID"`
	Timestamp int64 `form:"timestamp" binding:"required"`
}

type ChannelLineChartResp struct {
	Points []*Point `json:"points"`
}

type SupplierOrder struct {
	Name           string `json:"name"`
	Price          int    `json:"price"`
	SupplierName   string `json:"supplierName"`
	SupProductCode string `json:"supplierProductCode"`
	Total          int64  `json:"total"`
	Status         int    `json:"status"`
}

type SupplierOrderListResp struct {
	Total  int64           `json:"total"`
	Orders []SupplierOrder `json:"orders"`
}

type SupplierOrderListQueryForm struct {
	SupplierID     int  `form:"supplierID"`
	BrandSkuSpecID int  `form:"brandSkuSpecID"`
	StartTime      int  `form:"startTime" binding:"required"`
	EndTime        int  `form:"endTime" binding:"required"`
	Status         *int `form:"status"`
}
