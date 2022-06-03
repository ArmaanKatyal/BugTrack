import React from "react";
import { useCookies } from "react-cookie";

function Logout() {
  const [cookie, setCookie, removeCookie] = useCookies(["token", "role"]);
  React.useEffect(() => {
    if (document.cookie.split("=")[1]) {
        removeCookie("token");
        removeCookie("role");
        window.location.href = "/";
    } else {
      window.location.href = "/";
    }
  }, []);
}

export default Logout