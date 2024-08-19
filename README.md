# AI Definer
AI Definer is an API that allows you to get the definition of terms and phrases using the OpenAI chat completion API (model `gpt-3.5-turbo`).

## Usage
To use AI Definer, firstly set your OpenAI API key in a `.env` file (see [`.env.example`](./.env.example)).

Then, run the API using the following command :
```bash
go run .\main\main.go
```
*The API will be available at http://localhost:8080/.*

To compile the API, use the following command :
```bash
go build -o ai-definer.exe .\main\main.go
```

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