import React from 'react';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import Cv from './pages/CV';



const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Cv/>}/>
        <Route path="/admin" element={<Cv/>}/>
      </Routes>
    </Router>
  )
}

export default App;
