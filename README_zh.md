# Open Assistant Helper

[English](./README.md) | 简体中文

帮你对 [Open Assistant](https://open-assistant.io) 的贡献加把火 🔥。

## 功能

此项目将 Open Assistant 的各种任务发送到 ChatGPT（GPT 3.5），并提供相应的提示来帮助它完成任务。然后会把回复发送回 Open Assistant。（注意，目前是要花钱的）

## 使用方法

- 转到 [Release 页面](https://github.com/AnotiaWang/open-assistant-helper/releases/latest) 并根据您的操作系统（Windows / Linux / macOS）下载最新版本。
- 运行可执行文件。程序将要求您输入：
  - [OpenAI 密钥](https://platform.openai.com/account/api-keys) 
  - 您的 Open Assistant Cookie
    >   **Note** 您可以在浏览器的开发者工具中获取您的 Open Assistant Cookie。
    >
    >   单击 `Network` 标签，然后刷新页面。在 Network 顶部找到名为 `dashboard` 的请求，单击它。然后在 `Request Headers` 部分复制 `cookie` 字段。
  - 任务的首选语言（语言代码，例如 `zh` 为中文）。

配置文件保存在 `config.json` 中。

## 构建

- 克隆代码
- 安装 [Go](https://go.dev/dl/) (推荐 1.20)
- `go build -o oa-helper .`

## 计划

- [x] ~~支持任务 `label_initial_prompt`~~
- [ ] 支持其他语言模型（或许）

  最近出现了许多优秀的 LLM，比如 GPT4All，我有空的话可能会尝试支持它们（但不能保证）。如果您有想法，欢迎提 PR！
- [ ] ...

## 注意事项

对于在特殊网络条件下的用户（例如，您所在区域封锁了 ChatGPT），如果您有支持 HTTP/HTTPS 端口代理的代理，可以打开终端，运行以下命令：

```bash
export http_proxy=http://your-proxy:port
export https_proxy=http://your-proxy:port
```

然后在相同会话中执行此程序。
