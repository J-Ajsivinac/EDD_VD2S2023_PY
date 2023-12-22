// import { useState } from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import './App.css'
import Login from './pages/Login'
import LandingAdmin from './pages/LandingAdmin'
import LoadCourses from './pages/LoadCourses'
import AcceptBooks from './pages/AcceptBooks'
import LoadStudents from './pages/LoadStudents'
import Report from './pages/Report'

function App() {

  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/admin" element={<LandingAdmin />} />
          <Route path="/load/courses" element={<LoadCourses />} />
          <Route path="/load/students" element={<LoadStudents />} />
          <Route path="/accept" element={<AcceptBooks />} />
          <Route path="/report" element={<Report />} />
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
