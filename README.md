# AI Definer
AI Definer is an API that allows you to get the definition of terms and phrases using the OpenAI chat completion API (model `gpt-3.5-turbo`).

## Example
#### Request
```
GET /definition/en/cake
```
#### Response
```json
{
  "definition": "A sweet baked food made from a mixture of flour, sugar, eggs, and other ingredients, usually with a sweetening agent such as honey or sugar.",
  "lang": "en",
  "pronunciation": "keɪk",
  "term": "cake",
  "type": "noun"
}
```