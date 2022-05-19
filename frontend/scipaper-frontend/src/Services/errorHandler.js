import { toast } from "react-toastify";

const errorHandler = (errorResponse) => {
  if (!errorResponse) {
    toast.error("The server is unavailable.");
    return;
  }
  switch (errorResponse.status) {
    case 400:
      toast.error("Bad request - " + errorResponse.data.Message);
      break;

    case 401:
      toast.error("Unauthorized - Please log in to access this resource.");
      break;

    case 403:
      toast.error(
        "Forbidden - The client did not have permission to access the requested resource."
      );
      break;

    case 404:
      toast.error("Not found - " + errorResponse.data.Message);
      break;

    case 500:
      toast.error("Internal server error.");
      break;

    case 503:
      toast.error("The server was unavailable.");
      break;

    default:
      toast.error("Something wrong");
      break;
  }
};

export default errorHandler;