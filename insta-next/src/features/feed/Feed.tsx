"use client";

import { useQuery } from "@tanstack/react-query";
import PostForm from "./PostForm";
import { fetchFeed } from "./lib";
import Post from "./Post";

type Props = {
  currentUser?: string;
};

const Feed = (props: Props) => {
  const { data } = useQuery({
    queryKey: ["feed"],
    queryFn: fetchFeed,
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
