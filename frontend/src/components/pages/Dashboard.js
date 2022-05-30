import React from "react";
import axios from "axios";
import Sidebar from "../Sidebar";
import ProjectCard from "../ProjectCard";

const apiPath = "http://localhost:8080/api/v1";

function Dashboard() {
    const [data, setData] = React.useState([]);
    const [loading, setLoading] = React.useState(true);

    const toggleLoading = () => {
        setLoading(!loading);
    };
    const dashboard = async () => {
        try {
            var config = {
                headers: {
                    "Content-Type": "application/json",
                    token: document.cookie.split("=")[1],
                },
            };
            await axios.get(apiPath + "/project", config).then((res) => {
                toggleLoading();
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
            <div className="flex flex-col bg-blue-500 pl-60 h-72 shadow-xl">
            {loading && (
                        <div class="ml-5 mt-5">
                            <div
                                class="spinner-border animate-spin text-white inline-block w-10 h-10 border-3 rounded-full self-start"
                                role="status"
                            >
                                <span class="visually-hidden">Loading...</span>
                            </div>
                        </div>
                    )}
                <div className="relative flex flex-col items-center">
                    <div className="flex flex-col h-full mt-12 pl-36 self-start">
                        <h1 className="text-3xl font-sans-new text-white">Dashboard</h1>
                    </div>
                    <ProjectCard dashData={data} />
                </div>
            </div>
        </>
    );
}

export default Dashboard;
