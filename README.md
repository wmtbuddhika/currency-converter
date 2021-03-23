# Currency Converter
Golang based API end point that used to fetch currency exchange details using external API, fixer.io https://fixer.io/documentation

The env value `FIXER_API_KEY` is generated from fixer dashboard and used to authenticate the external endpoint

### Currency Converting Endpoint

- Endpoint: `/api/convert/`

- Method: `POST`

- Body: `{
         	"fromCurrency": "LKR",
         	"amount": 234.0,
         	"toCurrency": "USD"
         }`

- Responses: 

    `200`

        {
            "amount": 1.20343267683,
            "currency": "USD"
         }
     
    `500`
    
        {
            "status":false,
            "message":"Error on passing request body content"
        }

        {
            "status":false,
            "message":"Error on currency conversion"
        }  
        
    `400`
    
        {
            "status":false,
            "message":"Currency exchange request failed due to validations: [validations]"
        }
        
### Setting up the Development Environment

- Download go: https://golang.org/dl/
- Pull the source code: `git pull https://github.com/wmtbuddhika/currency-converter.git`
- Go to: currency-converter folder
- Set: `.env params`
- Execute: `go build`
- Execute: `go run currency-converter `