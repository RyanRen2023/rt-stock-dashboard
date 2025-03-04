import { ApolloClient, InMemoryCache, split, HttpLink } from "@apollo/client";
import { GraphQLWsLink } from "@apollo/client/link/subscriptions";
import { createClient } from "graphql-ws";
import { getMainDefinition } from "@apollo/client/utilities";

const API_URL = process.env.REACT_APP_GRAPHQL_API_URL || "http://localhost:8080/query";
const WS_URL = process.env.REACT_APP_GRAPHQL_WS_URL || "ws://localhost:8080/query";


// HTTP request
const httpLink = new HttpLink({
  uri: API_URL,
});

// WebSocket subscription
const wsLink = new GraphQLWsLink(
  createClient({
    url: WS_URL,
  })
);

// a split link that directs the query to the right place
const splitLink = split(
  ({ query }) => {
    const definition = getMainDefinition(query);
    return definition.kind === "OperationDefinition" && definition.operation === "subscription";
  },
  wsLink,
  httpLink
);

const client = new ApolloClient({
  link: splitLink,
  cache: new InMemoryCache(),
});

export default client;