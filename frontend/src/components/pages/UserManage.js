import React from 'react'
import Sidebar from '../Sidebar'

function UserManage() {
    React.useEffect(() => {
        if (document.cookie.split("=")[1]) {
            // do nothing
        } else {
          window.location.href = "/";
        }
      }, []);
  return (
    <>
        <Sidebar />
        <div className="flex flex-col bg-blue-500 h-72 pl-60">
            <div className="flex flex-col h-full mt-12 pl-36">
                <h1 className="text-3xl text-white">User Management</h1>
            </div>
        </div>
      </>
  )
}

export default UserManage