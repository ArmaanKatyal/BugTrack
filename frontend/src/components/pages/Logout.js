import React from "react";

function Logout() {
  React.useEffect(() => {
    if (document.cookie.split("=")[1]) {
        document.cookie = "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
        window.location.href = "/";
    } else {
      window.location.href = "/";
    }
  }, []);
}

export default Logout