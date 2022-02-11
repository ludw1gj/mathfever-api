# MathFever API

A golang gorilla mux server API which provides endpoints to find mathematical proof and answers to common calculation 
problems. Serves as a backend to the React frontend implementation: 
[mathfever-react](https://gitlab.com/ludw1gj/mathfever-react).

## Run
```
docker build -t mathfever-api-image .
docker run --name mathfever-api-container -p 8000:8000 -d mathfever-api-image
```

## AWS Lamba

The AWS Lambda / Serverless Framework version lives on branch: [aws-serverless](https://gitlab.com/ludw1gj/mathfever-api/-/tree/aws-serverless).
