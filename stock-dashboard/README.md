# Stock Dashboard

The **Stock Dashboard** is a real-time stock market monitoring web application built with **React**. It provides users with stock query and live stock price updates.


## Features
- **Real-time stock price updates** via WebSockets
- **Stock search functionality** for quick access to market data
- **Fast and responsive UI** built with React and Ant Design


## Run

### Run the Dashboard Locally

```sh
    npm start
```

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

### Run in docker
```sh
    docker build -t stock-dashboard .
    docker run -p 3000:3000 stock-dashboard
```