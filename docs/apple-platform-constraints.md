# Apple 平台约束（Dynamic Island / Live Activity）

> 以下为架构设计依据；实现细节以 [Apple Developer Documentation](https://developer.apple.com/documentation/activitykit) 为准。

## Dynamic Island 与 Live Activity

- **Dynamic Island** 是硬件裁切区域上的系统 UI；第三方内容主要通过 **Live Activity** 呈现，而不是随意绘制系统级 overlay。
- **ActivityKit**（iOS 16.1+）提供 Live Activity 的生命周期、UI 更新与 **Dynamic Island 布局**（compact / minimal / expanded）。
- Live Activity 需在 **Xcode 工程** 中配置 `ActivityAttributes`，用 **SwiftUI** 定义 `ActivityConfiguration` 内的 `dynamicIsland { ... }` 闭包。
- 更新方式：App 内 `Activity.update`、以及符合规则的 **推送**（ActivityKit push）；不能从 Go 或后台任意进程直接写岛 UI。

## WidgetKit 的关系

- **WidgetKit** 驱动主屏幕小组件与 Lock Screen 小组件，API 与 Live Activity **不同**。
- 灵动岛拓维可将 Widget 作为 **伴生入口**（例如点击小组件打开宿主 App 并激活某插件），但 **Widget 本身不是 Dynamic Island**。
- 本仓库在 `CompanionLayout` 中预留「主屏/锁屏伴生入口」字段，仅作产品占位；Go 侧不实现 WidgetKit。

## 展示形态（与 `internal/island` 对齐）

| 系统形态 | Apple 含义 | 本仓库模型 |
|----------|------------|------------|
| **Compact leading / trailing** | 药丸两侧紧凑视图 | `CompactLayout.Leading` / `Trailing` |
| **Minimal** | 多活动时的最小指示 | `MinimalLayout`（占位，后续 Swift 映射） |
| **Expanded** | 长按/点击展开 | `ExpandedLayout`（分区 `Region`） |
| **Lock Screen banner** | 锁屏 Live Activity 条 | 文档占位，未建模细节 |

## 必须原生 iOS 的部分

- Swift/SwiftUI UI、ActivityKit、`DynamicIsland` 视图构建。
- App Groups、推送证书、Background Modes、隐私清单与 App Store 审核材料。
- 真机交互（触觉、深度链接、Siri / Shortcuts 若接入）。

## 本 Go 仓库可原型化的部分

- 插件 ID、优先级、互斥与生命周期状态机。
- 紧凑/展开/伴生屏的 **结构化布局描述**（后续可 JSON 化给 Swift 生成器消费）。
- 宿主 **编排逻辑**（同时多个插件注册，但岛上「主展示」选一个）的单元测试。
- CLI **线框预览**，便于 PR review 与产品 markup。

## 政策与体验约束（设计时遵守）

- Live Activity 默认 **8 小时**上限（部分场景可延长，以 Apple 文档为准）；插件应声明预期时长与结束条件。
- 避免在紧凑态塞入过多文本；展开态承载主要信息与 **少量** 操作。
- 敏感数据遵守锁屏可见性设置；插件需提供「锁屏脱敏」占位策略（Swift 实现）。
