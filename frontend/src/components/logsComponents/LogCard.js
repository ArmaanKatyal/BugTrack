import React from 'react'
import LogItem from './LogItem'

function LogCard(props) {
  return (
    <div className="absolute w-11/12 flex flex-col bg-white h-auto top-40 rounded-lg shadow-xl font-sans-new">
        <div className="flex flex-col">
                <div className="overflow-x-auto sm:-mx-6 lg:-mx-8">
                    <div className="py-0 inline-block min-w-full sm:px-6 lg:px-8">
                        <div className="overflow-hidden rounded-tl-lg rounded-tr-lg shadow-lg">
                            <table className="min-w-full">
                                <thead className="border-b bg-gray-100 shadow-lg">
                                    <tr>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Type
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Author
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Date
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Description
                                        </th>
                                        <th
                                            scope="col"
                                            className="text-sm font-medium text-gray-900 px-6 py-4 text-left"
                                        >
                                            Table
                                        </th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {props.data &&
                                        props.data.map((itemvar, key) => (
                                            <LogItem
                                                key={key}
                                                item={itemvar}
                                            />
                                        ))}
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
    </div>
  )
}

export default LogCard