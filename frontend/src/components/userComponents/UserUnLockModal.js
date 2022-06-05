import React from 'react'
import axios from 'axios';
import { useCookies } from 'react-cookie';

const apiPath = 'http://localhost:8080/api/v1';

function UserUnLockModal(props) {
  const [cookie, setCookie] = useCookies(["token"]);
    const [username, setUsername] = React.useState(props.item.username);
    const [userInput, setUserInput] = React.useState("");
    const [alertModal, setAlertModal] = React.useState(false);

    const toggleAlertModal = () => {
        setAlertModal(!alertModal);
    };

    const unlockUser = async () => {
        if (cookie.token) {
            if (userInput === username) {
                try {
                    var config = {
                        headers: {
                            "Content-Type": "application/json",
                            token: cookie.token,
                        },
                    };
                    await axios
                        .get(apiPath + "/user/unlock/" + props.item.username, config)
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
                toggleAlertModal();
            }
        } else {
            window.location.href = "/";
        }
    };

    return (
        <div
            className="modal fade fixed top-0 left-0 hidden w-full h-full outline-none overflow-x-hidden overflow-y-auto font-sans-new"
            id={"unlockstaticBackdrop" + props.item.username}
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
                            Unlock User ⚠️
                        </h5>
                        <button
                            type="button"
                            className="btn-close box-content w-4 h-4 p-1 text-black border-none rounded-none opacity-50 focus:shadow-none focus:outline-none focus:opacity-100 hover:text-black hover:opacity-75 hover:no-underline"
                            data-bs-dismiss="modal"
                            aria-label="Close"
                        ></button>
                    </div>
                    <div className="flex flex-col items-start ml-5 gap-3 px-5">
                        <p className="mt-3">
                            Are you sure you want to Unlock user{" "}
                            <span className="font-bold text-purple-500">
                                {props.item.username}
                            </span>{" "}
                            ?
                        </p>
                        <p>
                            The user will gain back the ability to connect to the application.
                        </p>
                        <p>
                            Type{" "}
                            <span className="text-red-500">{props.item.username}</span> to
                            authorize the operation:
                        </p>
                    </div>
                    <div className="modal-body relative p-4">
                        <div className="form-group">
                            <input
                                type="text"
                                className="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                                id="exampleInput7"
                                placeholder="Username"
                                value={userInput}
                                onChange={(e) => setUserInput(e.target.value)}
                            />
                        </div>
                    </div>
                    {alertModal && (
                        <div className="mb-2 ml-3">
                            <p className="text-red-700">
                                Entered Username is not correct *
                            </p>
                        </div>
                    )}
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
                            onClick={unlockUser}
                            className="inline-block px-6 py-2.5 bg-red-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-red-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out ml-1"
                        >
                            Unlock
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default UserUnLockModal