"use client";

import { useQuery } from "@tanstack/react-query";
import { graphql } from "../../../gql";
import request from "graphql-request";
import { getUserApiKey } from "@/lib";

const currentUserDocument = graphql(`
  query CurrentUser {
    currentUser {
      user {
        name
      }
    }
  }
`);

const CurrentUser = () => {
  const currentUserQuery = useQuery({
    queryKey: ["currentUser"],
    queryFn: async () =>
      request("http://localhost:8080/query", currentUserDocument, undefined, {
        Authorization: `ApiKey ${getUserApiKey(document)}`,
      }),
  });
  return <div>User: {currentUserQuery.data?.currentUser.user.name}</div>;
};

export default CurrentUser;
