import { FragmentType, getFragmentData, graphql } from "../../../gql";
import moment from "moment";

const postFragment = graphql(`
  fragment Post on Post {
    id
    text
    createdAt
    user {
      name
    }
  }
`);

type Props = {
  post: FragmentType<typeof postFragment>;
};

const Post = (props: Props) => {
  const post = getFragmentData(postFragment, props.post);
  const createdAtDate = new Date(post.createdAt);
  const createAtTimeAgo = moment(createdAtDate).fromNow();

  return (
    <div key={post.id} className="my-4">
      <p className="opacity-60 flex justify-between">
        <span>@{post.user.name}</span>
        <span>{createAtTimeAgo}</span>
      </p>
      <p>{post.text}</p>
    </div>
  );
};

export default Post;
