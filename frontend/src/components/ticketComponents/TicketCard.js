import React from "react";
import TicketItem from "./TicketItem";
import axios from "axios";
import Multiselect from "multiselect-react-dropdown";

const apiPath = "http://localhost:8080/api/v1";

function TicketCard(props) {
    const [userRole, setUserRole] = React.useState("");
    const [userData, setUserData] = React.useState([]);
    const [title, setTitle] = React.useState("");
    const [description, setDescription] = React.useState("");
    const [priority, setPriority] = React.useState("");
    const [assigned_to, setAssigned_to] = React.useState([]);
    const [projectId, setProjectId] = React.useState("");
    const [projectName, setProjectName] = React.useState("");
    const [tags, setTags] = React.useState([]);
    const [projectData, setProjectData] = React.useState([]);
    const [objects, setObjects] = React.useState([]);

    const fillTags = (value) => {
        // make an array of strings which are comma separated values
        var tags = value.split(",");
        setTags(tags);
    };

    const createTicket = async () => {
        if (title === "" || description === "" || priority === "" || projectId === "") {
            alert("Please fill all the fields");
        }
        var data = {
            title: title,
            description: description,
            priority: priority,
            assigned_to: assigned_to,
            project_id: projectId,
            project_name: projectName,
            tags: tags,
        };
        var config = {
            headers: {
                "Content-Type": "application/json",
                token: document.cookie.split("=")[1],
            },
        };
        try {
            await axios.post(apiPath + "/ticket/create", data, config).then((res) => {
                if (res.status === 201) {
                    window.location.href = "/tickets";
                } else {
                    alert("Error");
                }
            });
        } catch (error) {
            // do nothing
        }
    };

    const getProjectData = async () => {
        try {
            var config = {
                headers: {
                    "Content-Type": "application/json",
                    token: document.cookie.split("=")[1],
                },
            };
            await axios.get(apiPath + "/project", config).then((res) => {
                setProjectData(res.data);
            });
        } catch (err) {
            // do nothing
        }
    };

    const setProjectInfo = (id) => {
        setProjectId(id);
        setProjectName(projectData.find((project) => project.id === id).title);
    };

    const getUserRole = async () => {
        try {
            var config = {
                headers: {
                    "Content-Type": "application/json",
                    token: document.cookie.split("=")[1],
                },
            };
            await axios.get(apiPath + "/user/role", config).then((res) => {
                setUserRole(res.data.role);
            });
        } catch (err) {
            // do nothing
        }
    };

    const getUsers = async () => {
        try {
            var config = {
                headers: {
                    "Content-Type": "application/json",
                    token: document.cookie.split("=")[1],
                },
            };

            await axios.get(`${apiPath}/user?role=developer`, config).then((res) => {
                setUserData(res.data);
                fillObjects(res.data);
            });
        } catch (err) {
            // do nothing
        }
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

    const fillObjects = (data) => {
        for (var i = 0; i < data.length; i++) {
            objects.push({
                username: data[i].username,
            });
        }
        // console.log(userData);
    };

    React.useEffect(() => {
        if (document.cookie.split("=")[1]) {
            getUserRole();
            getUsers();
            getProjectData();
        }
    }, []);

    return (
        <div className="absolute w-11/12 flex flex-col bg-white h-auto top-40 rounded-lg shadow-xl font-sans-new">
            <div className="flex flex-row justify-between bg-white rounded-tr-lg rounded-tl-lg shadow-lg">
                <div className="flex flex-col ml-4 p-2">
                    {/* <h2 className="text-xl font-sans-new hidden">Tickets</h2> */}
                    {userRole === "developer" && (
                        <h2 className="text-xl font-sans-new">Assigned Tickets</h2>
                    )}
                    {userRole === "admin" && (
                        <h2 className="text-xl font-sans-new">All Tickets</h2>
                    )}
                    {userRole === "project-manager" && (
                        <h2 className="text-xl font-sans-new">Project Tickets</h2>
                    )}
                </div>
                <div className="flex flex-row mr-10 p-2">
                    <button
                        // onClick={() => toggleCreate()}
                        data-bs-toggle="modal"
                        data-bs-target="#staticBackdrop"
                        type="button"
                        className=""
                    >
                        <span className="text-blue-500 font-semibold text-lg hover:text-blue-700 hover:transition">
                            + New Ticket
                        </span>
                    </button>
                    {/* {showCreate && ( */}
                    <div
                        className="modal fade fixed top-0 left-0 hidden w-full h-full outline-none overflow-x-hidden overflow-y-auto font-sans-new"
                        id="staticBackdrop"
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
                                        New Ticket
                                    </h5>
                                    <button
                                        type="button"
                                        className="btn-close box-content w-4 h-4 p-1 text-black border-none rounded-none opacity-50 focus:shadow-none focus:outline-none focus:opacity-100 hover:text-black hover:opacity-75 hover:no-underline"
                                        data-bs-dismiss="modal"
                                        aria-label="Close"
                                        onClick={clearFieldsOnClose}
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
                                            onChange={(e) =>
                                                setDescription(e.target.value)
                                            }
                                        ></textarea>
                                    </div>
                                    <div className="form-group mb-6">
                                        <textarea
                                            className=" form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                            id="exampleFormControlTextarea13"
                                            rows="1"
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
                                                onChange={(e) =>
                                                    setProjectInfo(e.target.value)
                                                }
                                            >
                                                <option defaultValue>
                                                    Select Project
                                                </option>
                                                {projectData.map((project) => (
                                                    <option
                                                        key={project.id}
                                                        value={project.id}
                                                    >
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
                                                onChange={(e) =>
                                                    setPriority(e.target.value)
                                                }
                                            >
                                                <option defaultValue>
                                                    Select Priority
                                                </option>
                                                <option value="low">Low</option>
                                                <option value="medium">Medium</option>
                                                <option value="high">High</option>
                                            </select>
                                        </div>
                                    </div>
                                    <div>
                                        <div className="mb-3 xl:w-96">
                                            <Multiselect
                                                options={objects}
                                                displayValue="username"
                                                showCheckbox={true}
                                                placeholder="Select Assigned"
                                                onSelect={(value) => onSelect(value)}
                                                onRemove={(value) => onSelect(value)}
                                            />
                                        </div>
                                    </div>
                                </div>
                                <div className="modal-footer flex flex-shrink-0 flex-wrap items-center justify-end p-4 border-t border-gray-200 rounded-b-md">
                                    <button
                                        type="button"
                                        className="inline-block px-6 py-2.5 bg-purple-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-purple-700 hover:shadow-lg focus:bg-purple-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-purple-800 active:shadow-lg transition duration-150 ease-in-out"
                                        data-bs-dismiss="modal"
                                        // onClick={clearFieldsOnClose}
                                    >
                                        Close
                                    </button>
                                    <button
                                        type="button"
                                        onClick={createTicket}
                                        className="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out ml-1"
                                    >
                                        Create
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                    {/* )} */}
                </div>
            </div>

            <div className="flex flex-col pb-5">
                <div className="overflow-x-auto sm:-mx-6 lg:-mx-8">
                    <div className="py-0 inline-block min-w-full sm:px-6 lg:px-8">
                        <div className="overflow-hidden">
                            <table className="min-w-full">
                                <thead className="border-b bg-gray-100 shadow-lg">
                                    <tr>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Ticket Title
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Description
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Project
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Author
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Priority
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Tags
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Status
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Action
                                        </th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {props.data &&
                                        props.data.map((itemvar, key) => (
                                            <TicketItem
                                                key={key}
                                                item={itemvar}
                                                ticketId={itemvar.id}
                                                role={userRole}
                                                users={userData}
                                            />
                                        ))}
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default TicketCard;
