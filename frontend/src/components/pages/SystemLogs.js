import React from "react";
import LogCard from "../logsComponents/LogCard";
import Sidebar from "../Sidebar";
import { useCookies } from "react-cookie";
import axios from "axios";

function SystemLogs() {
    const [loading, setLoading] = React.useState(true);
    const [data, setData] = React.useState([]);
    const [cookie, setCookie] = useCookies(["token"]);

    const toggleLoading = () => {
        setLoading(!loading);
    };

    React.useEffect(() => {
        if (cookie.token) {
            try {
              axios.get("http://localhost:8080/api/v1/logs", {
                headers: {
                  "Content-Type": "application/json",
                  token: cookie.token,
                },
              }).then((res) => {
                toggleLoading();
                setData(res.data);
              }
              );
            } catch (err) {
              window.location.href = "/";
            }
        } else {
            window.location.href = "/";
        }
    }, []);
    return (
        <>
            <Sidebar />
            <div className="flex flex-col bg-purple-600 h-72 pl-60">
                {loading && (
                    <div className="ml-5 mt-5">
                        <div
                            className="spinner-border animate-spin text-white inline-block w-10 h-10 border-3 rounded-full self-start"
                            role="status"
                        >
                            <span className="visually-hidden">Loading...</span>
                        </div>
                    </div>
                )}
                <div className="relative flex flex-col items-center">
                    <div className="flex flex-col h-full mt-12 ml-36 self-start">
                        <h1 className="text-4xl text-white font-sans-new">System Logs</h1>
                    </div>
                    <LogCard data={data}/>
                </div>
            </div>
        </>
    );
}

export default SystemLogs;
