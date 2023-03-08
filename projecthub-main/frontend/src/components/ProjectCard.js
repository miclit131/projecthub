import React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import SendIcon from '@mui/icons-material/Send';
import {createMuiTheme, ThemeProvider} from "@mui/material/styles";
import background from "../assets/background.png"


export function ProfileCard({information}) {
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
            <Card elevation={12} sx={{ maxWidth: 345, backgroundColor:"", mt: "50px", ml:"auto", mr:"auto",  backgroundImage: `url(${background})`, border:"1px solid #3E4847"}}>
                <CardContent>
                    <Typography gutterBottom variant="h5" component="div" color={"common.black"}>
                    {information.name}
                    </Typography>
                </CardContent>
                <CardActions>
                    <ThemeProvider theme={themeCustom}>
                            <Button color="secondary"  variant="contained" endIcon={<SendIcon />} href={'/projects/'+ information.stageId +'/showcase'}>
                                Demo
                            </Button>
                            <Button color="secondary"  variant="contained" endIcon={<SendIcon />} href={'/projects/'+ information.stageId}>
                                Profile
                            </Button>
                    </ThemeProvider>
                </CardActions>
            </Card>
        </div>
    );
}