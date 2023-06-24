export function getErrorMessage(e: any) {
  if (e.response && e.response.data && e.response.data.error) {
    return e.response.data.error;
  } else if (e.message) {
    return e.message;
  } else if (typeof e === "string") {
    return e;
  } else {
    return "Something went wrong";
  }
}
