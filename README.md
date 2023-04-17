# Open Assistant Helper

English | [简体中文](./README_zh.md)

Accelerate your contribution to [Open Assistant](https://open-assistant.io).

## Features

This project feeds tasks from Open Assistant to ChatGPT (GPT 3.5), along with corresponding prompts to help it complete the tasks. Then the replies will be sent back to Open Assistant.

## Usage

- Go to [releases](https://github.com/AnotiaWang/open-assistant-helper/releases/latest) and download the latest version, according to you OS (Windows / Linux / macOS).
- Run the executable file. The program will ask you for your [OpenAI Secret Key](https://platform.openai.com/account/api-keys), and your Open Assistant Cookie.
  > You can find your Open Assistant Cookie in the browser's developer tools. 
  > 
  > Click on the `Network` tab, and then refresh the page. Find the request called `dashboard` at the top, and click on it. Then copy the `cookie` field in the `Request Headers` section.

## Build

- Clone code
- Install [Go](https://go.dev/dl/) (1.20 recommended)
- `go build -o oa-helper .`

## Plans

- [ ] Support task `label_initial_prompt`
  
  This task is a bit hard to find in my language configuration, I'll keep finding it.
- [ ] Support other language models
  
  There have been many excellent LLMs coming out recently, like GPT4All. I may try to support them in my free time(but no guarantees). if you have ideas, PRs are welcome!

## Note

For those under special network conditions (for example, ChatGPT is blocked in your region), you can simply set up proxy if you have a proxy that supports HTTP/HTTPS port proxy.
Open your terminal, and run the following command:

```bash
export http_proxy=http://your-proxy:port
export https_proxy=http://your-proxy:port
```

Then run the executable file in the same terminal session.
