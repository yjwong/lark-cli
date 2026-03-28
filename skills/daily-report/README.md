# Daily Report Skill

A Claude Code skill that generates structured Lark/Feishu work reports as message cards — built from your task list and chat messages.

**Five report types:** morning briefing · evening wrap-up · quick scan · full snapshot · weekly summary

---

## For Humans

### What You Get

| Report | Default Schedule (UTC+8) | What it answers |
|--------|--------------------------|-----------------|
| Morning | 10:30 Mon–Fri | Top tasks today + overnight messages requiring action |
| Evening | 22:00 Mon–Fri | What shipped today + what to push tomorrow |
| Quick scan | On demand | Anything new since last check — @mentions + task updates |
| Full | On demand | Comprehensive snapshot across all sources |
| Weekly | Sun 15:00 | What I accomplished this week + what's next |

Morning and evening reports accumulate into a rolling weekly log. The weekly summary reads that log — no redundant API calls.

All output is formatted as a Lark message card, sent directly to your configured chat.

---

### Installation

**Prerequisites:** [lark-cli](../../README.md) installed and authenticated.

Tell Claude Code:

```
install daily report skill
```

Claude Code will guide you through the setup — scanning your groups, letting you pick which ones to include, and writing your config. No manual YAML editing required.

---

### Using the Skill

Tell Claude Code in English or Chinese:

| Say this | Report |
|----------|--------|
| `morning report` / `日报` | Morning briefing |
| `evening report` / `晚报` | Evening wrap-up |
| `weekly report` / `周报` | Weekly summary |
| `quick scan` | Quick scan |
| `full report` | Full snapshot |

---

### Required Permissions

In your Lark app console at [open.larksuite.com](https://open.larksuite.com) (Feishu: [open.feishu.cn](https://open.feishu.cn)), add the following under **Permissions & Scopes**. Copy and paste directly into the permission configuration:

```json
{
  "scopes": {
    "tenant": [
      "im:message",
      "im:message:send_as_bot",
      "task:task:readonly",
      "task:tasklist:readonly"
    ],
    "user": [
      "im:message:readonly",
      "task:task:readonly",
      "task:tasklist:readonly",
      "offline_access"
    ]
  }
}
```

> **Why both tenant and user scopes?** Message cards are sent as the bot (tenant token). Task and message history APIs require user identity to access personal data.

After adding permissions, re-authenticate:

```bash
lark auth login --add --scopes messages
```

---

## For AI Agents

This skill orchestrates three data sources to produce work reports:

```
Tasks (v2 API)  +  Chat messages  →  Message card report
```

### Core Flow

1. **Task collection** — `tasks list` discovers tasklist GUIDs from response; `tasklists tasks` per GUID fetches all tasks including follower tasks. (`tasklists list` returns error 1470500 — skip it entirely.)
2. **Message scanning** — Delta scan from `.last-scan-ts`. Whitelist groups get full pull; other active groups are discovered via `messages-search`. `@me` search as catchall backstop.
3. **Intelligence layer** — Match messages against open tasks by keyword and sender; extract action items; classify urgency; identify blockers.
4. **Card output** — Render Lark interactive card JSON per template; send via `messages-send --msg-type interactive --as bot`.

### Templates

Five templates in `templates/` — each defines card structure, extraction logic, and output rules:

| File | Report |
|------|--------|
| `morning-report.md` | TOP 3 tasks (DDL today > P0 > overdue) + message-driven action items |
| `evening-report.md` | Completed tasks + tomorrow push list + conversation highlights |
| `quick-scan.md` | @mentions → task-matched updates → group digest. Plain text if no updates. |
| `full-report.md` | Full task hierarchy by section + all active message groups |
| `weekly-summary.md` | Exec-facing (module-grouped, narrative) + internal (task list + tracking) |

Templates are the source of truth for card format and report logic — edit them to change structure, card theme, or extraction behavior.

### What's Configurable in `config.yaml`

| Field | What it controls |
|-------|-----------------|
| `tasks.section_names` | Map section GUIDs to display names (TODO / Ongoing Project / Tracking) |
| `messages.group_filter.whitelist` | Groups to always scan in full |
| `messages.group_filter.blacklist` | Groups to never scan |
| `scope.areas` | Keyword sets used for weekly report topic extraction |
| `report_types.*.schedule` | Cron schedule per report type |
| `output.send_to` | Auto-send destination (user or chat) |

### Full Skill Documentation

See [SKILL.md](SKILL.md) for complete execution flow, API workarounds, section GUID discovery, card formatting rules, and weekly log accumulation logic.

---

## 中文说明

### 这是什么

一个 Claude Code Skill，自动从飞书任务和群消息中生成结构化工作报告，以消息卡片形式直接发送。

**五种报告类型：** 日报 · 晚报 · 快速扫描 · 全量报告 · 周报

日报和晚报会累积到周日志，周报直接读取该日志，不重复拉取数据。所有报告以飞书消息卡片形式输出。

---

### 安装

**前置条件：** 已安装并完成认证的 [lark-cli](../../README.md)。

对 Claude Code 说：

```
install daily report skill
```

Claude 会自动引导你完成配置，包括扫描你的群组、选择要监控的群、写入 config 文件。无需手动编辑 YAML。

---

### 触发方式

直接告诉 Claude Code：

| 说这个 | 生成的报告 |
|--------|-----------|
| `日报` | 早间日报 |
| `晚报` | 晚间总结 |
| `周报` | 周报草稿 |
| `quick scan` | 快速扫描 |
| `full report` | 全量报告 |

---

### 必要权限

在飞书开放平台 [open.feishu.cn](https://open.feishu.cn) 的「权限管理」中添加以下权限。可直接复制使用：

```json
{
  "scopes": {
    "tenant": [
      "im:message",
      "im:message:send_as_bot",
      "task:task:readonly",
      "task:tasklist:readonly"
    ],
    "user": [
      "im:message:readonly",
      "task:task:readonly",
      "task:tasklist:readonly",
      "offline_access"
    ]
  }
}
```

> **为什么需要两类权限？** 消息卡片以 bot 身份发送（tenant token）；任务和消息历史需要用户身份才能读取个人数据。

添加权限后重新授权：

```bash
lark auth login --add --scopes messages
```
