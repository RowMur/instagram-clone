import { NextRequest, NextResponse } from "next/server";

export const middleware = (request: NextRequest) => {
  const currentUserCookie = request.cookies.get("key");

  if (currentUserCookie && request.url.includes("/sign-up")) {
    return NextResponse.redirect(new URL("/", request.url));
  }

  if (!currentUserCookie && !request.url.includes("/sign-up")) {
    return NextResponse.redirect(
      new URL(`/sign-up?continue-to=${request.url}`, request.url)
    );
  }
};

export const config = {
  matcher: ["/((?!api|_next/static|_next/image|favicon.ico).*)"],
};
