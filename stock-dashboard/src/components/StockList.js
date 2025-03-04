// src/components/StockList.js
import React from "react";
import { Table, Button } from "antd";



const StockList = ({ stock, handleSubscribe, handleUnsubscribe, subscribedStocks }) => {
  
  console.log("StockList subscribedStocks:", subscribedStocks);
  const columns = [
    {
      title: "Stock",
      dataIndex: "symbol",
      key: "symbol",
    },
    {
      title: "Price ($)",
      dataIndex: "price",
      key: "price",
      render: (price) => (price !== undefined ? `$${price.toFixed(2)}` : "N/A"),
    },
    {
      title: "Change ($)",
      dataIndex: "change",
      key: "change",
      render: (change) => (change !== undefined ? `${change.toFixed(2)}` : "N/A"),
    },
    {
      title: "Change (%)",
      dataIndex: "percentChange",
      key: "percentChange",
      render: (percentChange) => (percentChange !== undefined ? `${percentChange.toFixed(2)}%` : "N/A"),
    },
    {
      title: "Actions",
      key: "actions",
      render: (_, stockItem) => (
        <>
          <Button
            onClick={() => handleSubscribe(stockItem.symbol)}
            type="primary"
            disabled={!!subscribedStocks.find((s) => s.symbol === stockItem.symbol)}
            style={{ marginRight: 8 }}
          >
            Subscribe
          </Button>
          <Button
            onClick={() => handleUnsubscribe(stockItem.symbol)}
            type="default"
            disabled={!subscribedStocks.find((s) => s.symbol === stockItem.symbol)}
          >
            Unsubscribe
          </Button>
        </>
      ),
    },
  ];

  return <Table dataSource={stock} columns={columns} rowKey="symbol" pagination={false} />;
};

export default StockList;