import React from "react";
import { Link } from "react-router-dom";
import { BsThreeDotsVertical } from "react-icons/bs";
import UserDeleteModal from "./UserDeleteModal";
import { CgLock, CgLockUnlock } from "react-icons/cg";
import UserLockModal from "./UserLockModal";
import UserUnLockModal from "./UserUnLockModal";

function UserItem(props) {
    return (
        // <tr className="border-b hover:shadow-sm hover:bg-gray-100">
        <tr className={props.item.locked === true ? ("border-b hover:shadow-sm bg-red-100") : ("border-b hover:shadow-sm hover:bg-gray-100")}>
            <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                <Link to={`/profile/${props.item.username}`} className="text-purple-600">
                    <span>{`${props.item.first_name} ${props.item.last_name}`}</span>
                </Link>
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                {props.item.username}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                {props.item.email}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                {props.item.role}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                <button
                    className="dropdown-toggle py-2.5 text-black font-medium text-xs leading-tight uppercase roundedhover:shadow-lgfocus:shadow-lg focus:outline-none focus:ring-0active:shadow-lg active:text-white transition duration-150 ease-in-out flex items-center whitespace-nowrap"
                    type="button"
                    id="dropdownMenuButton1e"
                    data-bs-toggle="dropdown"
                    aria-expanded="false"
                >
                    <BsThreeDotsVertical />
                </button>
                <ul
                    className=" dropdown-menu min-w-max absolute hidden bg-white text-base z-50 float-left py-2 list-none text-left rounded-lg shadow-lg mt-1 m-0 bg-clip-padding border-none"
                    aria-labelledby="dropdownMenuButton1e"
                >
                    {(props.role === "admin" || props.role === "project-manager") && (
                        <>
                            <li>
                                <button
                                    className=" dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-gray-700 hover:bg-gray-100"
                                    data-bs-toggle="modal"
                                    data-bs-target={
                                        "#updatestaticBackdrop" + props.item.username
                                    }
                                >
                                    Update
                                </button>
                            </li>
                        </>
                    )}
                    {props.role === "admin" && (
                        <li>
                            <button
                                className=" dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-gray-700 hover:bg-gray-100"
                                data-bs-toggle="modal"
                                data-bs-target={
                                    "#deletestaticBackdrop" + props.item.username
                                }
                            >
                                Delete
                            </button>
                        </li>
                    )}
                    {props.role === "admin" && props.item.locked === false && (
                        <li>
                            <button
                                className=" dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-gray-700 hover:bg-gray-100"
                                data-bs-toggle="modal"
                                data-bs-target={
                                    "#lockstaticBackdrop" + props.item.username
                                }
                            >
                                <div className="flex flex-row gap-1 justify-center items-center">
                                    <CgLock /> Lock
                                </div>
                            </button>
                        </li>
                    )}
                    {props.role === "admin" && props.item.locked === true && (
                        <li>
                            <button
                                className=" dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-gray-700 hover:bg-gray-100"
                                data-bs-toggle="modal"
                                data-bs-target={
                                    "#unlockstaticBackdrop" + props.item.username
                                }
                            >
                                <div className="flex flex-row gap-1 justify-center items-center">
                                  <CgLockUnlock /> Unlock
                                </div>
                            </button>
                        </li>
                    )}
                </ul>
            </td>
            {/* <UpdateProjectModal
                projectId={props.projectId}
                item={props.item}
                users={props.users}
                role={props.role}
            /> */}
            <UserLockModal item={props.item} />
            <UserUnLockModal item={props.item} />
            <UserDeleteModal item={props.item} />
        </tr>
    );
}

export default UserItem;