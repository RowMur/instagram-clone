"use client";

import { useQuery } from "@tanstack/react-query";
import PostForm from "./PostForm";
import Post from "./Post";
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

type Props = {
  currentUser?: string;
};

const Feed = (props: Props) => {
  const { data } = useQuery({
    queryKey: ["feed"],
    queryFn: () =>
      request("http://localhost:8080/query", feedDocument, undefined),
  });
  return (
    <section className="my-10">
      <PostForm currentUser={props.currentUser} />
      {data?.posts.map((post) => (
        <Post key={post.id} post={post} />
      ))}
    </section>
  );
};

export default Feed;
