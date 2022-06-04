import React from 'react'
import { Link } from 'react-router-dom'
import { BsThreeDotsVertical } from 'react-icons/bs'

function UserItem(props) {
  return (
    <tr className="border-b hover:shadow-sm">
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
                                        "#updatestaticBackdrop" + props.projectId
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
                                data-bs-target={"#deletestaticBackdrop" + props.projectId}
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
                role={props.role}
            /> */}
            {/* <DeleteProjectModal projectId={props.projectId} item={props.item} /> */}
        </tr>
  )
}

export default UserItem