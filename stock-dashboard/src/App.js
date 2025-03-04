import React, { useState, useEffect } from "react";
import { useQuery } from "@apollo/client";
import { GET_STOCKS } from "./graphql/queries";
import Header from "./components/Header";
import StockSearch from "./components/StockSearch";
import StockList from "./components/StockList";
import SubscribedStocks from "./components/SubscribedStocks";

const App = () => {
  const [stockCode, setStockCode] = useState("");
  const [subscribedStocks, setSubscribedStocks] = useState([]);

  const { data, loading, error } = useQuery(GET_STOCKS, {
    variables: { stockCode },
    skip: stockCode === "",
  });

  useEffect(() => {
    console.log("Updated subscribedStocks:", subscribedStocks);
  }, [subscribedStocks]);

  const handleSearch = (symbol) => {
    setStockCode(symbol);
  };

  const handleSubscribe = (symbol) => {
    setSubscribedStocks((prev) => {
      if (!prev.some((s) => s.symbol === symbol)) {
        return [...prev, { symbol }];
      }
      return prev;
    });
  };

  const handleUnsubscribe = (symbol) => {
    setSubscribedStocks((prev) => prev.filter((s) => s.symbol !== symbol));
  };

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error loading stocks: {error.message}</p>;

  return (
    <div style={{ padding: "20px" }}>
      <Header />
      <StockSearch stockCode={stockCode} setStockCode={setStockCode} handleSearch={handleSearch} />

      <h2>Stock Details</h2>
      <StockList
        stock={data?.stock ? [data.stock] : []}
        handleSubscribe={handleSubscribe}
        subscribedStocks={subscribedStocks}
        handleUnsubscribe={handleUnsubscribe}
      />

      <SubscribedStocks subscribedStocks={subscribedStocks} handleUnsubscribe={handleUnsubscribe} />
    </div>
  );
};

export default App;