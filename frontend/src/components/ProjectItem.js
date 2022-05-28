import React from "react";
import { Link } from "react-router-dom";
import { BsThreeDotsVertical } from "react-icons/bs";

function ProjectItem(props) {
  return (
    <tr class="border-b">
      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
        <Link to={`/project/${props.projectId}`} className="text-purple-600">
          {props.item.title}
        </Link>
      </td>
      <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
        {props.item.description}
      </td>
      <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
        <button
          className="dropdown-toggle
          px-6
          py-2.5
          text-black
          font-medium
          text-xs
          leading-tight
          uppercase
          rounded
        hover:shadow-lg
        focus:shadow-lg focus:outline-none focus:ring-0
     active:shadow-lg active:text-white
          transition
          duration-150
          ease-in-out
          flex
          items-center
          whitespace-nowrap
        "
          type="button"
          id="dropdownMenuButton1e"
          data-bs-toggle="dropdown"
          aria-expanded="false"
        >
          <BsThreeDotsVertical />
        </button>
        <ul
          class="
          dropdown-menu
          min-w-max
          absolute
          hidden
          bg-white
          text-base
          z-50
          float-left
          py-2
          list-none
          text-left
          rounded-lg
          shadow-lg
          mt-1
          m-0
          bg-clip-padding
          border-none
        "
          aria-labelledby="dropdownMenuButton1e"
        >
          {(props.role === "admin" || props.role === "project-manager") && (
              <li>
                <a
                  class="
                  dropdown-item
                  text-sm
                  py-2
                  px-4
                  font-normal
                  block
                  w-full
                  whitespace-nowrap
                  bg-transparent
                  text-gray-700
                  hover:bg-gray-100
                "
                  href="#"
                >
                  Update
                </a>
              </li>
            )}
          {props.role === "admin" && (
            <li>
              <a
                class="
              dropdown-item
              text-sm
              py-2
              px-4
              font-normal
              block
              w-full
              whitespace-nowrap
              bg-transparent
              text-gray-700
              hover:bg-gray-100
            "
                href="#"
              >
                Delete
              </a>
            </li>
          )}
        </ul>
      </td>
    </tr>
  );
}

export default ProjectItem;
