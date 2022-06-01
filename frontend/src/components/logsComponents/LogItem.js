import React from "react";

function LogItem(props) {
    return (
        <tr className="border-b hover:shadow-sm">
            <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                    {props.item.type}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                {props.item.author}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                {props.item.date}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                {props.item.description}
            </td>
            <td className="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                {props.item.table}
            </td>
        </tr>
    );
}

export default LogItem;
