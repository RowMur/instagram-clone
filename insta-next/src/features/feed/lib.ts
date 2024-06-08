import request from "graphql-request";
import { graphql } from "../../../gql";

const feedDocument = graphql(`
  query Feed {
    posts {
      id
      ...Post
    }
  }
`);

export const fetchFeed = () => {
  return request("http://localhost:8080/query", feedDocument, undefined);
};
