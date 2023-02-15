import React from 'react';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import Cv from './pages/CV';
import Login from './pages/Login';



const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Cv/>}/>
        <Route path="/login" element={<Login/>}/>
        {/* <Route path="/admin" element={}/> */}
      </Routes>
    </Router>
  )
}

export default App;
