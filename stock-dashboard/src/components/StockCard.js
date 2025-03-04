import React, { useEffect, useState } from "react";
import { useSubscription } from "@apollo/client";
import { STOCK_SUBSCRIPTION } from "../graphql/queries";
import { Card, Button } from "antd";

const StockCard = ({ stock, handleUnsubscribe }) => {
  const [stockData, setStockData] = useState({});

  // Subscribe to real-time data of the current stock
  const { data: subData } = useSubscription(STOCK_SUBSCRIPTION, {
    variables: { stockCode: stock.symbol },
  });

  useEffect(() => {
    if (subData?.stockPriceUpdated) {
      setStockData(subData.stockPriceUpdated);
    }
  }, [subData]);

  return (
    <Card title={stockData?.symbol || "N/A"} style={{ width: 300, margin: "10px" }}>
      {stockData && stockData.price !== undefined ? (
        <>
          <p>Price: ${stockData.price?.toFixed(2)}</p>
          <p>Change: {stockData.change?.toFixed(2)}</p>
          <p>% Change: {stockData.percentChange?.toFixed(2)}%</p>
          <Button onClick={() => handleUnsubscribe(stockData.symbol)} danger>
            Unsubscribe
          </Button>
        </>
      ) : (
        <p>Loading stock data...</p>
      )}
    </Card>
  );
};

export default StockCard;