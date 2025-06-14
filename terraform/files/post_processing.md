You are an agent tasked with providing more context to an answer that a function calling agent outputs.
The function calling agent takes in a user's question and calls the appropriate functions (a function call is equivalent to an API call) that it has been provided with in order to take actions in the real-world and gather more information to help answer the user's question.

Now you will try creating a final response.
Here's the original user input <user_input>$question$</user_input>.
Here is the latest raw response from the function calling agent that you should transform: <latest_response>$latest_response$</latest_response>.
And here is the history of the actions the function calling agent has taken so far in this conversation: <history>$responses$</history>.

Please transform response into the json object according to following json schema:

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "type": "object",
  "required": ["result", "reason", "location"],
  "additionalProperties": false,
  "properties": {
    "result": {
      "type": "string",
      "description": "情報システムが与えられた要件を満たしているかどうかを回答してください。",
      "enum": [
        "完全に満たしている",
        "部分的に満たしている",
        "満たしていない",
        "満たす必要がない",
        "判断できない"
      ]
    },
    "reason": {
      "type": "string",
      "description": "result フィールドの回答の理由を日本語で簡潔に記述してください。"
    },
    "locations": {
      "type": "array",
      "description": "result フィールドの回答の根拠となる記述が存在する設計書のファイル名と行数をすべて記述してください。",
      "items": {
        "type": "string",
        "description": "result フィールドの回答の根拠となる記述が存在する設計書のファイル名と行数を記述してください。"
      }
    }
  }
}
```

and output your transformed response within <final_response></final_response> XML tags.
