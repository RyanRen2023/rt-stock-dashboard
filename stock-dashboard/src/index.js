import React from 'react';
import ReactDOM from 'react-dom/client';
import { ApolloProvider } from "@apollo/client";
import client from "./graphql/client"; // Import the Apollo client
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <ApolloProvider client={client}> {/* Wrap the App component with the ApolloProvider */}
    <App />
  </ApolloProvider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
