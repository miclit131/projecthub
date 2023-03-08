import React, { useEffect, useState } from 'react';
import { Header } from './Header';
import { ProfileCard } from './ProjectCard';
import $ from 'jquery'
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';


export function Browse() {

    //fetching projects from endpoint in backend
    const [projects, setProjects] = useState({});
    const fetchProjects = () => {fetch('http://localhost:8080/api/projects').then((response) => response.json()).then((json) => setProjects(json));}

    useEffect(() => {
        fetchProjects()
    }, []);

    return(
        <div>
            <Header/>
            <Box sx={{ width: '40%', mr:"auto", ml:"auto" }}>
                <Grid container rowSpacing={1} columnSpacing={{ xs: 1, sm: 2, md: 3 }} direction="row"
                justifyContent="space-between"
                alignItems="center">
                    {
                        //When response is JSON parsable then it will render a Grid and accordingly some Profilecard components
                        !$.isEmptyObject(JSON.parse(JSON.stringify(projects))) 
                        ?
                        JSON.parse(JSON.stringify(projects)).projects.map((element) => (
                            <div>
                                <Grid item xs={12} sx={{}}>
                                    <ProfileCard information={{ name: element.name, stageId: element.stageId}}/>
                                </Grid>
                            </div>))
                        

                        : console.log('rendering ...')
                    }
                </Grid>
            </Box>            
        </div>
    );
}