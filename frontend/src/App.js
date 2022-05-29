import React from "react";
import { Routes, Route, Link } from "react-router-dom";
import Dashboard from "./components/pages/Dashboard";
import Login from "./components/pages/Login";
import Signup from "./components/pages/Signup";
import Tickets from "./components/pages/Tickets";
import UserManagement from "./components/pages/UserManage";
import SystemLogs from "./components/pages/SystemLogs";
import Logout from "./components/pages/Logout";
import ProjectTickets from "./components/pages/ProjectTickets";
function App() {
    return (
        <div className="App">
            <Routes>
                <Route path="/" element={<Login />} />
                <Route path="/dashboard" element={<Dashboard />} />
                <Route path="/signup" element={<Signup />} />
                <Route path="/tickets" element={<Tickets />} />
                <Route path="/user-management" element={<UserManagement />} />
                <Route path="/system-logs" element={<SystemLogs />} />
                <Route path="/logout" element={<Logout />} />
                <Route path="/project/:id" element={<ProjectTickets />} />
            </Routes>
        </div>
    );
}

export default App;
