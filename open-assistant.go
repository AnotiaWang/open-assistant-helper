package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"open-assistant-helper-go/model"
	"strings"
)

var rty = resty.New()

func CancelTask(id string) error {
	logx.Infof("CancelTask")
	_, err := rty.R().
		SetHeaders(map[string]string{
			"cookie":       model.Conf.OaCookie,
			"Content-Type": "application/json",
		}).
		SetBody(map[string]string{
			"id": id,
		}).
		Post("https://open-assistant.io/api/reject_task")
	if err != nil {
		return err
	}
	return nil
}

func RefreshCookie() error {
	logx.Infof("RefreshCookie")
	resp, err := rty.R().
		SetHeaders(map[string]string{
			"cookie": model.Conf.OaCookie,
		}).
		Get("https://open-assistant.io/api/auth/session")
	if err != nil {
		return err
	}
	cs := resp.Header()["Set-Cookie"]
	if len(cs) <= 1 {
		return fmt.Errorf("no cookie. Please consider going to https://open-assistant.io , login, and update your cookie in config.json")
	}
	c := strings.Join(cs, "")
	return model.UpdateCookie(c)
}

func GetLabelsFromChatGPT(resp string) (map[model.OALabel]float32, error) {
	var j map[string]interface{}
	err := jsonx.Unmarshal([]byte(resp), &j)
	if err != nil {
		return nil, err
	}
	labels := make(map[model.OALabel]float32)
	for k, v := range j {
		if k == "text" {
			continue
		}
		f, _ := v.(json.Number).Float64()
		labels[model.OALabel(k)] = float32(f)
	}
	return labels, nil
}
func LabelAssistantReply(id string, task model.OALabelAssistantReplyTask) error {
	logx.Infof("LabelAssistantReply")
	text := ""

	for _, m := range task.Conversation.Messages {
		if m.IsAssistant {
			text += "Assistant"
		} else {
			text += "User"
		}
		text += fmt.Sprintf(": %s\n\n", m.Text)
	}
	if task.Reply != "" {
		text += fmt.Sprintf("Assistant: %s", task.Reply)
	}

	t, err := Complete(text, `You are a fine-tune tool for Open Assistant, a LLM. You will be given conversations between a user and a model, and you need to label the model's reply and return a JSON string. You should evaluate the conversations based on the following criteria:
- Spam: 0/1, whether the conversation contains spam / ads / porn / politics / etc.
- Fails Task: 0/1, whether the response is strongly related to the user's question
- Lang Mismatch: 0/1, whether the most part of response is in the same language as the user's question
- Not Appropriate: 0/1, whether the response is reasonable for the user's question
- pii: 0/1
- Hate Speech: 0/1, whether the response is aggressive / not respectful
- Sexual Content: 0/1
- Quality: 0~1, step 0.25, how well the response is written respecting grammar, spelling, use of words, etc.
- Helpfulness: 0~1, step 0.25
- Creativity: 0~1, step 0.25, how less is the model fixed to the user's question
- Humor: 0~1, step 0.25
- Toxicity: 0~1, step 0.25, how aggressive is the response
- Violence: 0~1, step 0.25

You MUST reply a JSON string, DO NOT include any other characters, DO NOT explain. Use snake_case for the keys.`)
	if err != nil {
		return err
	}

	logx.Infof("LabelAssistantReply: %s", t)
	labels, err := GetLabelsFromChatGPT(t)
	if err != nil {
		return err
	}
	body := model.PostBodyLabelAssistant{
		Id:         id,
		Lang:       model.Conf.Language,
		UpdateType: "text_labels",
		Content: model.PostBodyLabelAssistantContent{
			Labels:    labels,
			MessageId: task.MessageId,
			Text:      "unused?",
		},
	}
	resp, err := rty.R().
		SetHeaders(map[string]string{
			"Cookie":       model.Conf.OaCookie,
			"Content-Type": "application/json",
		}).
		SetBody(body).
		Post("https://open-assistant.io/api/update_task")
	if err != nil {
		return err
	}

	respStr := string(resp.Body())
	if respStr == "" {
		logx.Infof("LabelAssistantReply: OK!")
	} else {
		logx.Errorf("LabelAssistantReply: %s", respStr)
	}
	return nil
}

func LabelPrompterReply(id string, task model.OALabelPrompterReplyTask) error {
	logx.Infof("LabelPrompterReply")
	text := ""

	for _, m := range task.Conversation.Messages {
		if m.IsAssistant {
			text += "Assistant"
		} else {
			text += "User"
		}
		text += fmt.Sprintf(": %s\n\n", m.Text)
	}
	if task.Reply != "" {
		text += fmt.Sprintf("User: %s", task.Reply)
	}

	t, err := Complete(text, `You are a powerful fine-tuner of Open Assistant, an open-source LLM. You will be given conversations between a user and the model, and you need to label the user's last reply and return a JSON string. You should evaluate the conversations based on the following criteria:
- Spam: 0/1, whether the conversation contains spam / ads / porn / politics / etc.
- Not Appropriate: 0/1, whether the response is reasonable for the user's question
- pii: 0/1
- Hate Speech: 0/1, whether the response is aggressive / not respectful
- Sexual Content: 0/1
- Quality: 0-1, step 0.25, how well the response is written respecting grammar, spelling, use of words, etc.
- Lang Mismatch: 0-1, step 0.25, whether the most part of response is in the same language as the user's question
- Creativity: 0-1, step 0.25, how less is the model fixed to the user's question
- Humor: 0-1, step 0.25
- Toxicity: 0-1, step 0.25, how aggressive is the response
- Violence: 0-1, step 0.25

You must return a JSON string, DO NOT include any other characters, DO NOT explain. Use snake_case for the keys.`)
	if err != nil {
		return err
	}

	logx.Infof("LabelPrompterReply: %s", t)
	labels, err := GetLabelsFromChatGPT(t)
	if err != nil {
		return err
	}

	body := model.PostBodyLabelPrompter{
		Id:         id,
		Lang:       model.Conf.Language,
		UpdateType: "text_labels",
		Content: model.PostBodyLabelPrompterContent{
			Labels:    labels,
			MessageId: task.MessageId,
			Text:      "unused?",
		},
	}
	tmp, _ := jsonx.MarshalToString(body)
	logx.Infof(tmp)
	resp, err := rty.R().
		SetHeaders(map[string]string{
			"Cookie":       model.Conf.OaCookie,
			"Content-Type": "application/json",
		}).
		SetBody(body).
		Post("https://open-assistant.io/api/update_task")
	if err != nil {
		return err
	}

	respStr := string(resp.Body())
	if respStr == "" {
		logx.Infof("LabelPrompterReply: OK!")
	} else {
		logx.Errorf("LabelPrompterReply: %s", respStr)
	}
	return nil
}

func LabelInitialPrompt(id string, task model.OALabelInitialPromptTask) error {
	logx.Infof("LabelInitialPrompt")
	text := fmt.Sprintf("Prompt: %s\n\nUser's language code: %s", task.Prompt, model.Conf.Language)

	t, err := Complete(text, `You are a powerful fine-tuner of Open Assistant, an open-source LLM. You will be given a prompt from the user, and you need to label it and return a JSON string. You should evaluate the prompt based on the following criteria:
- Spam: 0/1, whether the message contains spam / ads / porn / politics / etc.
- Not Appropriate: 0/1, whether the message is offensive / not respectful
- pii: 0/1
- Hate Speech: 0/1, whether the prompt is aggressive / not respectful
- Sexual Content: 0/1
- Quality: 0-1, step 0.25, how well the response is written respecting grammar, spelling, use of words, etc.
- Lang Mismatch: 0-1, step 0.25, whether the prompt is in the same language as the user's language
- Creativity: 0-1, step 0.25, how less is the prompt
- Humor: 0-1, step 0.25
- Toxicity: 0-1, step 0.25, how aggressive is the prompt
- Violence: 0-1, step 0.25

You must return a JSON string, DO NOT include any other characters, DO NOT explain. Use snake_case for the keys.`)

	if err != nil {
		return err
	}

	logx.Infof("LabelInitialPrompt: %s", t)
	labels, err := GetLabelsFromChatGPT(t)
	if err != nil {
		return err
	}

	body := model.PostBodyLabelInitialPrompt{
		Id:         id,
		Lang:       model.Conf.Language,
		UpdateType: "text_labels",
		Content: model.PostBodyLabelInitialPromptContent{
			Labels:    labels,
			Text:      "unused?",
			MessageId: task.MessageId,
		},
	}

	resp, err := rty.R().
		SetHeaders(map[string]string{
			"Cookie":       model.Conf.OaCookie,
			"Content-Type": "application/json",
		}).
		SetBody(body).
		Post("https://open-assistant.io/api/update_task")

	if err != nil {
		return err
	}

	respStr := string(resp.Body())
	if respStr == "" {
		logx.Infof("LabelInitialPrompt: OK!")
	} else {
		logx.Errorf("LabelInitialPrompt: %s", respStr)
	}
	return nil
}

func AssistantReply(id string, task model.OAAssistantReplyTask) error {
	logx.Infof("AssistantReply")
	text := ""

	for _, m := range task.Conversation.Messages {
		if m.IsAssistant {
			text += "Assistant"
		} else {
			text += "User"
		}
		text += fmt.Sprintf(": %s\n\n", m.Text)
	}

	t, err := Complete(text, `You are a powerful fine-tuner of Open Assistant, an open-source LLM. You will be given conversations between a user and a model, and you need to generate an appropriate response as the model. You MUST return the response in plain text. DO NOT provide any other information, DO NOT explain.`)
	if err != nil {
		return err
	}

	logx.Infof("AssistantReply: %s", t)
	body := model.PostBodyAssistant{
		Id:         id,
		Lang:       model.Conf.Language,
		UpdateType: "text_reply_to_message",
		Content: model.PostBodyAssistantContent{
			Text: t,
		},
	}
	resp, err := rty.R().
		SetHeaders(map[string]string{
			"Cookie":       model.Conf.OaCookie,
			"Content-Type": "application/json",
		}).
		SetBody(body).
		Post("https://open-assistant.io/api/update_task")
	if err != nil {
		return err
	}

	respStr := string(resp.Body())
	if respStr == "" {
		logx.Infof("AssistantReply: OK!")
	} else {
		logx.Errorf("AssistantReply: %s", respStr)
	}

	return nil
}

func PrompterReply(id string, task model.OAPrompterReplyTask) error {
	logx.Infof("PrompterReply")
	text := ""

	for _, m := range task.Conversation.Messages {
		if m.IsAssistant {
			text += "Assistant"
		} else {
			text += "User"
		}
		text += fmt.Sprintf(": %s\n\n", m.Text)
	}

	t, err := Complete(text, `You are a powerful fine-tuner of Open Assistant, an open-source LLM. You will be given conversations between a user and the model, and you need to generate an appropriate response as the user. You MUST return the response in plain text. DO NOT provide any other information, DO NOT explain.`)
	if err != nil {
		return err
	}

	logx.Infof("PrompterReply: %s", t)
	body := model.PostBodyPrompter{
		Id:         id,
		Lang:       model.Conf.Language,
		UpdateType: "text_reply_to_message",
		Content: model.PostBodyPrompterContent{
			Text: t,
		},
	}

	resp, err := rty.R().
		SetHeaders(map[string]string{
			"Cookie":       model.Conf.OaCookie,
			"Content-Type": "application/json",
		}).
		SetBody(body).
		Post("https://open-assistant.io/api/update_task")
	if err != nil {
		return err
	}

	respStr := string(resp.Body())
	if respStr == "" {
		logx.Infof("PrompterReply: OK!")
	} else {
		logx.Errorf("PrompterReply: %s", respStr)
	}

	return nil
}

func RankAssistantReplies(id string, task model.OARankAssistantRepliesTask) error {
	logx.Infof("RankAssistantReplies")
	text := ""

	for _, m := range task.Conversation.Messages {
		if m.IsAssistant {
			text += "Assistant"
		} else {
			text += "User"
		}
		text += fmt.Sprintf(": %s\n\n", m.Text)
	}

	text += "Choices to be ranked:\n\n"
	for i, r := range task.Replies {
		text += fmt.Sprintf("%d: %s\n", i, r)
	}

	t, err := Complete(text, `You are a powerful fine-tuner of Open Assistant, an open-source LLM. You will be given conversations between a user and the model, and you need to rank the replies of the model, according to their quality, precision and readability, and return a JSON string.
Give the ranking in JSON format, DO NOT include any other characters, DO NOT explain. If all the replied are not appropriate, set 'not_rankable' to false. For example: {"not_rankable": false, "ranking": [0, 1, 2]}`)
	if err != nil {
		return err
	}

	logx.Infof("RankAssistantReplies: %s", t)
	var content model.PostBodyRankAssistantContent
	err = jsonx.UnmarshalFromString(t, &content)
	if err != nil {
		return err
	}
	body := model.PostBodyRankAssistant{
		Id:         id,
		Lang:       model.Conf.Language,
		UpdateType: "message_ranking",
		Content:    content,
	}

	resp, err := rty.R().
		SetHeaders(map[string]string{
			"Cookie":       model.Conf.OaCookie,
			"Content-Type": "application/json",
		}).
		SetBody(body).
		Post("https://open-assistant.io/api/update_task")
	if err != nil {
		return err
	}

	respStr := string(resp.Body())
	if respStr == "" {
		logx.Infof("RankAssistantReplies: OK!")
	} else {
		logx.Errorf("RankAssistantReplies: %s", respStr)
	}

	return nil
}

func StartTask() error {
	logx.Infof("GetTasks")
	resp, err := rty.R().
		SetHeaders(map[string]string{
			"Cookie":       model.Conf.OaCookie,
			"Content-Type": "application/json",
		}).
		Get("https://open-assistant.io/api/new_task/random?lang=" + model.Conf.Language)
	if err != nil {
		return err
	}

	var task model.OARandomTaskResponse
	body := resp.Body()
	err = jsonx.Unmarshal(body, &task)
	if resp.StatusCode() == 403 {
		logx.Errorf("GetTasks: cookie may have expired (403 Forbidden)")
		logx.Errorf("Please login to https://open-assistant.io/dashboard and update your cookie in config.json")
		return nil
	} else if task.Task == nil {
		var _json map[string]interface{}
		_ = jsonx.Unmarshal(body, &_json)
		if strings.Contains(_json["message"].(string), "No tasks") {
			logx.Infof("GetTasks: no tasks at this time")
		}
		logx.Errorf("GetTasks: get task failed: %s", _json["message"])
		logx.Errorf("Please check your network, or login to https://open-assistant.io/dashboard and update your cookie in config.json")
		return nil
	}

	t := task.Task.(map[string]interface{})
	j, _ := jsonx.Marshal(t)
	logx.Infof("Got task, id: %s, type: %s", task.Id, t["type"])
	if t["type"] == "label_assistant_reply" {
		var ch model.OALabelAssistantReplyTask
		_ = jsonx.Unmarshal(j, &ch)
		return LabelAssistantReply(task.Id, ch)
	} else if t["type"] == "label_prompter_reply" {
		var ch model.OALabelPrompterReplyTask
		_ = jsonx.Unmarshal(j, &ch)
		return LabelPrompterReply(task.Id, ch)
	} else if t["type"] == "assistant_reply" {
		var ch model.OAAssistantReplyTask
		_ = jsonx.Unmarshal(j, &ch)
		return AssistantReply(task.Id, ch)
	} else if t["type"] == "prompter_reply" {
		var ch model.OAPrompterReplyTask
		_ = jsonx.Unmarshal(j, &ch)
		return PrompterReply(task.Id, ch)
	} else if t["type"] == "rank_assistant_replies" {
		var ch model.OARankAssistantRepliesTask
		_ = jsonx.Unmarshal(j, &ch)
		return RankAssistantReplies(task.Id, ch)
	} else if t["type"] == "label_initial_prompt" {
		var ch model.OALabelInitialPromptTask
		_ = jsonx.Unmarshal(j, &ch)
		return LabelInitialPrompt(task.Id, ch)
	} else {
		logx.Infof("GetTasks: unknown task type: %s", t["type"])
		err = CancelTask(task.Id)
		if err != nil {
			return fmt.Errorf("GetTasks: cancel task failed: %s", err)
		}
	}

	return nil
}
