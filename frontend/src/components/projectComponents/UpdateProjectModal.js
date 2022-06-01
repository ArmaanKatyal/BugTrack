import React from "react";
import Multiselect from "multiselect-react-dropdown";
import axios from "axios";
import { useCookies } from "react-cookie";
const apiPath = "http://localhost:8080/api/v1";

function UpdateProjectModal(props) {
    const [id, setId] = React.useState(props.projectId);
    const [title, setTitle] = React.useState(props.item.title);
    const [description, setDescription] = React.useState(props.item.description);
    const [created_by, setCreated_by] = React.useState(props.item.created_by);
    const [created_on, setCreated_on] = React.useState(props.item.created_on);
    const [assigned_to, setAssigned_to] = React.useState(props.item.assigned_to);
    const [company_code, setCompany_code] = React.useState(props.item.company_code);
    const [project_manager, setProject_manager] = React.useState(
        props.item.project_manager
    );
    const [objects, setObjects] = React.useState([]);
    const [selected, setSelectedObjects] = React.useState([]);
    const [cookie, setCookie] = useCookies(["token"]);

    const updateProject = async () => {
        var data = {
            id: id,
            title: title,
            description: description,
            created_by: created_by,
            created_on: created_on,
            assigned_to: assigned_to,
            company_code: company_code,
            project_manager: project_manager,
        };
        var config = {
            headers: {
                "Content-Type": "application/json",
                token: cookie.token,
            },
        };
            await axios
                .put(apiPath + "/project/update/" + props.projectId, data, config)
                .then((res) => {
                    if (res.status === 200) {
                        window.location.href = "/dashboard";
                    } else {
                        alert("Error");
                    }
                });
    };

    for (var i = 0; i < props.item.assigned_to.length; i++) {
        selected.push({
            username: props.item.assigned_to[i],
        });
    }

    React.useEffect(() => {
        for (var i = 0; i < props.users.length; i++) {
            objects.push({
                username: props.users[i].username,
            });
        }
    }, []);

    const onSelect = (selectedList) => {
        var selectedArray = [];
        for (var i = 0; i < selectedList.length; i++) {
            selectedArray.push(selectedList[i].username);
        }
        setAssigned_to(selectedArray);
    };

    return (
        <div
            className="modal fade fixed top-0 left-0 hidden w-full h-full outline-none overflow-x-hidden overflow-y-auto font-sans-new"
            id={"updatestaticBackdrop" + props.projectId}
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
                            Update Project
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
                                onChange={(e) => setDescription(e.target.value)}
                            ></textarea>
                        </div>
                        {props.role === "admin" && (
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
                                    onChange={(e) => setProject_manager(e.target.value)}
                                >
                                    <option defaultValue>Select Project Manager</option>
                                    {props.users.map((user) => (
                                        <option key={user.username} value={user.username}>
                                            {user.username}
                                        </option>
                                    ))}
                                </select>
                            </div>
                        </div>
                        )}
                        <div>
                            <div className="mb-3 xl:w-96">
                                <Multiselect
                                    options={objects}
                                    displayValue="username"
                                    showCheckbox={true}
                                    onSelect={(value) => onSelect(value)}
                                    onRemove={(value) => onSelect(value)}
                                    selectedValues={selected}
                                />
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
                            onClick={updateProject}
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

export default UpdateProjectModal;
