import React from "react";
import axios from "axios";
import Sidebar from "../Sidebar";

const apiPath = "http://localhost:8080/api/v1";

function Dashboard() {
  const [data, setData] = React.useState([]);
  const dashboard = async () => {
    try {
      var config = {
        headers: {
          "Content-Type": "application/json",
          token: document.cookie.split("=")[1],
        },
      };
      await axios.get(apiPath + "/project", config).then((res) => {
        setData(res.data);
      });
    } catch (err) {
      window.location.href = "/";
    }
  };

  React.useEffect(() => {
    if (document.cookie.split("=")[1]) {
      dashboard();
    } else {
      window.location.href = "/";
    }
  }, []);
  return (
    <>
      <Sidebar />
      <div className="flex flex-col bg-blue-500 h-72 pl-60">
        <div className="flex flex-col h-full mt-12 pl-36">
          <h1 className="text-3xl text-white">Dashboard</h1>
        </div>
      </div>
    </>
  );
}

export default Dashboard;
