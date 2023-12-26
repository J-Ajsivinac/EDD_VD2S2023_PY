// import { useState } from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import './App.css'
import Login from './pages/Login'
import LandingAdmin from './pages/LandingAdmin'
import LoadCourses from './pages/LoadCourses'
import AcceptBooks from './pages/AcceptBooks'
import LoadStudents from './pages/LoadStudents'
import Report from './pages/Report'
import TutorBooks from './pages/TutorBooks'
import StudentBooks from './pages/StudentBooks'
import Courses from './pages/Courses'
import Pubs from './pages/Pubs'
import { AuthProvider } from './context/authContext'
import LoadTutor from './pages/LoadTutor'
import CreatePub from './pages/CreatePub'

function App() {

  return (
    <>
      <AuthProvider>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Login />} />
            {/* Administrador */}
            <Route path="/admin/index" element={<LandingAdmin />} />
            <Route path="/admin/load/courses" element={<LoadCourses />} />
            <Route path="/admin/load/students" element={<LoadStudents />} />
            <Route path="/admin/load/tutor" element={<LoadTutor />} />
            <Route path="/admin/accept" element={<AcceptBooks />} />
            <Route path="/admin/report/:graph" element={<Report />} />
            {/* Tutor */}
            <Route path="/tutor/books" element={<TutorBooks />} />
            <Route path="/tutor/pubs" element={<CreatePub />} />
            {/* Estudiante */}
            <Route path="/student/pubs" element={<Pubs />} />
            <Route path="/student/courses" element={<Courses />} />
            <Route path="/student/books" element={<StudentBooks />} />
          </Routes>
        </BrowserRouter>
      </AuthProvider>

    </>
  )
}

export default App
