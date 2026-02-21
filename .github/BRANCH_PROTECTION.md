# Branch Protection（main）

为避免试验性提交直接进入主分支，建议只允许通过 PR 从 `dev` 合并到 `main`，并强制 CI 通过后才能合并。

## 一次性设置

1. 打开仓库 `Settings` -> `Branches` -> `Add branch protection rule`。
2. Branch name pattern 填 `main`。
3. 勾选 `Require a pull request before merging`。
4. 勾选 `Require status checks to pass before merging`。
5. 在状态检查中选择：
   - `go-ci`
   - `frontend-ci`
6. 建议再勾选：
   - `Require branches to be up to date before merging`
   - `Include administrators`

## 推荐分支流程

- 日常开发：提交到 `dev`
- 发布流程：`dev` -> `main` 提 PR
- 合并条件：上述 CI 检查全部通过
