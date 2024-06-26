"use client";

import { useMutation } from "@tanstack/react-query";
import { graphql } from "../../../gql";
import request from "graphql-request";
import { useForm } from "@tanstack/react-form";
import { useRouter, useSearchParams } from "next/navigation";

const createNewUserDocument = graphql(`
  mutation CreateUser($name: String!) {
    createUser(name: $name) {
      apiKey
    }
  }
`);

const SignUpForm = () => {
  const router = useRouter();
  const params = useSearchParams();
  const createUserMutation = useMutation({
    mutationFn: async ({ name }: { name: string }) =>
      request("http://localhost:8080/query", createNewUserDocument, {
        name: name,
      }),
    onSuccess: (data) => {
      document.cookie = `key=${data.createUser?.apiKey}`;
      const continueToUrl = params.get("continue-to") ?? window.location.origin;
      const url = new URL(continueToUrl);
      router.push(url.pathname);
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
