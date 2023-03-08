import React from 'react';
import '../css/Header.css'
import logo from '../assets/playbtn_red.svg'
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import HelpOutlineIcon from '@mui/icons-material/HelpOutline';
import {IconButton} from "@mui/material";
import {createMuiTheme, ThemeProvider} from "@mui/material/styles";
import Modal from '@mui/material/Modal';
import ButtonGroup from '@mui/material/ButtonGroup';
import Button from '@mui/material/Button';
import RunningWithErrorsOutlinedIcon from '@mui/icons-material/RunningWithErrorsOutlined';

export function Header() {

    //toggle and event handlers
    const [open, setOpen] = React.useState(false);
    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);

    const style = {
        position: 'absolute',
        top: '50%',
        left: '50%',
        transform: 'translate(-50%, -50%)',
        width: 400,
        bgcolor: 'background.paper',
        boxShadow: 24,
        p: 4,
      };

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
        <div>
            <ThemeProvider theme={themeCustom}>
                <div className={"containerHeader"} >
                    <Box sx={{ flexGrow: 0 }}>
                        <AppBar position="static" sx={{ backgroundColor: '#3E4847',  color: ''}}>
                            <Toolbar>
                            <ButtonGroup variant="text" aria-label="text button group">
                                <a href="/">
                                    <img src={logo} alt={"Logo of ProjectHub"} height={"40px"}/>
                                </a>
                                <Typography
                                    variant="h5"
                                    noWrap
                                    component="a"
                                    href="/"
                                    sx={{
                                        ml: 2,
                                        mt: '4px',
                                        flexGrow: 0,
                                        fontFamily: 'Roboto',
                                        fontWeight: 400,
                                        letterSpacing: '.3rem',
                                        color: 'inherit',
                                        textDecoration: 'none',
                                    }}
                                >
                                    ProjectHub
                                </Typography>
                                {/* <Button color="secondary" href='/browse' sx={{ flexGrow: 0 }}>Discover</Button> */}
                            </ButtonGroup>
                            <IconButton color="secondary" aria-label="browse" href='/browse' sx={{flexGrow: 1, justifyContent: 'end'}}>
                                <RunningWithErrorsOutlinedIcon />
                            </IconButton>
                            <IconButton color="secondary" aria-label="information" onClick={handleOpen} sx={{flexGrow: 0, justifyContent: 'end'}}>
                                <HelpOutlineIcon />
                            </IconButton>
                                <Modal
                                    open={open}
                                    onClose={handleClose}
                                    aria-labelledby="modal-modal-title"
                                    aria-describedby="modal-modal-description"
                                    >
                                    <Box sx={style}>
                                        <Typography id="modal-modal-title" variant="h6" component="h2">
                                        🚀 Quick Information
                                        </Typography>
                                        <Typography id="modal-modal-description" sx={{ mt: 2 }}>
                                        Don't worry we've got your back! So you came here to <br></br>
                                        deploy your game or web based project right? Then let<br></br>
                                        us guide you through the following steps:<br></br>
                                        <br></br>
                                        {/* Something needs to be changend regarding the styling here - no more nbsp! */}
                                        💡 Step 1:&nbsp;&nbsp;Enter your desired project name.<br></br>
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;This property will be displayed<br></br>
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;at the browse page.<br></br>
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<br></br>
                                        💡 Step 2:&nbsp;&nbsp;Enter your Gitlab Repository URL.<br></br>
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Don't forget to set your Repository<br></br>
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Visibility to public or use an<br></br>
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Access Token (read only).<br></br>
                                        <br></br>
                                        💡 Step 3:&nbsp;&nbsp;Enter your HdM Stage ID.<br></br>
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Well you can get this ID at the<br></br>
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Stage. Just have a look at your<br></br>
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Stage project page. (read only).<br></br>
                                        <br></br>
                                        💡 Step 4:&nbsp;&nbsp;Press the button.<br></br>
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Our deployment process will take<br></br>
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;some time. You will find your<br></br>
                                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;embedded project at the browse page.<br></br>
                                        <br></br>

                                        </Typography>
                                    </Box>
                                </Modal>
                            </Toolbar>
                        </AppBar>
                    </Box>
                </div>
            </ThemeProvider>

        </div>

    );
}