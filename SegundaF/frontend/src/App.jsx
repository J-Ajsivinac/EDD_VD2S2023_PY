// import { useState } from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import './App.css'
import Login from './pages/Login'
import LandingAdmin from './pages/LandingAdmin'

function App() {

  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/admin" element={<LandingAdmin />} />
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
