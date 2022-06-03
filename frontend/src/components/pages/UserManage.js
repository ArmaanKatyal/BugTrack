import React from "react";
import Sidebar from "../Sidebar";
import UserCard from "../userComponents/UserCard";
import { useCookies } from "react-cookie";
import axios from "axios";

const apiPath = "http://localhost:8080/api/v1";

function UserManage() {
    const [loading, setLoading] = React.useState(true);
    const [cookie, setCookie] = useCookies(["token"]);
    const [users, setUsers] = React.useState([]);
    const [lockedUsers, setLockedUsers] = React.useState([]);

    const getUsers = async () => {
      try {
        var config = {
          headers: {
            token: cookie.token,
          },
        };
        await axios.get(`${apiPath}/user?role=`, config).then((res) => {
          if (res.status === 200) {
            setUsers(res.data);
            toggleLoading();
          } else {
            window.location.href = "/";
          }
        });
      } catch (error) {
        window.location.href = "/";
      }
    };

    const lockedUsersFunc = async () => {};

    const toggleLoading = () => {
        setLoading(!loading);
    };

    React.useEffect(() => {
        if (cookie.token) {
            getUsers();
        } else {
            window.location.href = "/";
        }
    }, []);
    return (
        <>
            <Sidebar />
            <div className="flex flex-col bg-blue-500 h-72 pl-60">
                <div className="relative flex flex-col items-center">
                    <div className="flex flex-row h-full mt-12 ml-36 self-start gap-10">
                        <h1 className="text-4xl text-white font-sans-new">
                            User Management
                        </h1>
                        {loading && (
                            <div className="">
                                <div
                                    className="spinner-border animate-spin text-white inline-block w-10 h-10 border-3 rounded-full self-start"
                                    role="status"
                                >
                                    <span className="visually-hidden">Loading...</span>
                                </div>
                            </div>
                        )}
                    </div>
                    <UserCard role={cookie.role} data={users} />
                </div>
            </div>
        </>
    );
}

export default UserManage;
