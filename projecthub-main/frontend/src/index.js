import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router } from 'react-router-dom';
import App from './App';

// Top level of this application - it starts the app component and sets the router
ReactDOM.render(
  <Router>
    <App />
  </Router>,
  document.getElementById('root')
);
