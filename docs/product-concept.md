# 灵动岛拓维（Dynamic Island expansion）

## 产品概念

**灵动岛拓维** 是一款以 iPhone **Dynamic Island（灵动岛）** 为载体的扩展框架：在系统允许的展示形态内，把更多场景（工具、状态、轻交互）**挂载到岛上**，并通过宿主 App 统一编排插件，而不是每个需求各自做一个独立 Live Activity。

### 目标用户场景（草案）

| 场景 | 岛上展示 | 深度操作 |
|------|----------|----------|
| 天气摘要 | 紧凑态：图标 + 温度；展开态：时段与地点 | 宿主 App「天气插件」设置页 |
| 计时 / 番茄钟 | 紧凑态：剩余时间；展开态：进度与控制 | 展开态按钮或 App 内控制 |
| 未来插件 | 由插件声明 `Compact` / `Expanded` 槽位 | `Companion` 屏配置 |

### 非目标（当前阶段）

- 不在本仓库实现真实 iOS UI（需 SwiftUI + ActivityKit）。
- 不绕过 Apple 对 Live Activity 时长、推送更新、敏感内容的政策。
- 不做「全局注入系统岛」的私有 API 方案（仅面向 **自家 App + Live Activity / ActivityKit** 合法路径）。

### 本 Go 仓库的角色

| 层级 | 本仓库 | 未来 iOS 宿主 |
|------|--------|----------------|
| 插件契约、生命周期 | ✅ 接口 + 注册表 + 测试 | Swift `Plugin` 协议对齐 |
| 布局数据结构 | ✅ `internal/island` | 映射到 `DynamicIsland` DSL |
| 岛上真实渲染 | ❌ ASCII 线框预览 | SwiftUI + ActivityKit |
| Live Activity 推送 | ❌ JSON 占位 | APNs + `ActivityKit` |

本仓库产出 **可评审的架构草稿 + 可运行的宿主模拟器 CLI**，便于在尚无设计稿前对齐布局与插件模型。
