import React from "react";
import axios from "axios";
import ProjectItem from "./ProjectItem";

const apiPath = "http://localhost:8080/api/v1";

function ProjectCard(props) {
    const [title, setTitle] = React.useState("");
    const [description, setDescription] = React.useState("");
    const [userRole, setUserRole] = React.useState("");
    const [userData, setUserData] = React.useState([]);

    const handleSubmit = async () => {
        try {
            var config = {
                headers: {
                    "Content-Type": "application/json",
                    token: document.cookie.split("=")[1],
                },
            };
            var data = {
                title: title,
                description: description,
                created_by: "",
                assigned_to: "",
                company_code: "",
            };
            await axios.post(apiPath + "/project/create", data, config).then((res) => {
                if (res.status === 201) {
                    setTitle("");
                    setDescription("");
                    window.location.href = "/dashboard";
                } else {
                    alert("Error");
                }
            });
        } catch (err) {
            alert("Error");
            setTitle("");
            setDescription("");
            window.location.href = "/dashboard";
        }
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

            await axios.get(`${apiPath}/user?role=`, config).then((res) => {
                setUserData(res.data);
            });
        } catch (err) {
            // do nothing
        }
    };

    React.useEffect(() => {
        if (document.cookie.split("=")[1]) {
            getUserRole();
            getUsers();
        }
    }, []);
    return (
        <div className="absolute w-11/12 flex flex-col bg-white h-auto top-40 rounded-lg shadow-xl font-sans-new">
            <div className="flex flex-row justify-between bg-white rounded-tr-lg rounded-tl-lg shadow-lg">
                <div className="flex flex-col ml-10 p-2">
                    <h2 className="text-xl">Projects</h2>
                </div>
                <div className="flex flex-row mr-10 p-2">
                    <button
                        // onClick={() => toggleCreate()}
                        data-bs-toggle="modal"
                        data-bs-target="#staticBackdrop"
                        type="button"
                        className="bg-purple-600 text-sm text-white px-3 rounded-lg hover:bg-black hover:text-white hover:transition"
                    >
                        New Project
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
                                        New Project
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
                                            placeholder="Message"
                                            value={description}
                                            onChange={(e) =>
                                                setDescription(e.target.value)
                                            }
                                        ></textarea>
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
                                        onClick={handleSubmit}
                                        className="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out ml-1"
                                    >
                                        Create Project
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
                    <div className="py-2 inline-block min-w-full sm:px-6 lg:px-8">
                        <div className="overflow-hidden">
                            <table className="min-w-full">
                                <thead className="border-b bg-gray-100">
                                    <tr>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Name
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
                                            Project-Manager
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
                                    {props.dashData &&
                                        props.dashData.map((itemvar, key) => (
                                            <ProjectItem
                                                key={key}
                                                item={itemvar}
                                                projectId={itemvar.id}
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

export default ProjectCard;
