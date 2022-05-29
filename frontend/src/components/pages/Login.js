import { FaUser, FaBuilding } from "react-icons/fa";
import { MdLockOutline } from "react-icons/md";
import { useState } from "react";
import axios from "axios";
import { Link } from "react-router-dom";

const apiPath = "http://localhost:8080/api/v1";

function Login() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [companyCode, setCompanyCode] = useState("");
    const [buttonDisabled, setButtonDisabled] = useState(false);
    const [show, setShow] = useState(false);

    const alert = () => {
        setShow(true);
    };

    const disableButton = () => {
        setButtonDisabled(true);
    };

    const login = async () => {
        if (!username || !password) {
            return;
        }

        setButtonDisabled(true);
        try {
            await axios
                .post(`${apiPath}/auth/login`, {
                    username: username,
                    password: password,
                    company_code: companyCode,
                })
                .then((res) => {
                    if (res.status === 200) {
                        document.cookie = `token=${res.data.token}`;
                        window.location.href = "/dashboard";
                    } else {
                        alert();
                        setUsername("");
                        setPassword("");
                        setCompanyCode("");
                    }
                });
        } catch (err) {
            alert();
            setUsername("");
            setPassword("");
            setCompanyCode("");
        }
    };

    return (
        <div className="flex flex-col items-center justify-center min-h-screen py-2 bg-gray-100 font-sans-new">
            {show && (
                <div
                    class="flex p-4 mb-4 text-sm text-yellow-700 bg-yellow-100 rounded-lg dark:bg-yellow-200 dark:text-yellow-800"
                    role="alert"
                >
                    <svg
                        class="inline flex-shrink-0 mr-3 w-5 h-5"
                        fill="currentColor"
                        viewBox="0 0 20 20"
                        xmlns="http://www.w3.org/2000/svg"
                    >
                        <path
                            fill-rule="evenodd"
                            d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
                            clip-rule="evenodd"
                        ></path>
                    </svg>
                    <div>
                        <span class="font-medium">Warning alert!</span> Wrong Username,
                        Password or Company Code.
                    </div>
                </div>
            )}
            <main className="flex flex-col items-center justify-center w-full flex-1 px-20 text-center">
                <div className="bg-white rounded-2xl shadow-2xl flex w-2/3 max-w-4xl">
                    <div className="w-3/5 p-5">
                        <div className="text-left font-bold">
                            <span className="text-black">bugTrack</span>
                        </div>
                        <div className="py-10">
                            <h2 className="text-3xl font-bold text-blue-600">
                                Sign in to Account
                            </h2>
                            <div className="border-2 w-10 border-blue-600 inline-block mb-2"></div>
                            <div className="flex flex-col items-center my-5">
                                <div className="bg-gray-100 w-64 p-2 flex items-center mb-2 outline-none">
                                    <FaUser className="text-gray-400 m-2"></FaUser>
                                    <input
                                        type="text"
                                        name="username"
                                        placeholder="Username"
                                        className="bg-gray-100 outline-none text-sm flex-1 border-transparent focus:border-transparent focus:ring-0"
                                        value={username}
                                        onChange={(e) => setUsername(e.target.value)}
                                    ></input>
                                </div>
                            </div>
                            <div className="flex flex-col items-center my-5">
                                <div className="bg-gray-100 w-64 p-2 flex items-center mb-2">
                                    <FaBuilding className="text-gray-400 m-2"></FaBuilding>
                                    <input
                                        type="text"
                                        name="companyCode"
                                        placeholder="Company Code"
                                        className="bg-gray-100 outline-none text-sm flex-1 border-transparent focus:border-transparent focus:ring-0"
                                        value={companyCode}
                                        onChange={(e) => setCompanyCode(e.target.value)}
                                    ></input>
                                </div>
                            </div>
                            <div className="flex flex-col items-center my-5">
                                <div className="bg-gray-100 w-64 p-2 flex items-center mb-3">
                                    <MdLockOutline className="text-gray-400 m-2"></MdLockOutline>
                                    <input
                                        type="password"
                                        name="password"
                                        placeholder="Password"
                                        className="bg-gray-100 outline-none text-sm flex-1 border-transparent focus:border-transparent focus:ring-0"
                                        value={password}
                                        onChange={(e) => setPassword(e.target.value)}
                                    ></input>
                                </div>
                                <div className="flex w-64 mb-5 justify-between">
                                    <label className="flex items-center text-xs">
                                        <input
                                            type="checkbox"
                                            name="remember"
                                            className="mr-1"
                                        />
                                        Remember me
                                    </label>
                                    <button href="#" className="text-xs text-blue-800">
                                        Forgot password?
                                    </button>
                                </div>
                                <button
                                    className="border-2 border-blue-600 text-blue-600 rounded-full px-12 py-2 inline-block font-semibold hover:bg-blue-600 hover:text-white hover: transition"
                                    onClick={login}
                                    disabled={buttonDisabled}
                                >
                                    Sign in
                                </button>
                            </div>
                        </div>
                    </div>
                    <div className="w-2/5 bg-blue-800 text-white rounded-tr-2xl rounded-br-2xl py-36 px-12">
                        <h2 className="text-3xl font-bold mb-2">Hello, Friend!</h2>
                        <div className="border-2 w-10 border-white inline-block mb-2"></div>
                        <p className="mb-10">
                            Register your company to get started with bugTrack.
                        </p>
                        <Link
                            to={"/signup"}
                            className="border-2 border-white rounded-full px-12 py-2 inline-block font-semibold hover:bg-white hover:text-blue-600 transition"
                        >
                            Sign up
                        </Link>
                    </div>
                </div>
            </main>
        </div>
    );
}

export default Login;
