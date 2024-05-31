export const getUserApiKey = (document: Document) => {
  return document.cookie
    .split(";")
    .find((cookie) => cookie.startsWith("key"))
    ?.split("=")[1];
};
