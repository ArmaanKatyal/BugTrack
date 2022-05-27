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
  const login = async () => {
    if (!username || !password) {
      return;
    }

    setButtonDisabled(true);
    try {
      await axios.post(`${apiPath}/auth/login`, {
        username: username,
        password: password,
        company_code: companyCode,
      }).then((res) => {
        if (res.status === 200) {
          document.cookie = `token=${res.data.token}`;
          window.location.href = "/dashboard";
        } else {
          alert("Invalid credentials");
          setUsername("");
          setPassword("");
          setCompanyCode("");
          setButtonDisabled(false);
        }
      });
    } catch (err) {
      setUsername("");
      setPassword("");
      setCompanyCode("");
      setButtonDisabled(false);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen py-2 bg-gray-100">
      <main className="flex flex-col items-center justify-center w-full flex-1 px-20 text-center">
        <div className="bg-white rounded-2xl shadow-2xl flex w-2/3 max-w-4xl">
          <div className="w-3/5 p-5">
            <div className="text-left font-bold">
              <span className="text-black">bugTrack</span>
            </div>
            <div className="py-10">
              <h2 className="text-3xl font-bold text-green-600">
                Sign in to Account
              </h2>
              <div className="border-2 w-10 border-green-600 inline-block mb-2"></div>
              <div className="flex flex-col items-center my-5">
                <div className="bg-gray-100 w-64 p-2 flex items-center mb-2">
                  <FaUser className="text-gray-400 m-2"></FaUser>
                  <input
                    type="text"
                    name="username"
                    placeholder="Username"
                    className="bg-gray-100 outline-none text-sm flex-1"
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
                    className="bg-gray-100 outline-none text-sm flex-1"
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
                    className="bg-gray-100 outline-none text-sm flex-1"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                  ></input>
                </div>
                <div className="flex w-64 mb-5 justify-between">
                  <label className="flex items-center text-xs">
                    <input type="checkbox" name="remember" className="mr-1" />
                    Remember me
                  </label>
                  <button href="#" className="text-xs text-green-800">
                    Forgot password?
                  </button>
                </div>
                <button
                  className="border-2 border-green-600 text-green-600 rounded-full px-12 py-2 inline-block font-semibold hover:bg-green-600 hover:text-white hover: transition"
                  onClick={login}
                  disabled={buttonDisabled}
                >
                  Sign in
                </button>
              </div>
            </div>
          </div>
          <div className="w-2/5 bg-green-600 text-white rounded-tr-2xl rounded-br-2xl py-36 px-12">
            <h2 className="text-3xl font-bold mb-2">Hello, Friend!</h2>
            <div className="border-2 w-10 border-white inline-block mb-2"></div>
            <p className="mb-10">
              Fill up personal information and start journey with us
            </p>
            <Link
              to={"/signup"}
              className="border-2 border-white rounded-full px-12 py-2 inline-block font-semibold hover:bg-white hover:text-green-600 transition"
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