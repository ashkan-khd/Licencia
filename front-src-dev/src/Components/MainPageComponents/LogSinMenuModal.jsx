import React, {Component} from 'react'
import {Button, Header, Image, Modal} from 'semantic-ui-react'
import {greenColor, loginMenu, signUpMenu} from "../../Js Functionals/MainPage/Login SignUp Show";
import licenciaImg from "../../Pics/Licencia-Logo.png";
import MainLoginMenu from "./mainLoginMenu";
import MainSignUpMenu from "./mainSignUpMenu";

class ModalLogSin extends Component{
    constructor(props, context) {
        super(props, context);
        this.setOpen = this.setOpen.bind(this)
    }

    state={
        open:false
    }

    setOpen(bool){
        this.setState({open:bool})
    }

    render() {
        let style1 = {
            display: "none"
        }
        let style2 = {
            backgroundColor: greenColor
        }
        return (
            <Modal
                onClose={()=>{
                    this.setOpen(false)
                    // emptyLoginFields()
                    // emptySignUpFields()
                }}
                onOpen={()=>{
                    this.setOpen(true)
                }}
                open={this.state.open}
                trigger={<Button className="loginButton">ورود / ثبت نام</Button>}
            >
                <Modal.Content>
                    <div className="header" id="Login-Menu-Header">
                        <div id="Signup-Login">
                            <div style={style2} className="Signup-login-text" id="LoginMenuButton"
                                 onClick={() => loginMenu()}>ورود
                            </div>
                            <div className="Signup-login-text" id="SignUpMenuButton" onClick={() => signUpMenu()}>ثبت
                                نام
                            </div>
                        </div>
                        <div className="image content">
                            <img src={licenciaImg} id="logoImage" alt="logoLicencia"/>
                        </div>
                        <h3 id="welcomeHeader">Welcome To Licencia</h3>
                    </div>
                    <MainLoginMenu onClose={()=>{
                        this.setState({open:false})
                    }} id='Login-Menu'/>
                    <MainSignUpMenu onClose={()=>{
                        this.setState({open:false})
                    }} style={style1} id='SignUp-Menu'/>
                </Modal.Content>
            </Modal>
        )
    }
}

export default ModalLogSin

