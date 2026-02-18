# Changelog

## 0.3.2 - 2026-02-18

### Fixed

- 修复双节点隧道（`type=2`）在节点重装后仅收到 `UpdateChains` 时无法自动恢复 chain 的问题：`UpdateChains` 失败 `not found` 时会自动回退执行 `AddChains`。
- 修复 outbox 仅按“已发送”判定成功的问题：改为等待节点执行响应（ACK）后再标记 `done`，并对幂等场景（如 `Add* already exists`、`Delete* not found`）按成功处理，避免误重试或阻塞队列。

## 0.3.1 - 2026-02-18

### Fixed

- 节点重装后本地无服务配置时，`UpdateService` 失败会自动回退执行 `AddService`，确保节点能重新拉起转发配置。
- 该修复覆盖“面板中节点与转发已存在，但节点清空本地配置后重新安装”的场景。

