import React from "react";
import axios from "axios";
import Sidebar from "../Sidebar";
import ProjectCard from "../projectComponents/ProjectCard";
import { useCookies } from "react-cookie";

const apiPath = "http://localhost:8080/api/v1";

function Dashboard() {
    const [data, setData] = React.useState([]);
    const [loading, setLoading] = React.useState(true);
    const [userRole, setUserRole] = React.useState("");
    const [cookie, setCookie] = useCookies(["token", "role"]);

    const getUserRole = async () => {
        try {
            var config = {
                headers: {
                    "Content-Type": "application/json",
                    token: cookie.token,
                },
            };
            await axios.get(apiPath + "/user/role", config).then((res) => {
                setUserRole(res.data.role);
            });
        } catch (err) {
            // do nothing
        }
    };
    

    const toggleLoading = () => {
        setLoading(!loading);
    };
    const dashboard = async () => {
        try {
            var config = {
                headers: {
                    "Content-Type": "application/json",
                    token: cookie.token,
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
        if (cookie.token) {
            getUserRole();
            dashboard();
        } else {
            window.location.href = "/";
        }
    }, []);
    return (
        <>
            <Sidebar />
            <div className="flex flex-col bg-blue-500 pl-60 h-72 shadow-xl">
                <div className="relative flex flex-col items-center">
                    <div className="flex flex-row h-full mt-12 pl-36 self-start gap-10">
                        <h1 className="text-4xl font-sans-new text-white">Dashboard</h1>
                        {loading && (
                        <div class="">
                            <div
                                class="spinner-border animate-spin text-white inline-block w-10 h-10 border-3 rounded-full self-start"
                                role="status"
                            >
                                <span class="visually-hidden">Loading...</span>
                            </div>
                        </div>
                    )}
                    </div>
                    <ProjectCard dashData={data} role={userRole} />
                </div>
            </div>
        </>
    );
}

export default Dashboard;
