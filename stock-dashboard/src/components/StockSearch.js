import React, { useState } from "react";
import { Input, Button, Row, Col } from "antd";

const StockSearch = ({handleSearch }) => {
  const [manualCode, setManualCode] = useState(""); // State to handle manual stock input

  // Handle input change for manual code
  const handleInputChange = (e) => {
    const value = e.target.value;
    console.log(`Manual code: ${value}`);
    setManualCode(value);
  };

  // Handle the submit button click
  const handleSubmit = () => {
    const code = manualCode; // Use manual code if present, otherwise use selected stockCode
    console.log(`Search for stock code: ${code}`);
    handleSearch(code);
  };

  return (
    <Row justify="center" style={{ marginBottom: "20px" }}>
      <Col>
        {/* Input field with datalist for predefined options */}
        <Input
          list="stockList"
          value={manualCode}
          onChange={handleInputChange}
          placeholder="Enter or select stock code"
          style={{ width: "200px", marginRight: "10px" }}
        />
        
        {/* The datalist element with predefined stock codes */}
        <datalist id="stockList">
          <option value="AAPL">Apple (AAPL)</option>
          <option value="GOOG">Google (GOOG)</option>
          <option value="AMZN">Amazon (AMZN)</option>
          <option value="TSLA">Tesla (TSLA)</option>
        </datalist>

        {/* Button to trigger the search */}
        <Button onClick={handleSubmit}>Search</Button>
      </Col>
    </Row>
  );
};

export default StockSearch;