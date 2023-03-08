import React from 'react';
import {Header} from "./Header";
import logoHDM from "../assets/hdmlogo.jpg";
import "../css/Home.css"
import {InputAdornment, TextField} from "@mui/material";
import Box from "@mui/material/Box";
import magic from '../assets/magic_button.svg';
import Button from "@mui/material/Button";
import SendIcon from '@mui/icons-material/Send';
import {createMuiTheme, ThemeProvider} from "@mui/material/styles";
import play from '../assets/playStage.svg';
import Snackbar from '@mui/material/Snackbar';
import {useNavigate} from 'react-router-dom';


export function Home() {

    const navigate = useNavigate();
    const [projectName, setProjectName] = React.useState('');
    const [projectRepo, setProjectRepo] = React.useState('');
    const [projectStageId, setProjectStageId] = React.useState(0);

    function handleSubmit(event) {
        
        
        let jsonData ={stageId: parseInt(projectStageId), name: projectName, gitUrl: projectRepo}
        console.log( JSON.stringify(jsonData) ); 

        //fetching data from endpoint
        fetch('http://localhost:8080/api/project', {
            method: 'POST', // or 'PUT'
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(jsonData),
            })
        .then(response => response.json())
        .then(jsonData => {
            console.log('Success:', jsonData);
        })
        .catch((error) => {
            console.error('Error:', error);
        });

    }

    const themeCustom = createMuiTheme({
        palette: {
            secondary: {
                main: '#e51433',
                light: '#e51433',
                dark: '#e51433'
            }
        }
    })

    const [open, setOpen] = React.useState(false);

    const handleClick = () => {
      setOpen(true);
      handleSubmit()
      document.getElementById("form1").reset()
    };
  
    const handleClose = (reason) => {
      if (reason === 'clickaway') {
        return;
      }
  
      setOpen(false);
      navigate('/browse')
    };


    return(

        <div>
            <Header/>
            <ThemeProvider theme={themeCustom}>
                <p className={"text"} >Embed your unity game!</p>
                <div className="wrapper">
                    <img src={logoHDM} alt="hdm logo" style={{width: '200px', height: '182px'}}/>
                    <div className="vl"></div>
                    <form id={"form1"} onSubmit={handleSubmit}>
                        <Box sx={{mt:"5px", marginLeft:"50px"}} className={"box"}>
                            <TextField id="input-with-icon-textfield1"
                                    label="Project name"
                                    color="secondary"
                                    
                                    onInput={ e=>setProjectName(e.target.value)}
                                    InputProps={{
                                        startAdornment: (
                                            <InputAdornment position="start">
                                                <img src={magic} alt="project name" width={"20px"}/>
                                            </InputAdornment>
                                        ),
                                    }}
                                    variant="filled" />
                            <Box sx={{mt:"10px"}}>
                                <TextField id="input-with-icon-textfield2"
                                        color="secondary"
                                        label="Repository URL"
                                        
                                        onInput={ e=>setProjectRepo(e.target.value)}
                                        InputProps={{
                                            startAdornment: (
                                                <InputAdornment position="start">
                                                    <img src={"https://cdn-icons-png.flaticon.com/512/25/25231.png"} alt="github logo" width={"20px"}/>
                                                </InputAdornment>
                                            ),
                                        }}
                                        variant="filled" />
                            </Box>
                            <Box sx={{mt:"10px"}}>
                                <TextField id="input-with-icon-textfield3"
                                        color="secondary"
                                        label="Stage ID"
                                        
                                        onInput={ e=>setProjectStageId(e.target.value)}
                                        InputProps={{
                                            startAdornment: (
                                                <InputAdornment position="start">
                                                    <img src={play} alt="stage logo" width={"20px"}/>
                                                </InputAdornment>
                                            ),
                                        }}
                                        variant="filled" />
                            </Box>
                            <Button type={"submit"} color="secondary" sx={{mt:"10px"}} variant="contained" endIcon={<SendIcon />} onClick={handleClick} disabled={open}>
                                Embed now!
                            </Button>
                            <Snackbar
                                open={open}
                                autoHideDuration={4000}
                                onClose={handleClose}
                                message="Please wait 5 minutes. Your project will be deployed soon... "
                                ContentProps={{
                                    sx: {
                                      background: "#e51433"
                                    }
                                  }}
                            />
                        </Box>
                    </form>
                </div>
            </ThemeProvider>
        </div>

    );
}