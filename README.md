# Binance API

- https://testnet.binance.vision
- https://academy.binance.com/en/articles/binance-api-series-pt-1-spot-trading-with-postman
- https://binance-docs.github.io/apidocs/spot/en/#websocket-market-streams

.env.local

```
SQL_DRIVER=sqlite3
SQL_URL=sqlite3.sqlite

ORDER_SYMBOL=BTCUSDT
ORDER_QTY=2
ORDER_PRICE=40899.59

BINANCE_WS_URL=wss://testnet.binance.vision/ws


```

- how did you approach the task

  - First I did I research for binance service. ( took me quite a bit of time)
  - Created some basic WS client to see how it is working
  - Modeling the data structs
  - basic structure of the code

- where did you have difficulties

  - The WS recive conn.ReadMessage() pasrsing the different object some for respone of request other for data ( Polymorphism in golang is not my strength)

- what part did you like the most

- I have never used WS in such way ( the service to act as a client ), ussually I have been using kafka

- how much time you spent on it in total

  - over all total spend: time 8-10 hours

- describe the next steps that you would work on when you had more time

  - more flexible orders ( price min/max , timeout for one order)

  - handling the actual order execution

  - depends on the use case, run these as a AWS Lambda or GCP Functinos, or create a scalable api service
