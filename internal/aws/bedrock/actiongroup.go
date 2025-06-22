package bedrock

import "fmt"

// https://docs.aws.amazon.com/bedrock/latest/userguide/agents-lambda.html

type ActionGroupRequest struct {
	MessageVersion          string                        `json:"messageVersion,omitempty"`
	Agent                   ActionGroupRequestAgent       `json:"agent,omitempty"`
	InputText               string                        `json:"inputText,omitempty"`
	SessionID               string                        `json:"sessionId,omitempty"`
	ActionGroup             string                        `json:"actionGroup,omitempty"`
	APIPath                 string                        `json:"apiPath,omitempty"`
	HTTPMethod              string                        `json:"httpMethod,omitempty"`
	Parameters              []ActionGroupRequestParameter `json:"parameters,omitempty"`
	RequestBody             ActionGroupRequestRequesyBody `json:"requestBody,omitempty"`
	SessionAttributes       map[string]string             `json:"sessionAttributes,omitempty"`
	PromptSessionAttributes map[string]string             `json:"promptSessionAttributes,omitempty"`
}

type ActionGroupRequestAgent struct {
	Name    string `json:"name,omitempty"`
	ID      string `json:"id,omitempty"`
	Alias   string `json:"alias,omitempty"`
	Version string `json:"version,omitempty"`
}

type ActionGroupRequestParameter struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type ActionGroupRequestRequesyBody struct {
	Content map[string]ActionGroupRequestRequesyBodyContentProperties `json:"content,omitempty"`
}

type ActionGroupRequestRequesyBodyContentProperties struct {
	Properties []ActionGroupRequestRequestBodyContentProperty `json:"properties,omitempty"`
}

type ActionGroupRequestRequestBodyContentProperty struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type ActionGroupResponse struct {
	MessageVersion          string                      `json:"messageVersion,omitempty"`
	Response                ActionGroupResponseResponse `json:"response,omitempty"`
	SessionAttributes       map[string]string           `json:"sessionAttributes,omitempty"`
	PromptSessionAttributes map[string]string           `json:"promptSessionAttributes,omitempty"`
	// KnowledgeBaseConfiguration is not supported yet.
}

type ActionGroupResponseResponse struct {
	ActionGroup    string                                             `json:"actionGroup,omitempty"`
	APIPath        string                                             `json:"apiPath,omitempty"`
	HTTPMethod     string                                             `json:"httpMethod,omitempty"`
	HTTPStatusCode int                                                `json:"httpStatusCode,omitempty"`
	ResponseBody   map[string]ActionGroupResponseResponseResponseBody `json:"responseBody,omitempty"`
}

type ActionGroupResponseResponseResponseBody struct {
	Body string `json:"body,omitempty"`
}

func NewActionGroupResponse(event *ActionGroupRequest, httpStatusCode int, responseBody map[string]ActionGroupResponseResponseResponseBody) *ActionGroupResponse {
	return &ActionGroupResponse{
		MessageVersion: "1.0",
		Response: ActionGroupResponseResponse{
			ActionGroup:    event.ActionGroup,
			APIPath:        event.APIPath,
			HTTPMethod:     event.HTTPMethod,
			HTTPStatusCode: httpStatusCode,
			ResponseBody:   responseBody,
		},
		SessionAttributes:       event.SessionAttributes,
		PromptSessionAttributes: event.PromptSessionAttributes,
	}
}

func (r *ActionGroupRequest) GetParameter(name string) (string, error) {
	for _, p := range r.Parameters {
		if p.Name == name {
			return p.Value, nil
		}
	}
	return "", fmt.Errorf("parameter '%s' was not found", name)
}
