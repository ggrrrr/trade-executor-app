# Trade Executor

## Introduction

On exchanges, there are two main order types that we can use to place trades. One is a market order, that executes the trade at the current market price. The advantage of a market order is that the trade is executed. The disadvantage is that you can not control the price much other than the time when you send the order to the exchange. Consequently, you might get a bit more or less, depending on the liquidity in the order book. The other order type is a limit order. With a limit order you can specify the price, but if the market moves too fast, you might end up having nothing or not the entire order size executed. So both have their pros and cons.

## Task

To address this situation, we would like you to create a trade execution service that takes an order size and an order price for an asset pair and generates market orders, when the liquidity in the order book allows it.

The goal is:

- have an asset pair, an order size and a price as input
- consume the order book ticker stream to see the current liquidity
- if you have some order size left that can be executed, store the trade info into a sqlite-db
- output a little summary how the order size was splitted at the end

You do not have to execute real trades. You can simply persist the trades you would have executed.

## Example

We would like to sell 25 BNB at a price no lower than 42 USDT.

We get the following data from the order book ticker stream, where we see that the best price someone is willing to pay is 40 USDT. So we have to wait.

```json
{
  "u": 400900217, // order book updateId
  "s": "BNBUSDT", // symbol
  "b": "40.0", // best bid price
  "B": "10.0", // best bid qty
  "a": "41.0", // best ask price
  "A": "10" // best ask qty
}
```

After a couple of minutes the price moves into our zone, and we get the following update where we see that we can sell 5 BNB at a price of 42 USDT.

```json
{
  "u": 400900223, // order book updateId
  "s": "BNBUSDT", // symbol
  "b": "42.0", // best bid price
  "B": "5.0", // best bid qty
  "a": "43.0", // best ask price
  "A": "10" // best ask qty
}
```

So we persist the info that we sell 5 BNB at 42 USDT with the timestamp and order book update id.

After a couple of minutes we see that more people are willing to sell above 42 USDT when we get the following update:

```json
{
  "u": 400900235, // order book updateId
  "s": "BNBUSDT", // symbol
  "b": "42.5", // best bid price
  "B": "30.0", // best bid qty
  "a": "43.0", // best ask price
  "A": "10" // best ask qty
}
```

So we persist another entry into the db, reporting that we can execute the remaining 20 BNB at a price of 42.5 USDT.

Since we have executed the entire order size, we can print the summary and then stop the service.

## Goals

What we are looking for:

- clean code
- quality is better than quantity
- configurability of the service
- unit tests
- put everything into a git repository, so we are able to review your submission
- at the end, please write a little summary about:
  - how did you approach the task
  - where did you have difficulties
  - what part did you like the most
  - how much time you spent on it in total
  - describe the next steps that you would work on when you had more time

If you implemented all the goals and still have capacity feel free to add additional features or improvements.

## Links

- [Description of order types](https://academy.binance.com/en/articles/understanding-the-different-order-types)
- [Binance API](https://binance-docs.github.io/apidocs/spot/en/#general-info)
- [Order Book Ticker streams](https://binance-docs.github.io/apidocs/spot/en/#individual-symbol-book-ticker-streams)
- [Sqlite](https://www.sqlite.org/index.html)
- [Go sqlite driver](https://github.com/mattn/go-sqlite3)
