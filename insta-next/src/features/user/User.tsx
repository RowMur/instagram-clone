import { cookies } from "next/headers";
import { graphql } from "../../../gql";
import request from "graphql-request";

const currentUserDocument = graphql(`
  query CurrentUser {
    currentUser {
      user {
        name
      }
    }
  }
`);

const User = async () => {
  const cookieStore = cookies();
  const currentUser = cookieStore.get("key");

  const user = await request(
    "http://localhost:8080/query",
    currentUserDocument,
    undefined,
    {
      Authorization: `ApiKey ${currentUser?.value}`,
    }
  );

  return <div>User: {user.currentUser.user.name}</div>;
};

export default User;
