import React from 'react'
import { Routes, Route, Link } from 'react-router-dom'
import Dashboard from './components/pages/Dashboard'
import Login from './components/pages/Login'
import Signup from './components/pages/Signup'

function App() {
  return (
    <div className='App'>
      <Routes>
        <Route path='/' element={<Login />} />
        <Route path='/dashboard' element={<Dashboard />} />
        <Route path='/signup' element={<Signup />} />
      </Routes>
    </div>
  )
}

export default App