"use client";

import { useMutation, useQueryClient } from "@tanstack/react-query";
import { graphql } from "../../../gql";
import request from "graphql-request";
import { useForm } from "@tanstack/react-form";

const createNewUserDocument = graphql(`
  mutation CreateUser($name: String!) {
    createUser(name: $name) {
      apiKey
    }
  }
`);

const SignUpForm = () => {
  const queryClient = useQueryClient();
  const createUserMutation = useMutation({
    mutationFn: async ({ name }: { name: string }) =>
      request("http://localhost:8080/query", createNewUserDocument, {
        name: name,
      }),
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["users"],
      });
    },
  });

  const form = useForm({
    defaultValues: {
      name: "",
    },
    onSubmit: async ({ value }) => {
      createUserMutation.mutate(value);
    },
  });
  return (
    <>
      Sign Up
      <form
        onSubmit={(e) => {
          e.preventDefault();
          e.stopPropagation();
          void form.handleSubmit();
        }}
      >
        <form.Field
          name="name"
          validators={{
            onChange: ({ value }) =>
              value === "" ? "Name is a required field" : undefined,
          }}
          children={(field) => (
            <>
              Name:
              <input
                type="text"
                className="border-2 border-red-400"
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

export default SignUpForm;
