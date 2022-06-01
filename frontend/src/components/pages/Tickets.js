import React from "react";
import Sidebar from "../Sidebar";
import axios from "axios";
import TicketCard from "../ticketComponents/TicketCard";
import { useCookies } from "react-cookie";

const apiPath = "http://localhost:8080/api/v1";

function Tickets() {
    const [loading, setLoading] = React.useState(true);
    const [data, setData] = React.useState([]);
    const [cookie, setCookie] = useCookies(["token"]);

    const toggleLoading = () => {
        setLoading(!loading);
    };

    const getTickets = async () => {
        try {
            var config = {
                headers: {
                    "Content-Type": "application/json",
                    token: cookie.token,
                },
            };
            await axios.get(apiPath + "/ticket", config).then((res) => {
                toggleLoading();
                setData(res.data);
            });
        } catch (err) {
            window.location.href = "/";
        }
    };

    React.useEffect(() => {
        if (cookie.token) {
            getTickets();
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
                        <h1 className="text-4xl text-white font-sans-new">Tickets</h1>
                    </div>
                    <TicketCard data={data} />
                </div>
            </div>
        </>
    );
}

export default Tickets;
