import { graphql } from "../../../gql";
import request from "graphql-request";
import Post from "./Post";

const feedDocument = graphql(`
  query Feed {
    posts {
      id
      ...Post
    }
  }
`);

const Feed = async () => {
  const feed = await request("http://localhost:8080/query", feedDocument);
  return (
    <section className="my-10">
      {feed.posts.map((post) => (
        <Post key={post.id} post={post} />
      ))}
    </section>
  );
};

export default Feed;
