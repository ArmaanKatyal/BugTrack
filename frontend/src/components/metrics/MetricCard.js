import React from "react";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Doughnut } from "react-chartjs-2";

const apiPath = "http://localhost:8080/api/v1";

ChartJS.register(ArcElement, Tooltip, Legend);

function MetricCard(props) {
    
    var Statusdata = {
        labels: ["Open", "In Progress", "Resolved"],
        datasets: [
            {
                label: "Tickets By Status",
                data: [props.metrics.tickets_by_status.open, props.metrics.tickets_by_status.inprogress, props.metrics.tickets_by_status.resolved],
                backgroundColor: [
                    "rgba(54, 162, 235, 0.2)",
                    "rgba(153, 102, 255, 0.2)",
                    "rgba(75, 192, 192, 0.2)",
                ],
                borderColor: [
                    "rgba(54, 162, 235, 1)",
                    "rgba(153, 102, 255, 1)",
                    "rgba(75, 192, 192, 1)",
                ],
                borderWidth: 1,
            },
        ],
    };

    var Prioritydata = {
        labels: ["Low", "Medium", "High"],
        datasets: [
            {
                label: "Tickets By Priority",
                data: [props.metrics.tickets_by_priority.low, props.metrics.tickets_by_priority.medium, props.metrics.tickets_by_priority.high],
                backgroundColor: [
                    "rgba(54, 162, 235, 0.2)",
                    "rgba(255, 206, 86, 0.2)",
                    "rgba(255, 99, 132, 0.2)",
                ],
                borderColor: [
                    "rgba(54, 162, 235, 1)",
                    "rgba(255, 206, 86, 1)",
                    "rgba(255, 99, 132, 1)",
                ],
                borderWidth: 1,
            },
        ],
    };


    return (
        <div className="mt-10 flex flex-row flex-wrap gap-10">
            <div className="h-fit flex flex-col gap-10">
                <div className="flex flex-col items-center bg-white rounded-2xl shadow-xl p-5 hover:transition hover:shadow-2xl">
                    <h1 className="font-medium font-sans-new">Total Projects</h1>
                    <h1 className="font-normal font-sans-new">{props.metrics && props.metrics.total_projects}</h1>
                </div>
                <div className="flex flex-col items-center bg-white rounded-2xl shadow-xl p-5 hover:transition hover:shadow-2xl">
                    <h1 className="font-medium font-sans-new">Total Tickets</h1>
                    <h1 className="font-normal font-sans-new" >{props.metrics && props.metrics.total_tickets}</h1>
                </div>
            </div>
            {props.metrics ? (
                <div className="h-80 w-72 bg-white rounded-2xl shadow-xl p-5 flex flex-col items-center hover:transition hover:shadow-2xl">
                    <Doughnut data={Statusdata} />
                    <h1 className="font-thin font-sans-new py-4">Tickets By Status</h1>
                </div>
            ) : (
                <div></div>
            )}

            {props.metrics ? (
                <div className="h-80 w-72 bg-white rounded-2xl shadow-xl p-5 flex flex-col items-center hover:transition hover:shadow-2xl">
                    <Doughnut data={Prioritydata} />
                    <h1 className="font-thin font-sans-new py-4">Tickets By Priority</h1>
                </div>
            ) : (
                <div></div>
            )}
        </div>
        // <></>
    );
}

export default MetricCard;
