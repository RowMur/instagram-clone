import Feed from "@/features/feed/Feed";
import User from "@/features/user/User";
import { cookies } from "next/headers";

export default function Home() {
  const cookieStore = cookies();
  const currentUser = cookieStore.get("key");
  return (
    <>
      <User />
      <Feed currentUser={currentUser?.value} />
    </>
  );
}
