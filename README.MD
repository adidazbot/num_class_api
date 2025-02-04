# Number Classification API FOR STAGE ONE 

## 🚀 Live API URL
https://railway.com/project/018804f7-6e6c-44e2-8fa9-6a1b3285b019/service/f2dd8c0c-168a-4fbc-9a7d-a3d84062c1bd?environmentId=f10a999d-60ed-4a56-998a-69705445c233&id=b1b9423d-fbe6-4154-845c-a91a8a59d875

## 📌 Endpoint
GET `/api/classify-number?number=371`

### 📌 Example Response
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number"
}

