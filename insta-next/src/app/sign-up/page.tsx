import SignUpForm from "@/features/user/SignUpForm";
import { Suspense } from "react";

const Page = () => {
  return (
    <>
      <Suspense>
        <SignUpForm />
      </Suspense>
    </>
  );
};

export default Page;
