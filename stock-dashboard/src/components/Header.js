import React from "react";
import { Typography, Row, Col } from "antd";

const { Title } = Typography;

const Header = () => {
  return (
    <Row justify="center">
      <Col>
        <Title level={1}>Real-Time Stock Tracking System</Title>
      </Col>
    </Row>
  );
};

export default Header;