# 插件集成模型

## 角色

```text
┌─────────────┐     registers      ┌──────────────────┐
│   Plugin    │ ─────────────────► │  Host (Registry) │
│  (天气/计时) │ ◄── callbacks ──── │  apple-dog CLI   │
└─────────────┘                    └────────┬─────────┘
                                            │
                                            │ selects Primary
                                            ▼
                                    ┌───────────────┐
                                    │ Island snapshot│
                                    │ Compact/Expanded│
                                    └───────────────┘
```

| 角色 | 职责 |
|------|------|
| **Plugin** | 声明元数据、优先级；在 `Start`/`Stop` 中管理自身资源；通过 `Snapshot()` 提供当前岛布局。 |
| **Host** | 注册表、生命周期、**主展示插件**选举、冲突日志；未来 iOS 侧负责 `Activity.request/update/end`。 |
| **Registry** | 插件 ID 唯一；禁止重复注册。 |

## 插件接口（Go 草稿 → 未来 Swift 对齐）

见 `internal/plugin/plugin.go`：

- `Descriptor()` — id、名称、优先级、预期 Live Activity 时长说明。
- `Start(ctx, HostCallbacks)` — 激活；可注册定时刷新（CLI 用 `HostCallbacks.RequestRefresh`）。
- `Stop(ctx)` — 释放。
- `Snapshot()` — 返回 `island.Presentation`（compact + expanded + minimal + companion）。

## 生命周期

```text
        Register()
            │
            ▼
      ┌───────────┐
      │ Registered │  在 Registry 中，未 Start
      └─────┬─────┘
            │ Host.Activate(id)
            ▼
      ┌───────────┐
      │  Active   │  可参与 Primary 选举
      └─────┬─────┘
            │ Host.Deactivate(id) / Stop
            ▼
      ┌───────────┐
      │  Stopped  │  仍注册，可再次 Activate
      └───────────┘
```

- **Registered**：已 `Register`，未运行。
- **Active**：已 `Start`；多个 Active 时，**Primary** = 最高优先级（数值小者优先，见 stub）。
- iOS 映射：`Activate` ≈ 请求 Live Activity；`Deactivate` ≈ `Activity.end`（简化）。

## 主展示（Primary）选举

1. 收集所有 **Active** 插件的 `Snapshot()`。
2. 按 `Descriptor.Priority` 升序选第一个为非空 compact 的插件（可配置策略，当前为默认）。
3. 其余 Active 插件可贡献 **Minimal** glyph（草案：仅最高两个，TODO 在 Swift 实现）。

## 注册内容到「岛」

插件 **不** 直接写 UI；只返回 **声明式** `island.Presentation`：

```go
type Presentation struct {
    Compact   CompactLayout
    Minimal   MinimalLayout
    Expanded  ExpandedLayout
    Companion CompanionLayout
}
```

宿主职责：

1. 校验槽位（文本长度软限制，见 `island.Validate`）。
2. 合并 Primary 与 Minimal 列表。
3. （iOS）映射为 SwiftUI；**（Go CLI）** 调用 `render.Wireframe`。

## 内置示例插件

| ID | 包路径 | 用途 |
|----|--------|------|
| `weather` | `internal/plugins/weather` | 紧凑温度 + 展开时段 |
| `timer` | `internal/plugins/timer` | 倒计时紧凑 + 展开控制 |

注册入口：`internal/plugins.Builtin()`。

## 占位与后续工作

- [ ] JSON Schema 导出 `Presentation`，供 Swift 代码生成。
- [ ] ActivityKit Push 载荷结构（仅文档）。
- [ ] 插件沙箱与签名（产品化）。
- [ ] 真实优先级 / 用户手动 pin 主插件。
