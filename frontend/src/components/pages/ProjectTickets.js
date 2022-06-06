import React from "react";
import { useParams } from "react-router-dom";
import Sidebar from "../Sidebar";
import axios from "axios";
import { useCookies } from "react-cookie";
import TicketCard from "../ticketComponents/TicketCard";

const apiPath = "http://localhost:8080/api/v1";

function ProjectTickets(props) {
    const params = useParams();
    const [loading, setLoading] = React.useState(true);
    const [projectData, setProjectData] = React.useState([]);
    const [ticketData, setTicketData] = React.useState([]);
    const [cookie, setCookie] = useCookies(["token", "role"]);

    const getProjectData = async () => {
        if (cookie.token) {
            var config = {
                headers: {
                    token: cookie.token,
                },
            };
            try {
                await axios.get(`${apiPath}/project/${params.id}`, config).then((res) => {
                    setProjectData(res.data);
                });
            } catch (err) {
                // do nothing
            }
        } else {
            window.location.href = "/";
        }
    };

    const getTicketData = async () => {
        if (cookie.token) {
            var config = {
                headers: {
                    token: cookie.token,
                },
            };
            try {
                await axios
                    .get(`${apiPath}/ticket/project/${params.id}`, config)
                    .then((res) => {
                        setTicketData(res.data);
                    });
            } catch (err) {
                // do nothing
            }
        } else {
            window.location.href = "/";
        }
    };

    const toggleLoading = () => {
        setLoading(!loading);
    };

    const convertDate = (date) => {
      // convert string to date
      var d = new Date(date);
      // convert date to string
      var dateString = d.toLocaleDateString();
      return dateString;
    }


    React.useEffect(() => {
        getProjectData();
        getTicketData();
        toggleLoading();
    }, []);

    return (
        <>
            <Sidebar />
            <div className="flex flex-col bg-blue-500 pl-60 h-72 shadow-xl">
                <div className="relative flex flex-col items-center">
                    <div className="flex flex-row h-full mt-12 pl-36 self-start gap-10">
                        <h1 className="text-4xl font-sans-new text-white">
                            {projectData.title}
                        </h1>
                        <div className="font-sans-new text-white self-start">
                            <h5 className="font-light leading-tight text-lg mt-1 mb-1 text-white">
                                Created by: {projectData.created_by}
                            </h5>
                            <h5 className="font-light leading-light text-lg mt-1 mb-1 text-white">
                                Project Manager: {projectData.project_manager}
                            </h5>
                            <h5 className="font-light leading-tight text-lg mt-1 mb-1 text-white">
                                {/* {projectData.created_on} */}
                                Created on: {!loading && (convertDate(projectData.created_on))}
                            </h5>
                        </div>
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

                    <TicketCard data={ticketData} />
                </div>
            </div>
        </>
    );
}

export default ProjectTickets;
