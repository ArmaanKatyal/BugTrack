import React from "react";
import { Link } from "react-router-dom";
import { FaTicketAlt, FaUserCog, FaBug } from "react-icons/fa";
import { MdSettingsBackupRestore, MdDashboard, MdLogout } from "react-icons/md";
import { useCookies } from "react-cookie";

function Sidebar(props) {
    const [cookies, setCookies] = useCookies(["token", "role"]);

    return (
        <div className="w-60 h-full shadow-2xl bg-white px-1 fixed font-sans-new">
            <div className="mt-10 flex flex-col items-center justify-center">
                <FaBug className="mb-3 text-5xl text-gray-600 hover:animate-pulse" />
                <p className="text-2xl font-semibold">
                    Bug
                    <span className="text-2xl font-semibold text-blue-600">Track</span>
                </p>
                <p className="text-sm mt-2">
                    Role : <span className="text-blue-600">{cookies.role}</span>
                </p>
            </div>
            <ul className="relative mt-10">
                <li className="relative">
                    <Link
                        to="/dashboard"
                        className="flex items-center text-sm py-4 px-6 h-12 overflow-hidden text-gray-700 text-ellipsis whitespace-nowrap rounded hover:text-gray-900 hover:bg-gray-100 transition duration-300 ease-in-out"
                        data-mdb-ripple="true"
                        data-mdb-ripple-color="dark"
                    >
                        <MdDashboard className="mr-2 text-lg" />
                        <span>Dashboard</span>
                    </Link>
                </li>
                <li className="relative">
                    <Link
                        to="/tickets"
                        className="flex items-center text-sm py-4 px-6 h-12 overflow-hidden text-gray-700 text-ellipsis whitespace-nowrap rounded hover:text-gray-900 hover:bg-gray-100 transition duration-300 ease-in-out"
                        data-mdb-ripple="true"
                        data-mdb-ripple-color="dark"
                    >
                        <FaTicketAlt className="mr-2 text-lg" />
                        <span>Tickets</span>
                    </Link>
                </li>
                {cookies.role === "admin" && (
                    <li className="relative">
                        <Link
                            to="/user-management"
                            className="flex items-center text-sm py-4 px-6 h-12 overflow-hidden text-gray-700 text-ellipsis whitespace-nowrap rounded hover:text-gray-900 hover:bg-gray-100 transition duration-300 ease-in-out"
                            data-mdb-ripple="true"
                            data-mdb-ripple-color="dark"
                        >
                            <FaUserCog className="mr-2 text-lg" />
                            <span>User Management</span>
                        </Link>
                    </li>
                )}
                {cookies.role === "admin" && (
                    <li className="relative">
                    <Link
                        to="/system-logs"
                        className="flex items-center text-sm py-4 px-6 h-12 overflow-hidden text-gray-700 text-ellipsis whitespace-nowrap rounded hover:text-gray-900 hover:bg-gray-100 transition duration-300 ease-in-out"
                        data-mdb-ripple="true"
                        data-mdb-ripple-color="dark"
                    >
                        <MdSettingsBackupRestore className="mr-2 text-lg" />
                        <span>System Logs</span>
                    </Link>
                </li>
                )}
            </ul>
            <Link
                to="/logout"
                className="flex relative items-center text-sm py-4 px-6 h-12 overflow-hidden text-red-700 text-ellipsis whitespace-nowrap rounded hover:bg-gray-100 transition duration-300 ease-in-out"
            >
                <MdLogout className="mr-2 text-lg" />
                <span>Logout</span>
            </Link>
        </div>
    );
}

export default Sidebar;
