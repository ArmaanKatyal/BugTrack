import React from "react";
import { Link } from "react-router-dom";
import { BsThreeDotsVertical } from "react-icons/bs";
import TicketDeleteModal from "./TicketDeleteModal";

function TicketItem(props) {
    return (
        <tr className="border-b hover:shadow-sm">
            <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {props.item.title}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-wrap">
                {props.item.description}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                {props.item.project_name}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                {props.item.created_by}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                {props.item.priority === "high" && (
                    <div class="flex items-start">
                        <span class="px-3 py-1 rounded-full text-red-700 bg-red-200 font-semibold text-sm flex align-center w-max cursor-pointer active:bg-gray-300 transition duration-300 ease">
                            High
                        </span>
                    </div>
                )}
                {props.item.priority === "medium" && (
                    <div class="flex items-start">
                        <span class="px-3 py-1 rounded-full text-yellow-700 bg-yellow-200 font-semibold text-sm flex align-center w-max cursor-pointer active:bg-gray-300 transition duration-300 ease">
                            Medium
                        </span>
                    </div>
                )}
                {props.item.priority === "low" && (
                    <div class="flex items-start">
                        <span class="px-3 py-1 rounded-full text-green-700 bg-green-200 font-semibold text-sm flex align-center w-max cursor-pointer active:bg-gray-300 transition duration-300 ease">
                            Low
                        </span>
                    </div>
                )}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap flex flex-row gap-2 flex-wrap">
                {props.item.tags.map((tag, key) => (
                    <span key={key} class="px-3 rounded-full text-gray-500 bg-gray-200 font-thin text-sm flex align-center w-max cursor-pointer active:bg-gray-300 transition duration-300 ease">
                        {tag}
                </span>
                ))}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-wrap">
                {props.item.status === "resolved" ? (
                    <span className="text-green-600">{props.item.status}</span>
                ) : (<span>{props.item.status}</span>)}
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
                    {(props.role === "admin" || props.role === "project-manager" || props.role === "developer") && (
                        <>
                            <li>
                                <button
                                    className=" dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-gray-700 hover:bg-gray-100"
                                    data-bs-toggle="modal"
                                    data-bs-target={
                                        "#updatestaticBackdrop" + props.ticketId
                                    }
                                >
                                    Update
                                </button>
                            </li>
                        </>
                    )}
                    {(props.role === "admin" || props.role === "project-manager") && (
                        <li>
                            <button
                                className=" dropdown-item text-sm py-2 px-4 font-normal block w-full whitespace-nowrap bg-transparent text-gray-700 hover:bg-gray-100"
                                data-bs-toggle="modal"
                                data-bs-target={"#deletestaticBackdrop" + props.ticketId}
                            >
                                Delete
                            </button>
                        </li>
                    )}
                </ul>
            </td>
            {/* <UpdateProjectModal
                projectId={props.projectId}
                item={props.item}
                users={props.users}
            /> */}
            <TicketDeleteModal ticketId={props.ticketId} item={props.item} />
        </tr>
    );
}

export default TicketItem;
