"use client";

import { useQuery } from "@tanstack/react-query";
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
  const data = useQuery({
    queryKey: ["users"],
    queryFn: async () => request("http://localhost:8080/query", queryDocument),
  });
  return (
    <>
      Users:
      {data.data?.users.map((user) => (
        <p key={user.id}>
          {user.id}: {user.name}
        </p>
      ))}
    </>
  );
}
