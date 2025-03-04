import React from "react";
import { List } from "antd";
import StockCard from "./StockCard";

const SubscribedStocks = ({ subscribedStocks, handleUnsubscribe }) => {
  return (
    <>
      <h2>Subscribed Stocks</h2>
      <List
        grid={{ gutter: 16, column: 3 }}
        dataSource={subscribedStocks}
        renderItem={(stock) => (
          <List.Item>
            <StockCard stock={stock} handleUnsubscribe={handleUnsubscribe} />
          </List.Item>
        )}
      />
    </>
  );
};

export default SubscribedStocks;