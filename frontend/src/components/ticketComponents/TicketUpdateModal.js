import React from "react";
import Multiselect from "multiselect-react-dropdown";
import axios from "axios";
import { useCookies } from "react-cookie";

const apiPath = "http://localhost:8080/api/v1";

function TicketUpdateModal(props) {
    const [id, setId] = React.useState(props.ticketId);
    const [title, setTitle] = React.useState(props.item.title);
    const [description, setDescription] = React.useState(props.item.description);
    const [priority, setPriority] = React.useState(props.item.priority);
    const [assigned_to, setAssigned_to] = React.useState(props.item.assigned_to);
    const [projectId, setProjectId] = React.useState(props.item.project_id);
    const [projectName, setProjectName] = React.useState(props.item.project_name);
    const [tags, setTags] = React.useState(props.item.tags);
    const [status, setStatus] = React.useState(props.item.status);
    const [selected, setSelected] = React.useState([]);
    const [cookie, setCookie] = useCookies(["token"]);

    for (var i = 0; i < props.item.assigned_to.length; i++) {
        selected.push({
            username: props.item.assigned_to[i],
        });
    }

    const fillTags = (value) => {
        // make an array of strings which are comma separated values
        var tags = value.split(",");
        setTags(tags);
    };

    const setProjectInfo = (id) => {
        setProjectId(id);
        setProjectName(props.projects.find((project) => project.id === id).title);
    };

    const clearFieldsOnClose = () => {
        setTitle("");
        setDescription("");
        setPriority("");
        setAssigned_to([]);
    };

    const onSelect = (selectedList) => {
        var selectedArray = [];
        for (var i = 0; i < selectedList.length; i++) {
            selectedArray.push(selectedList[i].username);
        }
        setAssigned_to(selectedArray);
    };

    const updateTicket = async () => {
        if (cookie.token) {
            var data = {
                title: title,
                description: description,
                priority: priority,
                assigned_to: assigned_to,
                project_id: projectId,
                project_name: projectName,
                tags: tags,
                status: status,
            };
            var config = {
                headers: {
                    "Content-Type": "application/json",
                    token: cookie.token,
                },
            };
            try {
                await axios
                    .put(apiPath + "/ticket/update/" + props.ticketId, data, config)
                    .then((res) => {
                        if (res.status === 200) {
                            window.location.href = "/tickets";
                        } else {
                            alert("Error");
                        }
                    });
            } catch (error) {
                alert("Error");
            }
        } else {
            window.location.href = "/";
        }
    };

    return (
        <div
            className="modal fade fixed top-0 left-0 hidden w-full h-full outline-none overflow-x-hidden overflow-y-auto font-sans-new"
            id={"updatestaticBackdrop" + props.ticketId}
            data-bs-backdrop="static"
            data-bs-keyboard="false"
            tabIndex="-1"
            aria-labelledby="staticBackdropLabel"
            aria-hidden="true"
        >
            <div className="modal-dialog relative w-auto pointer-events-none">
                <div className="modal-content border-none shadow-lg relative flex flex-col w-full pointer-events-auto bg-white bg-clip-padding rounded-md outline-none text-current">
                    <div className="modal-header flex flex-shrink-0 items-center justify-between p-4 border-b border-gray-200 rounded-t-md">
                        <h5
                            className="text-xl font-medium leading-normal text-gray-800"
                            id="exampleModalLabel"
                        >
                            Update Ticket
                        </h5>
                        <button
                            type="button"
                            className="btn-close box-content w-4 h-4 p-1 text-black border-none rounded-none opacity-50 focus:shadow-none focus:outline-none focus:opacity-100 hover:text-black hover:opacity-75 hover:no-underline"
                            data-bs-dismiss="modal"
                            aria-label="Close"
                        ></button>
                    </div>
                    <div className="modal-body relative p-4">
                        <div className="form-group mb-6">
                            <input
                                type="text"
                                className="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                id="exampleInput7"
                                placeholder="Title"
                                value={title}
                                onChange={(e) => setTitle(e.target.value)}
                            />
                        </div>
                        <div className="form-group mb-6">
                            <textarea
                                className=" form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                id="exampleFormControlTextarea13"
                                rows="3"
                                placeholder="Description"
                                value={description}
                                onChange={(e) => setDescription(e.target.value)}
                            ></textarea>
                        </div>
                        <div className="form-group mb-6">
                            <textarea
                                className=" form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                id="exampleFormControlTextarea13"
                                rows="1"
                                value={tags}
                                placeholder="Tags ( , seperated)"
                                onChange={(e) => fillTags(e.target.value)}
                            ></textarea>
                        </div>
                        <div className="">
                            <div className="mb-3 xl:w-96">
                                <select
                                    className="form-select appearance-none
                                            block
                                            w-full
                                            px-3
                                            py-1.5
                                            text-base
                                            font-normal
                                            text-gray-700
                                            bg-white bg-clip-padding bg-no-repeat
                                            border border-solid border-gray-300
                                            rounded
                                            transition
                                            ease-in-out
                                            m-0
                                            focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                    aria-label="Default select example"
                                    onChange={(e) => setProjectInfo(e.target.value)}
                                >
                                    <option defaultValue>Select Project</option>
                                    {props.projects.map((project) => (
                                        <option key={project.id} value={project.id}>
                                            {project.title}
                                        </option>
                                    ))}
                                </select>
                            </div>
                        </div>
                        <div className="">
                            <div className="mb-3 xl:w-96">
                                <select
                                    className="form-select appearance-none
                                            block
                                            w-full
                                            px-3
                                            py-1.5
                                            text-base
                                            font-normal
                                            text-gray-700
                                            bg-white bg-clip-padding bg-no-repeat
                                            border border-solid border-gray-300
                                            rounded
                                            transition
                                            ease-in-out
                                            m-0
                                            focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                    aria-label="Default select example"
                                    onChange={(e) => setPriority(e.target.value)}
                                >
                                    <option defaultValue>Select Priority</option>
                                    <option value="low">Low</option>
                                    <option value="medium">Medium</option>
                                    <option value="high">High</option>
                                </select>
                            </div>
                        </div>
                        <div>
                            <div className="mb-3 xl:w-96">
                                <Multiselect
                                    options={props.objects}
                                    selectedValues={selected}
                                    displayValue="username"
                                    showCheckbox={true}
                                    placeholder="Select Assigned"
                                    onSelect={(value) => onSelect(value)}
                                    onRemove={(value) => onSelect(value)}
                                />
                            </div>
                        </div>
                        <div className="">
                            <div className="mb-3 xl:w-96">
                                <select
                                    className="form-select appearance-none
                                            block
                                            w-full
                                            px-3
                                            py-1.5
                                            text-base
                                            font-normal
                                            text-gray-700
                                            bg-white bg-clip-padding bg-no-repeat
                                            border border-solid border-gray-300
                                            rounded
                                            transition
                                            ease-in-out
                                            m-0
                                            focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                    aria-label="Default select example"
                                    onChange={(e) => setStatus(e.target.value)}
                                >
                                    <option defaultValue>Select Status</option>
                                    <option value="open">Open</option>
                                    <option value="in-progress">In-Progress</option>
                                    <option value="resolved">Resolved</option>
                                </select>
                            </div>
                        </div>
                    </div>
                    <div className="modal-footer flex flex-shrink-0 flex-wrap items-center justify-end p-4 border-t border-gray-200 rounded-b-md">
                        <button
                            type="button"
                            className="inline-block px-6 py-2.5 bg-purple-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-purple-700 hover:shadow-lg focus:bg-purple-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-purple-800 active:shadow-lg transition duration-150 ease-in-out"
                            data-bs-dismiss="modal"
                        >
                            Close
                        </button>
                        <button
                            type="button"
                            onClick={updateTicket}
                            className="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out ml-1"
                        >
                            Update
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default TicketUpdateModal;
