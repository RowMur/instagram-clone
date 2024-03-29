"use client";

import { useQuery } from "react-query";
import { graphql } from "../../gql";
import request from "graphql-request";

const queryDocument = graphql(`
  query test {
    users {
      id
      name
    }
  }
`);

export default function TestComponent() {
  const data = useQuery(["users"], async () =>
    request("http://localhost:8080/query", queryDocument)
  );
  return (
    <>
      {data.data?.users.map((user) => (
        <p key={user.id}>
          {user.id}: {user.name}
        </p>
      ))}
    </>
  );
}
