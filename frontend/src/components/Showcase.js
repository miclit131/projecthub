import React, { useEffect, useState } from 'react';
import {useParams} from "react-router-dom";
import $ from 'jquery'

export function Showcase() {
    let { projectID } = useParams();
    
    const [project, setProject] = useState({});
    const fetchProject = async () => { await fetch('http://localhost:8080/api/project/' + projectID).then((response) => response.json()).then((json) => setProject(json));}

    useEffect(() => {
        fetchProject()
    }, []);

    return(
        <>
            <iframe title={projectID} src={'http://kubernetes.docker.internal/'} allow={"fullscreen"} allowFullScreen={"true"}></iframe>       
        </>
    );
}