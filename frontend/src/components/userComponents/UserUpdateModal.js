import React from "react";
import axios from "axios";
import { useCookies } from "react-cookie";

const apiPath = "http://localhost:8080/api/v1";

function UserUpdateModal(props) {
    const [cookie, setCookie] = useCookies(["token", "role"]);
    const [first_name, setFirstName] = React.useState(props.item.first_name);
    const [last_name, setLastName] = React.useState(props.item.last_name);
    const [email, setEmail] = React.useState(props.item.email);
    const [role, setRole] = React.useState(props.item.role);

    const updateUser = async () => {
        if (cookie.token) {
            if (cookie.role === "admin") {
                try {
                    var config = {
                        headers: {
                            "Content-Type": "application/json",
                            token: cookie.token,
                        },
                    };
                    var data = {
                        first_name: first_name,
                        last_name: last_name,
                        email: email,
                        role: role,
                    };
                    await axios
                        .put(apiPath + "/user/update/" + props.item.username, data, config)
                        .then((res) => {
                            if (res.status === 200) {
                                window.location.href = "/user-management";
                            } else {
                                alert("Error");
                            }
                        });
                } catch (error) {
                    // do nothing
                }
            } else {
                alert("You are not authorized to perform this action");
            }
        } else {
            window.location.href = "/";
        }
    };

    return (
        <div
            className="modal fade fixed top-0 left-0 hidden w-full h-full outline-none overflow-x-hidden overflow-y-auto font-sans-new"
            id={"updatestaticBackdrop" + props.item.username}
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
                            Update User
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
                                placeholder="First Name"
                                value={first_name}
                                onChange={(e) => setFirstName(e.target.value)}
                            />
                        </div>
                        <div className="form-group mb-6">
                            <input
                                type="text"
                                className="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                id="exampleInput7"
                                placeholder="Last Name"
                                value={last_name}
                                onChange={(e) => setLastName(e.target.value)}
                            />
                        </div>
                        <div className="form-group mb-6">
                            <input
                                type="email"
                                className="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                id="exampleInput7"
                                placeholder="Email"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                            />
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
                                    onChange={(e) => setRole(e.target.value)}
                                >
                                    <option defaultValue>Select Role</option>
                                    <option value="submitter">Submitter</option>
                                    <option value="developer">Developer</option>
                                    <option value="project-manager">
                                        Project Manager
                                    </option>
                                    <option value="admin">Admin</option>
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
                            onClick={updateUser}
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

export default UserUpdateModal;
