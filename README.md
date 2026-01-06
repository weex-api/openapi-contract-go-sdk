# WEEX Contract API - Golang SDK

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

官方 WEEX 合约交易 API 的 Golang SDK，提供完整的 REST API 和 WebSocket 支持。

[English](README.md) | [简体中文](README_zh.md)

## 功能特性

- ✅ **完整的 API 覆盖**
  - REST API: 市场数据、账户管理、交易操作
  - WebSocket: 公开和私有频道实时数据推送

- ✅ **生产级特性**
  - HMAC SHA256 认证
  - 自动重试机制（指数退避）
  - 基于权重的速率限制
  - 上下文支持（超时和取消）
  - 结构化日志

- ✅ **类型安全**
  - 强类型请求和响应模型
  - 枚举类型定义
  - 小数精度保护（使用字符串类型）

- ✅ **开发者友好**
  - 清晰的错误处理
  - 丰富的示例代码
  - 完整的文档
  - 单元测试覆盖

## 安装

```bash
go get github.com/weex/openapi-contract-go-sdk
```

**要求**: Go 1.20 或更高版本

## 快速开始

### 安装依赖

```bash
go get github.com/weex/openapi-contract-go-sdk
```

### 示例 1：获取市场数据（公开端点，无需API密钥）

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/weex/openapi-contract-go-sdk/weex"
)

func main() {
    // 创建公开客户端（无需API密钥）
    config := weex.NewDefaultConfig()
    client, err := weex.NewPublicClient(config)
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // 获取BTC永续合约的行情
    ticker, err := client.Market().GetTicker(ctx, "cmt_btcusdt")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("BTC/USDT:\n")
    fmt.Printf("  Last Price: %s\n", ticker.LastPrice)
    fmt.Printf("  24h Change: %s%%\n", ticker.PriceChangePercent)
    fmt.Printf("  24h Volume: %s\n", ticker.Volume)
}
```

### 示例 2：私有端点（需要API密钥）

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/weex/openapi-contract-go-sdk/weex"
)

func main() {
    // 创建配置
    config := weex.NewDefaultConfig()
    config.APIKey = "your-api-key"
    config.SecretKey = "your-secret-key"
    config.Passphrase = "your-passphrase"

    // 或使用链式调用
    config = weex.NewDefaultConfig().
        WithAPIKey("your-api-key").
        WithSecretKey("your-secret-key").
        WithPassphrase("your-passphrase").
        WithLogLevel(weex.LogLevelInfo)

    // 创建客户端
    client := weex.NewClient(config)

    // 使用客户端...
}
```

### 示例 3：WebSocket 实时行情（公开频道）

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/weex/openapi-contract-go-sdk/weex"
    "github.com/weex/openapi-contract-go-sdk/weex/websocket"
    "github.com/weex/openapi-contract-go-sdk/weex/websocket/public"
)

func main() {
    config := weex.NewDefaultConfig()
    client := public.NewClient(config)

    // 连接 WebSocket
    ctx := context.Background()
    if err := client.Connect(ctx); err != nil {
        log.Fatal(err)
    }

    // 订阅实时行情
    client.SubscribeTicker("cmt_btcusdt", func(ticker *websocket.TickerData) error {
        if len(ticker.Data) > 0 {
            t := ticker.Data[0]
            fmt.Printf("BTC/USDT: Price=%s, Change=%s%%\n",
                t.LastPrice, t.PriceChangePercent)
        }
        return nil
    })

    // 订阅深度数据
    client.SubscribeDepth("cmt_btcusdt", func(depth *websocket.DepthData) error {
        if len(depth.Data) > 0 {
            d := depth.Data[0]
            fmt.Printf("Best Bid: %s @ %s\n", d.Bids[0].Quantity, d.Bids[0].Price)
        }
        return nil
    })

    // 保持连接...
    select {}
}
```

### 示例 4：WebSocket 账户更新（私有频道）

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/weex/openapi-contract-go-sdk/weex"
    "github.com/weex/openapi-contract-go-sdk/weex/websocket"
    "github.com/weex/openapi-contract-go-sdk/weex/websocket/private"
)

func main() {
    config := weex.NewDefaultConfig().
        WithAPIKey(os.Getenv("WEEX_API_KEY")).
        WithSecretKey(os.Getenv("WEEX_SECRET_KEY")).
        WithPassphrase(os.Getenv("WEEX_PASSPHRASE"))

    auth := weex.NewAuthenticator(
        config.APIKey, config.SecretKey, config.Passphrase,
    )

    client := private.NewClient(config, auth)

    // 连接并认证
    ctx := context.Background()
    if err := client.Connect(ctx); err != nil {
        log.Fatal(err)
    }

    // 订阅账户余额变化
    client.SubscribeAccount(func(account *websocket.AccountData) error {
        for _, asset := range account.Data {
            fmt.Printf("%s: Available=%s, Frozen=%s\n",
                asset.CoinName, asset.Available, asset.Frozen)
        }
        return nil
    })

    // 订阅持仓变化
    client.SubscribePositions(func(position *websocket.PositionData) error {
        for _, pos := range position.Data {
            fmt.Printf("%s [%s]: Size=%s, PnL=%s\n",
                pos.Symbol, pos.PositionSide, pos.Size, pos.UnrealizedPnl)
        }
        return nil
    })

    // 订阅订单更新
    client.SubscribeOrders(func(order *websocket.OrderData) error {
        for _, o := range order.Data {
            fmt.Printf("Order %s: %s, Filled=%s/%s\n",
                o.OrderId, o.Symbol, o.FilledSize, o.Size)
        }
        return nil
    })

    // 保持连接...
    select {}
}
```


## 项目状态

✅ **生产就绪** - REST API 和 WebSocket 已全部完成！

### 已完成 ✅

**基础设施**:
- [x] 项目结构和配置
- [x] 核心类型定义（枚举、常量）
- [x] HMAC SHA256 认证实现
- [x] 错误处理框架（错误分类、错误码映射）
- [x] 日志接口（默认logger和no-op logger）
- [x] 配置系统（支持链式调用）
- [x] 重试机制（指数退避）
- [x] 速率限制器（基于令牌桶算法）

**REST API (40/40 端点 - 100%)**:
- [x] REST API 客户端核心
- [x] Market API (13个公开端点)
- [x] Account API (11个私有端点)
- [x] Trade API (16个私有端点)

**WebSocket API (8/8 频道 - 100%)**:
- [x] WebSocket 客户端核心（自动重连、心跳）
- [x] 订阅管理器
- [x] 公开频道：ticker, depth, candlestick, trades
- [x] 私有频道：account, positions, orders, fill

**示例和文档**:
- [x] REST API 示例（市场数据、账户、交易）
- [x] WebSocket 示例（公开频道、私有频道）
- [x] 快速入门指南
- [x] README 文档

### 待完善 📋

- [ ] 单元测试（目标 80% 覆盖率）
- [ ] 集成测试
- [ ] 更多 API 文档（认证、错误处理、WebSocket详细说明）
- [ ] CI/CD 流程

## 核心组件

### 认证 (Authentication)

SDK 使用 HMAC SHA256 签名算法对请求进行签名：

```go
// 签名算法
// Message = timestamp + method + requestPath + body
// Signature = Base64(HMAC-SHA256(secretKey, Message))

auth := weex.NewAuthenticator(apiKey, secretKey, passphrase)
headers := auth.GetRESTHeaders(0, "GET", "/capi/v2/market/contracts", "")
```

### 错误处理 (Error Handling)

SDK 提供详细的错误分类：

```go
err := client.SomeMethod(ctx)
if err != nil {
    var apiErr *weex.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API Error Code: %s\n", apiErr.Code)
        fmt.Printf("Message: %s\n", apiErr.Message)
        fmt.Printf("Retriable: %v\n", apiErr.IsRetriable())
        fmt.Printf("Type: %s\n", apiErr.Category.Type)
    }
}
```

**错误类型**:
- `ErrTypeAuth` - 认证错误（不可重试）
- `ErrTypeRateLimit` - 速率限制（可重试）
- `ErrTypeValidation` - 验证错误（不可重试）
- `ErrTypeNetwork` - 网络错误（可重试）
- `ErrTypeSystem` - 系统错误（可重试）
- `ErrTypePermission` - 权限错误（不可重试）
- `ErrTypeBusiness` - 业务错误（通常不可重试）

### 重试机制 (Retry Mechanism)

自动重试可重试的错误（速率限制、系统错误、网络错误）：

```go
config := weex.NewDefaultConfig()
config.MaxRetries = 3                        // 最多重试 3 次
config.InitialBackoff = 1 * time.Second      // 初始退避 1 秒
config.MaxBackoff = 30 * time.Second         // 最大退避 30 秒
config.BackoffFactor = 2.0                   // 指数因子 2.0
```

### 速率限制 (Rate Limiting)

基于权重的速率限制，符合 WEEX API 规范：

```go
config := weex.NewDefaultConfig()
config.EnableRateLimit = true
config.IPWeight = 300   // IP 权重限制：300/5分钟
config.UIDWeight = 100  // UID 权重限制：100/5分钟
```

### 日志 (Logging)

支持自定义日志实现：

```go
// 使用默认logger
config.Logger = weex.NewDefaultLogger(weex.LogLevelInfo)

// 使用自定义logger（实现 weex.Logger 接口）
type MyLogger struct{}

func (l *MyLogger) Debug(msg string, args ...interface{}) { /* ... */ }
func (l *MyLogger) Info(msg string, args ...interface{}) { /* ... */ }
func (l *MyLogger) Warn(msg string, args ...interface{}) { /* ... */ }
func (l *MyLogger) Error(msg string, args ...interface{}) { /* ... */ }
func (l *MyLogger) SetLevel(level weex.LogLevel) { /* ... */ }

config.Logger = &MyLogger{}
```

## 类型系统

### 枚举类型

SDK 提供了强类型的枚举定义：

```go
// 保证金模式
types.MarginModeShared   // 全仓
types.MarginModeIsolated // 逐仓

// 订单类型
types.OrderTypeOpenLong   // 开多
types.OrderTypeOpenShort  // 开空
types.OrderTypeCloseLong  // 平多
types.OrderTypeCloseShort // 平空

// 订单执行类型
types.OrderExecNormal            // 普通委托
types.OrderExecPostOnly          // 只做 maker
types.OrderExecFillOrKill        // FOK
types.OrderExecImmediateOrCancel // IOC

// 价格类型
types.PriceMatchLimit  // 限价
types.PriceMatchMarket // 市价
```

### 精度处理

所有小数使用 `Decimal` 类型（字符串），避免浮点数精度损失：

```go
price := types.NewDecimal(50000.5)
size := types.NewDecimalFromString("1.5")

// 转换为 float64（需要时）
priceFloat, err := price.Float64()
```

## 目录结构

```
sdk/golang/
├── weex/                    # 主包
│   ├── auth.go              # 认证和签名
│   ├── config.go            # 配置系统
│   ├── errors.go            # 错误处理
│   ├── logger.go            # 日志接口
│   ├── retry.go             # 重试机制
│   ├── rate_limiter.go      # 速率限制
│   ├── types/               # 通用类型
│   │   ├── common.go        # 枚举和常量
│   │   └── errors.go        # 错误码映射
│   ├── rest/                # REST API
│   │   ├── market/          # 市场 API
│   │   ├── account/         # 账户 API
│   │   └── trade/           # 交易 API
│   └── websocket/           # WebSocket API
│       ├── public/          # 公开频道
│       └── private/         # 私有频道
├── examples/                # 示例代码
├── docs/                    # 文档
└── tests/                   # 测试
```

## API 文档

详细的 API 文档请参考：
- [快速入门指南](docs/QUICKSTART.md)
- [认证指南](docs/AUTHENTICATION.md)
- [错误处理指南](docs/ERROR_HANDLING.md)
- [WebSocket 指南](docs/WEBSOCKET.md)

## 贡献

欢迎贡献！请查看 [贡献指南](CONTRIBUTING.md)。

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件。

## 支持

- 问题反馈: [GitHub Issues](https://github.com/weex/openapi-contract-go-sdk/issues)
- API 文档: [WEEX API Documentation](https://www.weex.com/api-doc/)
- 官方网站: [https://www.weex.com](https://www.weex.com)

## 免责声明

本 SDK 为非官方实现。使用本 SDK 进行交易时，请自行承担风险。作者不对使用本 SDK 造成的任何损失负责。

---

**⚠️ 风险警告**: 加密货币交易存在重大风险。请确保您充分理解这些风险并谨慎交易。
