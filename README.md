# apple-dog · 灵动岛拓维（Dynamic Island expansion）

Go 模块 `github.com/liucong009/apple-dog`：**灵动岛拓维** 的架构草稿与宿主模拟器。真实 Dynamic Island UI 必须在 iOS（SwiftUI + ActivityKit）实现；本仓库提供可评审的 **插件模型、布局数据结构、ASCII 线框预览**。

## 文档（建议 review 顺序）

| 文档 | 内容 |
|------|------|
| [docs/product-concept.md](docs/product-concept.md) | 产品概念、Go vs iOS 分工 |
| [docs/apple-platform-constraints.md](docs/apple-platform-constraints.md) | Dynamic Island / Live Activity / WidgetKit 约束 |
| [docs/ui-layout.md](docs/ui-layout.md) | Compact / Expanded / Companion 线框 |
| [docs/plugin-integration.md](docs/plugin-integration.md) | 插件接口、生命周期、Primary 选举 |

## 仓库布局

```text
.
├── cmd/apple-dog/           # CLI 宿主模拟器（线框预览）
├── internal/
│   ├── island/              #  declarative 岛 + 伴生屏布局
│   ├── plugin/              # Plugin 接口、Registry、Host
│   ├── plugins/             # 内置 stub：weather、timer
│   └── render/              # ASCII wireframe 渲染
├── docs/                    # 产品与平台说明
├── go.mod
└── Makefile
```

## 要求

- Go 1.22+

## 构建与测试

```bash
go build ./...
go test ./...
# 或
make build
make test
```

## 运行草稿（CLI）

列出插件：

```bash
go run ./cmd/apple-dog -list
```

预览单个插件（默认 `weather`，模式 `full` = compact + expanded + minimal）：

```bash
go run ./cmd/apple-dog -plugin weather -mode full
go run ./cmd/apple-dog -plugin timer -mode compact
go run ./cmd/apple-dog -plugin weather -mode companion
```

多插件激活（Primary 按 **priority 数字小者优先**）：

```bash
go run ./cmd/apple-dog -activate weather,timer -mode full
```

构建二进制：

```bash
make build
./bin/apple-dog -list
```

## 仍为占位的内容

- 无 Swift / Xcode 工程；无 ActivityKit、APNs、WeatherKit。
- 插件数据为 **静态 stub**（计时器用相对 deadline，仅演示刷新形态）。
- `Minimal` 多活动合并策略、JSON Schema 导出、用户 pin 主插件均未实现。
- `internal/island.Validate` 仅为草案长度检查，非 Apple HIG 完整校验。

## 扩展插件（开发向）

1. 在 `internal/plugins/<name>/` 实现 `plugin.Plugin`。
2. 在 `internal/plugins/register.go` 的 `RegisterBuiltin` 中注册。
3. 为 `Snapshot()` 与 Primary 选举添加测试。

## Module path

`github.com/liucong009/apple-dog`
