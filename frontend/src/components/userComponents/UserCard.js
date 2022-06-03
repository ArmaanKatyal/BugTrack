import React from "react";
import UserItem from "./UserItem";

function UserCard(props) {
    return (
        <div className="absolute w-11/12 flex flex-col bg-white h-auto top-40 rounded-lg shadow-xl font-sans-new">
            <div className="flex flex-row justify-between bg-white rounded-tr-lg rounded-tl-lg shadow-lg">
                <div className="flex flex-col ml-4 p-2">
                    <h2 className="text-xl font-sans-new ">All Users</h2>
                    {/* {props.role === "developer" && (
                            <h2 className="text-xl font-sans-new">Assigned Projects</h2>
                    )}
                    {props.role === "admin" && (
                            <h2 className="text-xl font-sans-new">All Projects</h2>
                    )}
                    {props.role === "project-manager" && (
                            <h2 className="text-xl font-sans-new">Projects</h2>
                    )} */}
                </div>
                <div className="flex flex-row p-2 items-end gap-5 mr-7">
                    {props.role === "admin" && (
                        <>
                            <button
                                // onClick={() => toggleCreate()}
                                data-bs-toggle="modal"
                                data-bs-target="#createStaticBackdrop"
                                type="button"
                                className=""
                            >
                                <span className="text-blue-500 font-semibold text-lg hover:text-blue-700 hover:transition">
                                    + New User
                                </span>
                            </button>
                            <button
                                // onClick={() => toggleCreate()}
                                data-bs-toggle="modal"
                                data-bs-target="#lockedUsersStaticBackdrop"
                                type="button"
                                className=""
                            >
                                <span className="text-blue-500 font-semibold text-lg hover:text-blue-700 hover:transition">
                                    Locked User(s)
                                </span>
                            </button>
                        </>
                    )}
                    {/* {showCreate && ( */}
                    <div
                        className="modal fade fixed top-0 left-0 hidden w-full h-full outline-none overflow-x-hidden overflow-y-auto font-sans-new"
                        id="createStaticBackdrop"
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
                                        New User
                                    </h5>
                                    <button
                                        type="button"
                                        className="btn-close box-content w-4 h-4 p-1 text-black border-none rounded-none opacity-50 focus:shadow-none focus:outline-none focus:opacity-100 hover:text-black hover:opacity-75 hover:no-underline"
                                        data-bs-dismiss="modal"
                                        aria-label="Close"
                                        // onClick={clearFieldsOnClose}
                                    ></button>
                                </div>
                                <div className="modal-body relative p-4">
                                    <div className="form-group mb-6">
                                        <input
                                            type="text"
                                            className="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                            id="exampleInput7"
                                            placeholder="First Name"
                                            // value={title}
                                            // onChange={(e) => setTitle(e.target.value)}
                                        />
                                    </div>
                                    <div className="form-group mb-6">
                                        <input
                                            type="text"
                                            className="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                            id="exampleInput7"
                                            placeholder="Last Name"
                                            // value={title}
                                            // onChange={(e) => setTitle(e.target.value)}
                                        />
                                    </div>
                                    <div className="form-group mb-6">
                                        <input
                                            type="email"
                                            className="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                            id="exampleInput7"
                                            placeholder="Email"
                                            // value={title}
                                            // onChange={(e) => setTitle(e.target.value)}
                                        />
                                    </div>
                                    <div className="form-group mb-6">
                                        <input
                                            type="text"
                                            className="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                            id="exampleInput7"
                                            placeholder="Username"
                                            // value={title}
                                            // onChange={(e) => setTitle(e.target.value)}
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
                                                // onChange={(e) => setPriority(e.target.value)}
                                            >
                                                <option defaultValue>Select Role</option>
                                                <option value="submitter">
                                                    submitter
                                                </option>
                                                <option value="developer">
                                                    developer
                                                </option>
                                                <option value="project-manager">
                                                    project-manager
                                                </option>
                                                <option value="admin">admin</option>
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
                                                // onChange={(e) => setPriority(e.target.value)}
                                            >
                                                <option defaultValue>Locked / Unlocked</option>
                                                <option value={true}>
                                                    true
                                                </option>
                                                <option value={false}>
                                                    false
                                                </option>
                                            </select>
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
                                        // onClick={handleSubmit}
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
                                            Name
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Username
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Email
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Role
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
                                            <UserItem
                                                key={key}
                                                item={itemvar}
                                                userId={itemvar.id}
                                                role={props.role}
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

export default UserCard;
