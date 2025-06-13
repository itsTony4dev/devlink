import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import LoginPage from './components/Login'
import SignupPage from './components/Signup'

function App() {
  return (
    <Router>
      <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-900 via-slate-800 to-slate-900 px-4 sm:px-6 lg:px-8">
        {/* Background pattern */}
        <div className="absolute inset-0 bg-grid-pattern opacity-5"></div>
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route path="/signup" element={<SignupPage />} />
          {/* Redirect to login if no specific path is matched */}
          <Route path="*" element={<LoginPage />} /> 
        </Routes>
      </div>
    </Router>
  )
}

export default App
