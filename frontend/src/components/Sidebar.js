import React from "react";
import { Link } from "react-router-dom";
import { FaTicketAlt, FaUserCog } from "react-icons/fa";
import { MdSettingsBackupRestore, MdDashboard, MdLogout } from "react-icons/md";

const menuItems = [
    { id: 1, name: "Dashboard", link: "/dashboard", icon: MdDashboard },
    { id: 2, name: "Tickets", link: "/tickets", icon: FaTicketAlt },
    { id: 3, name: "User Management", link: "/user-management", icon: FaUserCog },
    { id: 4, name: "System Logs", link: "/system-logs", icon: MdSettingsBackupRestore },
];

function Sidebar(props) {
    return (
        <div className="w-60 h-full shadow-2xl bg-white px-1 absolute font-sans-new">
            <ul className="relative mt-10">
                {menuItems.map((item, key) => {
                    return (
                        <li className="relative" key={key}>
                            <Link
                                to={item.link}
                                className="flex items-center text-sm py-4 px-6 h-12 overflow-hidden text-gray-700 text-ellipsis whitespace-nowrap rounded hover:text-gray-900 hover:bg-gray-100 transition duration-300 ease-in-out"
                                data-mdb-ripple="true"
                                data-mdb-ripple-color="dark"
                            >
                                <item.icon className="mr-2 text-lg" />
                                <span>{item.name}</span>
                            </Link>
                        </li>
                    );
                })}
            </ul>
            <Link
                to="/logout"
                className="flex relative items-center text-sm py-4 px-6 h-12 overflow-hidden text-gray-700 text-ellipsis whitespace-nowrap rounded hover:text-red-600 hover:bg-gray-100 transition duration-300 ease-in-out"
            >
                <MdLogout className="mr-2 text-lg" />
                <span>Logout</span>
            </Link>
        </div>
    );
}

export default Sidebar;
