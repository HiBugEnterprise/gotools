package openai

// LLMAnswer
// 模型返回的结果，在
// 非流式的格式中模型返回的消息在 LLMAnswer.choices[index].message.content中
// 流式的格式中模型返回的消息在 LLMAnswer.choices[index].delta.content中
// 非流式
//
//	{
//	   "id": "string",
//	   "object": "string",
//	   "created": 0,
//	   "model": "string",
//	   "choices": [
//	       {
//	           "index": 0,
//	           "message": {
//	               "role": "string",
//	               "content": "string"
//	           },
//	           "finish_reason": "string"
//	       }
//	   ],
//	   "usage": {
//	       "prompt_tokens": 0,
//	       "completion_tokens": 0,
//	       "total_tokens": 0
//	   },
//	   "corr_lst": [
//	       "string"
//	   ],
//	   "code": 0
//	}
//
// 流式
//
//	{
//	   "id": "string",
//	   "object": "string",
//	   "created": 0,
//	   "model": "string",
//	   "choices": [
//	       {
//	           "index": 0,
//	           "delta": {
//	               "role": "string",
//	               "content": "string"
//	           },
//	           "finish_reason": "string",
//	           "usage": {
//	               "prompt_tokens": 0,
//	               "completion_tokens": 0,
//	               "total_tokens": 0
//	           }
//	       }
//	   ],
//	   "corr_lst": [
//	       "string"
//	   ],
//	   "code": 0
//	}
type LLMAnswer struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int64     `json:"created"`
	Model   string    `json:"model"`
	Choices []*Choice `json:"choices"`
	Usage   *Usage    `json:"usage"`
	CorrLst []*Corr   `json:"corr_lst"`
	Code    int64     `json:"code"`
}

type Choice struct {
	Index        int64    `json:"index"`
	Message      *Message `json:"message"` // 非流式消息在此处
	Delta        *Message `json:"delta"`   // 流式消息在此处
	FinishReason string   `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

type Corr struct {
	Content string `json:"content"`
	Score   string `json:"score"`
}

type ChatReq struct {
	Id       string     `json:"id"`       // 传入session_uuid
	Model    string     `json:"model"`    // 使用的是什么模型
	Stream   bool       `json:"stream"`   // 是否是流式接口
	Messages []*Message `json:"messages"` // 历史消息
}
