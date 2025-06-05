package awswellarchitectedframework_test

import (
	"testing"

	wa "github.com/shibataka000/heimdall/internal/checklist/awswellarchitectedframework"
	"github.com/stretchr/testify/require"
)

func TestPrompt(t *testing.T) {
	tests := []struct {
		name        string
		requirement wa.Requirement
		prompt      string
	}{{
		name:        "Normal",
		requirement: "GivenRequirement",
		prompt:      "<instruction>\n\nあなたは優秀な IT エンジニアです。\n\nあなたの仕事は、ナレッジベースに格納された情報システムの設計書を参照し、それが与えられた要件を満たしているか否かを判断することです。\n\nあなたは仕事を実行する上で以下のルールを守る必要があります。\n\n- 判断は必ず与えられた情報（情報システムの設計書、情報システムが満たすべき要件）に基づいて行ってください。\n- 情報システムの設計が要件の Common anti-patterns のいずれかに当てはまる場合、情報システムは要件を満たしていないと判断してください。\n\n</instruction>\n\n<input>\n\nあなたには以下の 2 つの情報が与えられます。\n\n- 情報システムの設計書。\n- 情報システムが満たすべき要件。\n\n情報システムの設計書はナレッジベースに格納されています。必要に応じて参照してください。\n\n情報システムが満たすべき要件は以下のとおりです。\n\n<requirement>\n\nGivenRequirement\n\n</requirement>\n\n</input>\n\n<output>\n\n結果は以下のフィールドで構成された json オブジェクトとして出力してください。\n\n| フィールド | 型     | 説明                                                                                                                                                                                           |\n| :--------- | :----- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |\n| title      | 文字列 | 要件のタイトルを記述してください。これは与えられた要件の 1 行目に記述されています。                                                                                                            |\n| result     | 文字列 | 情報システムが与えられた要件を満たしているか否かを「完全に満たしている」「部分的に満たしている」「満たしていない」「満たす必要がない」「判断できない」「その他」のいずれかで回答してください。 |\n| reason     | 文字列 | result フィールドの回答の理由を日本語で簡潔に記述してください。                                                                                                                                |\n| locations  | 文字列 | result フィールドの回答の根拠となる記述が設計書のどこに書かれているか記述してください。                                                                                                        |\n\n以下は期待される結果の例です。\n\n```json\n{\n  \"title\": \"REL10-BP01 Deploy the workload to multiple locations\",\n  \"result\": \"部分的に満たしている\",\n  \"reason\": \"サブシステム A は High Availability な構成になっていますが、サブシステム B には単一障害点が存在しているため、 Availability Zone がダウンするとシステム全体が停止する恐れがあります。\",\n  \"locations\": \"\"\n}\n```\n\n</output>\n",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prompt, err := wa.Prompt(tt.requirement)
			require.NoError(t, err)
			require.Equal(t, tt.prompt, prompt)
		})
	}
}
