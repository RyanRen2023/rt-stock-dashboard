import { gql } from "@apollo/client";

export const GET_STOCKS = gql`
  query GetStock($stockCode: String!) {
    stock(stockCode: $stockCode) {
      symbol
      price
      change
      percentChange
    }
  }
`;

export const STOCK_SUBSCRIPTION = gql`
  subscription OnStockPriceUpdated($stockCode: String!) {
    stockPriceUpdated(stockCode: $stockCode) {
      symbol
      price
      change
      percentChange
    }
  }
`;