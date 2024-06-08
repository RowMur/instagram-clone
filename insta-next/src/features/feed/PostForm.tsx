"use client";

import {
  useMutation,
  QueryClient,
  useQueryClient,
} from "@tanstack/react-query";
import { graphql } from "../../../gql";
import request from "graphql-request";
import { useForm } from "@tanstack/react-form";

const postDocument = graphql(`
  mutation Post($text: String!) {
    createPost(text: $text) {
      __typename
    }
  }
`);

type Props = {
  currentUser?: string;
};

const PostForm = (props: Props) => {
  const queryClient = useQueryClient();
  const postMutation = useMutation({
    mutationFn: async ({ text }: { text: string }) =>
      request(
        "http://localhost:8080/query",
        postDocument,
        {
          text: text,
        },
        {
          Authorization: `ApiKey ${props.currentUser}`,
        }
      ),
    onSuccess: () => {
      form.reset();
      queryClient.invalidateQueries({
        queryKey: ["feed"],
      });
    },
  });

  const form = useForm({
    defaultValues: {
      text: "",
    },
    onSubmit: async ({ value }) => {
      postMutation.mutate(value);
    },
  });
  return (
    <>
      <form
        onSubmit={(e) => {
          e.preventDefault();
          e.stopPropagation();
          void form.handleSubmit();
        }}
      >
        <form.Field
          name="text"
          validators={{
            onSubmit: ({ value }) =>
              value ? undefined : "You have to type something...",
          }}
          children={(field) => (
            <>
              <input
                type="text"
                className="border-2 border-red-400"
                placeholder="Post..."
                name={field.name}
                value={field.state.value}
                onBlur={field.handleBlur}
                onChange={(e) => field.handleChange(e.target.value)}
              />
              {field.state.meta.errors ? (
                <em>{field.state.meta.errors}</em>
              ) : null}
            </>
          )}
        />
      </form>
    </>
  );
};

export default PostForm;
