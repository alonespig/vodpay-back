# VodPay 后端API

这是一个使用Go和Gin框架构建的简单后端API，用于测试前端接口。

## 功能

- 提供供应商数据查询接口

## API端点

### GET /api/supplier

返回供应商列表，包含以下字段：
- `id`: 供应商ID (int)
- `code`: 供应商编码 (string)
- `name`: 供应商名称 (string)
- `create_date`: 创建日期 (ISO 8601格式)

#### 响应格式

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "code": "SUP001",
      "name": "北京科技有限公司",
      "create_date": "2025-12-15T16:01:04.32664189+08:00"
    }
  ]
}
```

## 运行方式

1. 确保已安装Go 1.23+
2. 运行以下命令启动服务器：

```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动。

## 测试

可以使用以下命令测试API：

```bash
curl -X GET http://localhost:8080/api/supplier
```

