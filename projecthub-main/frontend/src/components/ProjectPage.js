import React from 'react';
import {useParams} from "react-router-dom";
import "../css/ProjectPage.css"
import {Header} from "./Header";
import IframeResizer from 'iframe-resizer-react'
import Button from "@mui/material/Button";
import SendIcon from '@mui/icons-material/Send';
import {createMuiTheme, ThemeProvider} from "@mui/material/styles";


export function ProjectPage() {
    // Get the userId param from the URL.
    let { projectID } = useParams();

    const themeCustom = createMuiTheme({
        palette: {
            secondary: {
                main: '#e51433',
                light: '#e51433',
                dark: '#e51433'
            }
        }
    })


    return(
    <div id={"container"}>
        <Header/>
        <ThemeProvider theme={themeCustom}>
            <div id={"iframe-overlay-top"}>
                <Button color="secondary" sx={{mt:"50px", ml:"30%", mr:"30%", display:"flex"}} variant="contained" endIcon={<SendIcon />} href={'/projects/'+ projectID +'/showcase'}>
                    Showcase now!
                </Button>
            </div>
        </ThemeProvider>

        {/* Redesigning an iframe is kinda difficult since cross side scripting is not allowed */}
        <IframeResizer
          heightCalculationMethod="bodyScroll"
          inPageLinks
          log
          scrolling="yes"
          src={"https://www.hdm-stuttgart.de/stage/projekt_detail/projekt_details?projekt_ID=" + projectID}
          style={{ width: '1px', minWidth: '100%'}}
        />
    </div>



    );
}