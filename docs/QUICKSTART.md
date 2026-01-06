# 快速开始指南

本指南将帮助你快速上手 WEEX Contract API Golang SDK。

## 目录

- [安装](#安装)
- [获取 API 凭证](#获取-api-凭证)
- [初始化客户端](#初始化客户端)
- [市场数据示例](#市场数据示例)
- [账户管理示例](#账户管理示例)
- [交易操作示例](#交易操作示例)
- [错误处理](#错误处理)
- [最佳实践](#最佳实践)

## 安装

```bash
go get github.com/weex/openapi-contract-go-sdk
```

**要求**: Go 1.20 或更高版本

## 获取 API 凭证

1. 登录 WEEX 账户
2. 进入 API 管理页面
3. 创建新的 API Key
4. 保存以下信息：
   - API Key
   - Secret Key
   - Passphrase

⚠️ **安全提示**: 永远不要将 API 凭证硬编码在代码中或提交到版本控制系统。建议使用环境变量。

## 初始化客户端

### 公开端点客户端（无需凭证）

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/weex/openapi-contract-go-sdk/weex"
)

func main() {
    // 创建公开客户端（用于市场数据）
    config := weex.NewDefaultConfig()
    client, err := weex.NewPublicClient(config)
    if err != nil {
        log.Fatal(err)
    }

    // 使用客户端...
}
```

### 完整功能客户端（需要凭证）

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/weex/openapi-contract-go-sdk/weex"
)

func main() {
    // 从环境变量读取凭证
    config := weex.NewDefaultConfig().
        WithAPIKey(os.Getenv("WEEX_API_KEY")).
        WithSecretKey(os.Getenv("WEEX_SECRET_KEY")).
        WithPassphrase(os.Getenv("WEEX_PASSPHRASE")).
        WithLogLevel(weex.LogLevelInfo)

    // 创建客户端
    client, err := weex.NewClient(config)
    if err != nil {
        log.Fatal(err)
    }

    // 使用客户端...
}
```

## 市场数据示例

### 获取实时行情

```go
ctx := context.Background()

// 获取单个合约行情
ticker, err := client.Market().GetTicker(ctx, "cmt_btcusdt")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("BTC/USDT 价格: %s\n", ticker.LastPrice)
fmt.Printf("24h涨跌幅: %s%%\n", ticker.PriceChangePercent)
fmt.Printf("24h成交量: %s\n", ticker.Volume)
```

### 获取订单簿深度

```go
depth, err := client.Market().GetDepth(ctx, &market.GetDepthRequest{
    Symbol: "cmt_btcusdt",
    Limit:  20, // 获取20档深度
})
if err != nil {
    log.Fatal(err)
}

fmt.Println("买单（前5档）:")
for i := 0; i < 5 && i < len(depth.Bids); i++ {
    fmt.Printf("  价格: %s, 数量: %s\n", depth.Bids[i].Price, depth.Bids[i].Quantity)
}

fmt.Println("卖单（前5档）:")
for i := 0; i < 5 && i < len(depth.Asks); i++ {
    fmt.Printf("  价格: %s, 数量: %s\n", depth.Asks[i].Price, depth.Asks[i].Quantity)
}
```

### 获取K线数据

```go
klines, err := client.Market().GetKlines(ctx, &market.GetKlinesRequest{
    Symbol:   "cmt_btcusdt",
    Interval: types.Interval1Hour,
    Limit:    10,
})
if err != nil {
    log.Fatal(err)
}

for _, kline := range klines {
    fmt.Printf("时间: %v, 开: %s, 高: %s, 低: %s, 收: %s, 量: %s\n",
        time.UnixMilli(kline.OpenTime).Format("2006-01-02 15:04"),
        kline.Open, kline.High, kline.Low, kline.Close, kline.Volume)
}
```

## 账户管理示例

### 查询账户余额

```go
assets, err := client.Account().GetAccountBalance(ctx)
if err != nil {
    log.Fatal(err)
}

for _, asset := range assets {
    fmt.Printf("%s:\n", asset.CoinName)
    fmt.Printf("  可用: %s\n", asset.Available)
    fmt.Printf("  冻结: %s\n", asset.Frozen)
    fmt.Printf("  权益: %s\n", asset.Equity)
    fmt.Printf("  未实现盈亏: %s\n", asset.UnrealizedPnl)
}
```

### 查询持仓

```go
positions, err := client.Account().GetAllPositions(ctx, &account.GetAllPositionsRequest{
    Symbol: "cmt_btcusdt", // 可选，不填则查询所有
})
if err != nil {
    log.Fatal(err)
}

for _, pos := range positions {
    fmt.Printf("%s [%s]:\n", pos.Symbol, pos.PositionSide)
    fmt.Printf("  持仓数量: %s\n", pos.Size)
    fmt.Printf("  开仓均价: %s\n", pos.AverageOpenPrice)
    fmt.Printf("  未实现盈亏: %s\n", pos.UnrealizedPnl)
    fmt.Printf("  杠杆倍数: %sx\n", pos.Leverage)
    fmt.Printf("  强平价格: %s\n", pos.LiquidatePrice)
}
```

### 调整杠杆

```go
leverageResp, err := client.Account().AdjustLeverage(ctx, &account.AdjustLeverageRequest{
    Symbol:     "cmt_btcusdt",
    MarginMode: int(types.MarginModeShared), // 全仓模式
    Leverage:   types.NewDecimalFromString("20"), // 20倍杠杆
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("杠杆已调整为: %sx\n", leverageResp.Leverage)
```

## 交易操作示例

### 下限价单

```go
import "time"

// 生成唯一的客户端订单ID
clientOid := fmt.Sprintf("order_%d", time.Now().UnixMilli())

order, err := client.Trade().PlaceOrder(ctx, &trade.PlaceOrderRequest{
    Symbol:     "cmt_btcusdt",
    ClientOid:  clientOid,
    Size:       types.NewDecimalFromString("0.01"), // 下单数量
    Type:       types.OrderTypeOpenLong,            // 开多
    OrderType:  types.OrderExecNormal,              // 普通委托
    MatchPrice: types.PriceMatchLimit,              // 限价
    Price:      types.NewDecimalFromString("50000"), // 限价价格
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("订单已提交!\n")
fmt.Printf("订单ID: %s\n", order.OrderId)
fmt.Printf("客户端订单ID: %s\n", order.ClientOid)
```

### 下市价单

```go
order, err := client.Trade().PlaceOrder(ctx, &trade.PlaceOrderRequest{
    Symbol:     "cmt_btcusdt",
    ClientOid:  fmt.Sprintf("market_%d", time.Now().UnixMilli()),
    Size:       types.NewDecimalFromString("0.01"),
    Type:       types.OrderTypeOpenLong,
    OrderType:  types.OrderExecNormal,
    MatchPrice: types.PriceMatchMarket, // 市价
    // 市价单不需要设置 Price
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("市价单已提交: %s\n", order.OrderId)
```

### 批量下单

```go
batchResp, err := client.Trade().PlaceOrdersBatch(ctx, &trade.PlaceOrdersBatchRequest{
    Orders: []trade.PlaceOrderRequest{
        {
            Symbol:     "cmt_btcusdt",
            ClientOid:  "batch1",
            Size:       types.NewDecimalFromString("0.01"),
            Type:       types.OrderTypeOpenLong,
            OrderType:  types.OrderExecNormal,
            MatchPrice: types.PriceMatchLimit,
            Price:      types.NewDecimalFromString("49000"),
        },
        {
            Symbol:     "cmt_btcusdt",
            ClientOid:  "batch2",
            Size:       types.NewDecimalFromString("0.01"),
            Type:       types.OrderTypeOpenLong,
            OrderType:  types.OrderExecNormal,
            MatchPrice: types.PriceMatchLimit,
            Price:      types.NewDecimalFromString("48000"),
        },
    },
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("成功: %d, 失败: %d\n", len(batchResp.Success), len(batchResp.Failed))
```

### 撤销订单

```go
cancelResp, err := client.Trade().CancelOrder(ctx, &trade.CancelOrderRequest{
    OrderId: order.OrderId, // 或使用 ClientOid
    Symbol:  "cmt_btcusdt",
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("订单已撤销: %s\n", cancelResp.OrderId)
```

### 查询当前订单

```go
orders, err := client.Trade().GetCurrentOrderStatus(ctx, &trade.GetOrdersRequest{
    Symbol: "cmt_btcusdt",
    Limit:  20,
})
if err != nil {
    log.Fatal(err)
}

for _, order := range orders.Orders {
    fmt.Printf("订单: %s, 状态: %d, 价格: %s, 数量: %s, 已成交: %s\n",
        order.OrderId, order.State, order.Price, order.Size, order.FilledSize)
}
```

### 平仓

```go
closeResp, err := client.Trade().ClosePositions(ctx, &trade.ClosePositionsRequest{
    Symbol:       "cmt_btcusdt",
    PositionSide: "LONG",
    Size:         types.NewDecimalFromString("0"), // 0表示全部平仓
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("平仓订单ID: %s\n", closeResp.OrderId)
```

## 错误处理

SDK 提供了详细的错误分类，帮助你更好地处理各种错误情况。

```go
import "errors"

ticker, err := client.Market().GetTicker(ctx, "cmt_btcusdt")
if err != nil {
    // 检查是否为 API 错误
    var apiErr *weex.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API 错误码: %s\n", apiErr.Code)
        fmt.Printf("错误信息: %s\n", apiErr.Message)
        fmt.Printf("HTTP状态: %d\n", apiErr.HTTPStatus)

        // 判断是否可重试
        if apiErr.IsRetriable() {
            fmt.Println("这个错误可以重试")
            // 实现重试逻辑...
        }

        // 判断错误类型
        if apiErr.IsAuthError() {
            fmt.Println("认证错误，请检查API凭证")
        } else if apiErr.IsRateLimitError() {
            fmt.Println("速率限制，请稍后重试")
        }
    }

    // 检查是否为网络错误
    var netErr *weex.NetworkError
    if errors.As(err, &netErr) {
        fmt.Printf("网络错误: %v\n", netErr)
    }

    return
}
```

## 最佳实践

### 1. 使用环境变量存储凭证

```bash
export WEEX_API_KEY="your-api-key"
export WEEX_SECRET_KEY="your-secret-key"
export WEEX_PASSPHRASE="your-passphrase"
```

```go
config := weex.NewDefaultConfig().
    WithAPIKey(os.Getenv("WEEX_API_KEY")).
    WithSecretKey(os.Getenv("WEEX_SECRET_KEY")).
    WithPassphrase(os.Getenv("WEEX_PASSPHRASE"))
```

### 2. 使用 Context 控制超时

```go
// 设置5秒超时
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

ticker, err := client.Market().GetTicker(ctx, "cmt_btcusdt")
```

### 3. 使用客户端订单ID追踪订单

```go
// 使用有意义的客户端订单ID
clientOid := fmt.Sprintf("strategy_a_%d", time.Now().UnixMilli())

order, err := client.Trade().PlaceOrder(ctx, &trade.PlaceOrderRequest{
    ClientOid: clientOid,
    // ... 其他参数
})

// 稍后可以通过客户端订单ID查询
orderInfo, err := client.Trade().GetSingleOrderInfo(ctx, "", clientOid, "cmt_btcusdt")
```

### 4. 处理精度问题

```go
// 使用 Decimal 类型保证精度
price := types.NewDecimalFromString("50123.45")
size := types.NewDecimalFromString("0.001")

// 如果需要转换为 float64
priceFloat, err := price.Float64()
if err != nil {
    log.Fatal(err)
}
```

### 5. 启用日志调试

```go
config := weex.NewDefaultConfig().
    WithLogLevel(weex.LogLevelDebug) // 开发环境使用Debug

// 生产环境使用Info或Warn
config.WithLogLevel(weex.LogLevelInfo)
```

### 6. 自定义配置

```go
config := weex.NewDefaultConfig().
    WithHTTPTimeout(30 * time.Second).  // HTTP超时
    WithMaxRetries(5).                  // 最大重试次数
    WithLogLevel(weex.LogLevelInfo)

// 关闭速率限制（不推荐）
config.EnableRateLimit = false
```

### 7. 批量操作提高效率

```go
// 批量下单比单个下单更高效
batchResp, err := client.Trade().PlaceOrdersBatch(ctx, &trade.PlaceOrdersBatchRequest{
    Orders: orders, // 最多10个订单
})

// 批量撤单
cancelResp, err := client.Trade().CancelOrdersBatch(ctx, &trade.CancelOrdersBatchRequest{
    Orders: cancelRequests, // 最多10个订单
})
```

### 8. 错误重试示例

```go
func placeOrderWithRetry(client *weex.Client, req *trade.PlaceOrderRequest) (*trade.PlaceOrderResponse, error) {
    maxRetries := 3
    var lastErr error

    for i := 0; i < maxRetries; i++ {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        resp, err := client.Trade().PlaceOrder(ctx, req)
        if err == nil {
            return resp, nil
        }

        lastErr = err

        // 检查是否可重试
        var apiErr *weex.APIError
        if errors.As(err, &apiErr) && !apiErr.IsRetriable() {
            return nil, err // 不可重试，直接返回
        }

        // 等待后重试
        time.Sleep(time.Second * time.Duration(i+1))
    }

    return nil, lastErr
}
```

## 下一步

- 查看 [examples/](../examples/) 目录获取更多示例
- 阅读 [README.md](../README.md) 了解完整功能
- 查看 API 文档了解所有可用端点

## 支持

如有问题，请提交 Issue 或查看官方文档。
