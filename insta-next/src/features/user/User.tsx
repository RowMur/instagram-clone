"use client";

import { getUserApiKey } from "@/lib";
import CurrentUser from "./CurrentUser";
import SignUpForm from "./SignUpForm";
import { useQuery } from "@tanstack/react-query";

const User = () => {
  const currentUserApiKey = useQuery({
    queryKey: ["currentUserApiKey"],
    queryFn: () => getUserApiKey(document),
  });
  return <>{currentUserApiKey.data ? <CurrentUser /> : <SignUpForm />}</>;
};

export default User;
