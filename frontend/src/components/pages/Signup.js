import React from "react";
import {useState} from "react";
import { FaUser, FaBuilding, FaEnvelope } from "react-icons/fa";
import { MdLockOutline } from "react-icons/md";
import axios from "axios";

const apiPath = "http://localhost:8080/api/v1";

function Signup() {
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [repeatPassword, setRepeatPassword] = useState("");
    const [companyName, setCompanyName] = useState("");
    const [companyCode, setCompanyCode] = useState("");
    const [buttonDisabled, setButtonDisabled] = useState(false);

    const disableButton = () => {
        setButtonDisabled(true);
      };

    const signup = async () => {
        if (!firstName || !lastName || !username || !email|| !password || !repeatPassword || !companyName || !companyCode) {
            return;
        }
        if (password !== repeatPassword) {
            return;
        }
        setButtonDisabled(true);
        try {
            await axios.post(`${apiPath}/auth/signup`, {
                first_name: firstName,
                last_name: lastName,
                username: username,
                password: password,
                email: email,
                company_name: companyName,
                company_code: companyCode,
            }).then((res) => {
                if (res.status === 200) {
                    window.location.href = "/";
                } else {
                    setFirstName("");
                    setLastName("");
                    setEmail("");
                    setUsername("");
                    setPassword("");
                    setRepeatPassword("");
                    setCompanyName("");
                    setCompanyCode("");
                    disableButton();
                }
            });
        } catch (err) {
            setFirstName("");
            setLastName("");
            setEmail("");
            setUsername("");
            setPassword("");
            setRepeatPassword("");
            setCompanyName("");
            setCompanyCode("");
            disableButton();
        }

    };

  return (
    <div className="flex flex-col items-center bg-gray-100 py-2 min-h-screen font-sans-new">
      <main className="flex flex-col items-center justify-center w-full flex-1 px-20 text-center py-10">
          <div className="bg-white rounded-2xl shadow-2xl flex items-center flex-col max-w-4xl w-2/3">
            <div className="py-10">
                <h2 className="text-3xl font-bold text-green-600 mb-2">Let's get started</h2>
                <div className="border-2 w-10 border-green-600 inline-block mb-2"></div>
            </div>
            <div className="flex flex-col items-center my-5">
                <div className="bg-gray-100 w-96 p-2 flex items-center mb-2 hover:shadow-xl hover:transition">
                    <FaUser className="text-gray-400 m-2"></FaUser>
                    <input
                        type="text"
                        name="firstName"
                        placeholder="First Name"
                        className="bg-gray-100 outline-none text-sm flex-1 border-transparent focus:border-transparent focus:ring-0"
                        value={firstName}
                        onChange={(e) => setFirstName(e.target.value)}
                    ></input>
                </div>
                <div className="bg-gray-100 w-96 p-2 flex items-center mb-2  hover:shadow-xl hover:transition">
                    <FaUser className="text-gray-400 m-2"></FaUser>
                    <input
                        type="text"
                        name="lastName"
                        placeholder="Last Name"
                        className="bg-gray-100 outline-none text-sm flex-1 border-transparent focus:border-transparent focus:ring-0"                        
                        value={lastName}
                        onChange={(e) => setLastName(e.target.value)}
                    ></input>
                </div>
                <div className="bg-gray-100 w-96 p-2 flex items-center mb-2  hover:shadow-xl hover:transition">
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
                <div className="bg-gray-100 w-96 p-2 flex items-center mb-2  hover:shadow-xl hover:transition">
                    <FaEnvelope className="text-gray-400 m-2"></FaEnvelope>
                    <input
                        type="email"
                        name="email"
                        placeholder="Email"
                        className="bg-gray-100 outline-none text-sm flex-1 border-transparent focus:border-transparent focus:ring-0"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    ></input>
                </div>
                <div className="bg-gray-100 w-96 p-2 flex items-center mb-2 hover:shadow-xl hover:transition">
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
                <div className="bg-gray-100 w-96 p-2 flex items-center mb-2 hover:shadow-xl hover:transition">
                    <MdLockOutline className="text-gray-400 m-2"></MdLockOutline>
                    <input
                        type="password"
                        name="repeatPassword"
                        placeholder="Repeat Password"
                        className="bg-gray-100 outline-none text-sm flex-1 border-transparent focus:border-transparent focus:ring-0"
                        value={repeatPassword}
                        onChange={(e) => setRepeatPassword(e.target.value)}
                    ></input>
                </div>
                <div className="bg-gray-100 w-96 p-2 flex items-center mb-2 hover:shadow-xl hover:transition">
                    <FaBuilding className="text-gray-400 m-2"></FaBuilding>
                    <input
                        type="text"
                        name="companyName"
                        placeholder="Company Name"
                        className="bg-gray-100 outline-none text-sm flex-1 border-transparent focus:border-transparent focus:ring-0"
                        value={companyName}
                        onChange={(e) => setCompanyName(e.target.value)}
                    ></input>
                </div>
                <div className="bg-gray-100 w-96 p-2 flex items-center mb-5  hover:shadow-xl hover:transition">
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
            <button
                  className="border-2 m-10 border-green-600 text-green-600 rounded-full px-12 py-2 inline-block font-semibold hover:bg-green-600 hover:text-white hover: transition"
                  onClick={signup}
                  disabled={buttonDisabled}
                >
                  Sign up
            </button>
            </div>
          </div>
      </main>
    </div>
  );
}

export default Signup;
