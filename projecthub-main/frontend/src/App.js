import React from 'react';
import { Routes, Route } from 'react-router-dom';
import './css/App.css';
import {ProjectPage} from './components/ProjectPage';
import {Home} from "./components/Home";
import {Browse} from "./components/Browse";
import { Showcase } from './components/Showcase';


export default function App() {
    return (
        <Routes>
            <Route path={"/"} element={<Home/>}/>
            <Route path="projects">
                <Route path=":projectID" element={<ProjectPage />} />
                <Route path=":projectID/showcase" element={<Showcase />} />
            </Route>
            <Route path={"browse"} element={<Browse/>}/>
            <Route path={"*"} element={<h1>Error: Not found.</h1>}/>
        </Routes>
    );
}